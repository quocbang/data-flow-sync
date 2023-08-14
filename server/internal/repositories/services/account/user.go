package account

import (
	"context"
	"fmt"
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
	pg *gorm.DB
	rd *redis.Client
}

func NewService(pg *gorm.DB, rd *redis.Client) repositories.AccountServices {
	return service{
		pg: pg,
		rd: rd,
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
func (s service) CreateAccount(ctx context.Context, req repositories.CreateAccountRequest) (repositories.CreateAccountReply, error) {
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
