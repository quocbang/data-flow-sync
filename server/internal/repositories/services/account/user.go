package account

import (
	"context"
	"fmt"
	"math/rand"
	"net/smtp"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/quocbang/data-flow-sync/server/internal/repositories"
	e "github.com/quocbang/data-flow-sync/server/internal/repositories/errors"
	"github.com/quocbang/data-flow-sync/server/internal/repositories/orm/models"
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
func (s service) getAccount(ctx context.Context, userID string) (models.Account, error) {
	user := models.Account{}
	if err := s.pg.Where(`id=?`, userID).Take(&user).Error; err != nil {
		return models.Account{}, err
	}

	return user, nil
}

// CreateAccount is create new account
func (s service) createAccount(ctx context.Context, req repositories.CreateAccountRequest) (repositories.CreateAccountReply, error) {
	// hash password
	pwd, err := toHashPassword(req.Password)
	if err != nil {
		return repositories.CreateAccountReply{}, e.Error{
			Detail: fmt.Sprintf("failed to generate password, error: %v", err.Error()),
		}
	}

	reply := s.pg.Create(&models.Account{
		UserID:   req.UserID,
		Password: pwd,
		Roles:    req.Roles,
	})

	return repositories.CreateAccountReply{RowsAffected: reply.RowsAffected}, reply.Error
}

// DeleteAccount is delete existing account
func (s service) DeleteAccount(ctx context.Context, req repositories.DeleteAccountRequest) (repositories.CommonUpdateAndDeleteReply, error) {
	reply := s.pg.Where(`id=?`, req.UserID).Delete(&models.Account{})
	return repositories.CommonUpdateAndDeleteReply{RowsAffected: reply.RowsAffected}, reply.Error
}

// SignIn is access system with user and password
func (s service) SignIn(ctx context.Context, req repositories.SignInRequest) (repositories.SignInReply, error) {
	userInfo, err := s.getAccount(ctx, req.UserID)
	if err != nil {
		return repositories.SignInReply{}, err
	}

	// compare given password and stored password.
	if err := bcrypt.CompareHashAndPassword(userInfo.Password, []byte(req.Password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			// Passwords don't match, handle the invalid login
			return repositories.SignInReply{}, e.Error{
				Code:   e.Code_WRONG_PASSWORD,
				Detail: "wrong password",
			}
		} else {
			// Handle the error
			return repositories.SignInReply{}, err
		}
	}

	// create JWT.
	token, err := s.generateJWT(ctx, userInfo)
	if err != nil {
		return repositories.SignInReply{}, err
	}

	return repositories.SignInReply{Token: token}, nil
}

// SignOut is logout the system and save token to black list of
// the time haven't expired
func (s service) SignOut(ctx context.Context) error {
	return nil
}

// toHashPassword hashes the password using bcrypt
func toHashPassword(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hashedPassword, nil
}

type JwtCustomClaims struct {
	Email      string    `json:"email"`
	ExpiryTime time.Time `json:"expiry_time"`
	Roles      []int64   `json:"roles"`
	jwt.StandardClaims
}

// generate JWT.
func (s *service) generateJWT(ctx context.Context, userInfo models.Account) (string, error) {
	if secretKey == "" {
		return "", e.Error{
			Detail: "secret key not found",
		}
	}

	claims := &JwtCustomClaims{
		Email:      userInfo.UserID,
		ExpiryTime: time.Now().Add(time.Hour * 8),
		Roles:      userInfo.Roles,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s service) SignUp(ctx context.Context, req repositories.CreateAccountRequest) (repositories.SignInReply, error) {
	_, err := s.createAccount(ctx, req)
	if err != nil {
		return repositories.SignInReply{}, err
	}

	return repositories.SignInReply{}, nil
}

func (s service) VerifyAccount(ctx context.Context, req repositories.VerifyAccountRequest) (repositories.VerifyAccountReply, error) {

	return repositories.VerifyAccountReply{}, nil
}

func optCreator() int {
	// Initialize the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random 6-digit number
	randOTP := rand.Intn(900000) + 100000

	return randOTP
}

func (s service) SendMail(ctx context.Context, recipience string) error {
	if err := s.smtp.Rcpt(recipience); err != nil {
		return err
	}

	// generate otp
	otp := optCreator()

	// Compose the HTML email message
	message := "To: " + recipience + "\n" +
		"Subject: OTP verifier\n" +
		"MIME-Version: 1.0\n" +
		"Content-Type: text/html; charset=\"utf-8\"\n" +
		"\n" +
		"<html><body>" +
		"<h1>Hello from Go!</h1>" +
		"<p> your OTP is<strong>" + fmt.Sprintf("%v", otp) + "</strong></p>" +
		"<p>this will be available within two minute </p>" +
		"</body></html>"

	wc, err := s.smtp.Data()
	if err != nil {
		return err
	}
	defer wc.Close()

	_, err = fmt.Fprintf(wc, message)
	if err != nil {
		return err
	}

	return nil
}
