package account

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/smtp"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/quocbang/data-flow-sync/server/internal/repositories"
	e "github.com/quocbang/data-flow-sync/server/internal/repositories/errors"
	"github.com/quocbang/data-flow-sync/server/internal/repositories/orm/models"
	"github.com/quocbang/data-flow-sync/server/utils/mail"
	"github.com/quocbang/data-flow-sync/server/utils/roles"
)

var secretKey = os.Getenv("DATA_FLOW_SYNC_SECRET_KEY")

type service struct {
	pg   *gorm.DB
	rd   *redis.Client
	smtp *smtp.Client
}

func NewService(pg *gorm.DB, rd *redis.Client, sm *smtp.Client) repositories.AccountServices {
	return service{
		pg:   pg,
		rd:   rd,
		smtp: sm,
	}
}

// getAccount definition.
func (s service) GetAccount(ctx context.Context, Identifier string) (models.Account, error) {
	user := models.Account{}
	err := s.pg.Where(`user_id=?`, Identifier).Take(&user).Error
	if err != nil {
		err = s.pg.Where(`email=?`, Identifier).Take(&user).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Account{}, e.Error{
				Code:    e.Code_ACCOUNT_NOT_FOUND,
				Details: "account not found",
			}
		}
		return models.Account{}, err
	}

	return user, nil
}

// CreateAccount is create new account with unspecified role.
func (s service) createAccount(ctx context.Context, req repositories.CreateAccountRequest) (repositories.CreateAccountReply, error) {
	// hash password
	pwd, err := models.ToHashPassword(req.Password)
	if err != nil {
		return repositories.CreateAccountReply{}, e.Error{
			Details: fmt.Sprintf("failed to generate password, error: %v", err.Error()),
		}
	}

	reply := s.pg.Create(&models.Account{
		UserID:   req.UserID,
		Email:    req.Email,
		Password: pwd,
		Roles:    roles.Roles_UNSPECIFIED,
	})

	return repositories.CreateAccountReply{RowsAffected: reply.RowsAffected}, reply.Error
}

// DeleteAccount is delete existing account
func (s service) DeleteAccount(ctx context.Context, req repositories.DeleteAccountRequest) (repositories.CommonUpdateAndDeleteReply, error) {
	reply := s.pg.Where(`user_id=?`, req.UserID).Delete(&models.Account{})
	return repositories.CommonUpdateAndDeleteReply{RowsAffected: reply.RowsAffected}, reply.Error
}

// SignIn is access system with user and password
func (s service) SignIn(ctx context.Context, req repositories.SignInRequest) (repositories.SignInReply, error) {
	userInfo, err := s.GetAccount(ctx, req.Identifier)
	if err != nil {
		return repositories.SignInReply{}, err
	}

	// compare given password and stored password.
	if err := bcrypt.CompareHashAndPassword(userInfo.Password, []byte(req.Password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			// Passwords don't match, handle the invalid login
			return repositories.SignInReply{}, e.Error{
				Code:    e.Code_WRONG_PASSWORD,
				Details: "wrong password",
			}
		} else {
			// Handle the error
			return repositories.SignInReply{}, err
		}
	}

	// create JWT.
	token, err := userInfo.GenerateJWT(ctx, req.Options.TokenLifeTime, secretKey)
	if err != nil {
		return repositories.SignInReply{}, err
	}

	return repositories.SignInReply{Token: token}, nil
}

// SignOut is logout the system and save token to black list of
// the time haven't expired
func (s service) SignOut(ctx context.Context, token string) error {
	claims, err := models.VerifyToken(token, secretKey)
	if err != nil {
		return err
	}

	timeRemaining := time.Until(time.Unix(claims.ExpiresAt, 0))
	reply := s.rd.Set(ctx, token, nil, timeRemaining)

	return reply.Err()
}

// Authorization verify token and parse it to check auth.
func (s service) Authorization(ctx context.Context, token string) (*models.JwtCustomClaims, error) {
	// check black list
	dataCount, err := s.getBlackList(ctx, token)
	if err != nil {
		return nil, err
	}
	if dataCount > 0 {
		return nil, e.Error{
			Code:    e.Code_TOKEN_BLOCKED,
			Details: "token was blocked",
		}
	}

	// verify token.
	claims, err := models.VerifyToken(token, secretKey)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

// getBlackList get data in black list.
func (s service) getBlackList(ctx context.Context, token string) (int64, error) { // return row number and error.
	return s.rd.Exists(ctx, token).Result()
}

// UpdateRole updates the role for specified account
func (s service) UpdateToUserRole(ctx context.Context, email string) (repositories.CommonUpdateAndDeleteReply, error) {
	reply := s.pg.Model(&models.Account{}).Where(`email = ?`, email).Update("roles", roles.Roles_USER)
	return repositories.CommonUpdateAndDeleteReply{RowsAffected: reply.RowsAffected}, reply.Error
}

func (s service) SignUp(ctx context.Context, req repositories.SignUpAccountRequest) (repositories.SignInReply, error) {
	_, err := s.createAccount(ctx, repositories.CreateAccountRequest{
		UserID:   req.UserID,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return repositories.SignInReply{}, err
	}

	reply, err := s.SignIn(ctx, repositories.SignInRequest{
		Identifier: req.UserID,
		Password:   req.Password,
		Options: repositories.Option{
			TokenLifeTime: req.TokenLifeTime,
		},
	})

	if err != nil {
		return repositories.SignInReply{}, err
	}

	if err := s.SendMail(ctx, repositories.SendMailRequest{
		Email: req.Email,
	}); err != nil {
		return repositories.SignInReply{}, err
	}

	return repositories.SignInReply{Token: reply.Token}, nil
}

func (s service) VerifyAccount(ctx context.Context, req repositories.VerifyAccountRequest) (repositories.VerifyAccountReply, error) {
	// get otp of the account
	actual, err := s.rd.Get(ctx, req.Email).Result()
	if err != nil {
		log.Fatal(err)
	}

	if req.Otp != actual {
		return repositories.VerifyAccountReply{}, e.Error{
			Code:    e.Code_WRONG_OPT,
			Details: "wrong otp",
		}
	}
	if err != nil {
		return repositories.VerifyAccountReply{}, err
	}

	_, err = s.UpdateToUserRole(ctx, req.Email)
	if err != nil {
		return repositories.VerifyAccountReply{}, err
	}

	// Newly update user
	newUser, err := s.GetAccount(ctx, req.Email)
	if err != nil {
		return repositories.VerifyAccountReply{}, err
	}

	// Generate the token for new user
	// create JWT.z
	token, err := newUser.GenerateJWT(ctx, req.Option.TokenLifeTime, secretKey)
	if err != nil {
		return repositories.VerifyAccountReply{}, err
	}

	return repositories.VerifyAccountReply{Token: token}, nil
}

func OptCreator() string {
	// Initialize the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random 6-digit number
	randOTP := rand.Intn(999999)

	stringOTP := fmt.Sprintf("%v", randOTP)

	for i := 0; i < 6-len(stringOTP); i++ {
		stringOTP = "0" + stringOTP
	}

	return stringOTP
}

func (s service) SendMail(ctx context.Context, req repositories.SendMailRequest) error {
	// Set the sender
	if err := s.smtp.Mail(mail.GetOtpMailSender()); err != nil {
		return err
	}

	if err := s.smtp.Rcpt(req.Email); err != nil {
		return err
	}

	// generate otp
	otp := OptCreator()

	// Compose the HTML email message
	// html active form
	message := "To: " + req.Email + "\n" +
		"Subject: OTP verifier\n" +
		"MIME-Version: 1.0\n" +
		"Content-Type: text/html; charset=\"utf-8\"\n" +
		"\n" +
		"<html><body><div class='active-form' style='display: flex; justify-content: center'>" +
		"<div style='background-color: rgba(228, 241, 254, 1); height: 600px; width: 600px;text-align: center; font-size: large'>" +
		"<h1>Data Flow Sync active mail</h1><p>you just registered an account at data flow sync, here is verify code:</p><h1>" + otp + "</h1>" +
		"<p><strong>*Note:</strong> this code will be available within two minute </p></div></div></body></html>"

	wc, err := s.smtp.Data()
	if err != nil {
		return err
	}
	defer wc.Close()

	_, err = fmt.Fprintf(wc, message)
	if err != nil {
		return err
	}

	// Add otp to redis server with 5 minutes expire time
	if err := s.addOTP(ctx, req.Email, otp); err != nil {
		return err
	}

	return nil
}
