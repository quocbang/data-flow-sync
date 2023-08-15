package models

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"

	repoErrors "github.com/quocbang/data-flow-sync/server/internal/repositories/errors"
	"github.com/quocbang/data-flow-sync/server/utils/roles"
)

type Account struct {
	UserID   string      `gorm:"column:id;type:text;primaryKey"`
	Password []byte      `gorm:"type:bytea;not null"`
	Role     roles.Roles `gorm:"not null"`
}

func (a *Account) TableName() string {
	return "account"
}

type JwtCustomClaims struct {
	UserID            string      `json:"user_id"`
	Role              roles.Roles `json:"role"`
	IsUnspecifiedUser bool        `json:"is_unspecified_user"`
	jwt.StandardClaims
}

// generate JWT.
func (a Account) GenerateJWT(ctx context.Context, tokenLifeTime time.Duration, secretKey string) (string, error) {
	if secretKey == "" {
		return "", repoErrors.Error{
			Details: "secret key not found",
		}
	}

	claims := &JwtCustomClaims{
		UserID:            a.UserID,
		Role:              a.Role,
		IsUnspecifiedUser: a.IsUnSpecifiedUser(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * tokenLifeTime).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// check whether is unspecified user
func (a Account) IsUnSpecifiedUser() bool {
	return a.Role == roles.Roles_UNSPECIFIED
}

// ToHashPassword hashes the password using bcrypt
func ToHashPassword(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hashedPassword, nil
}

// VerifyToken is check the validity of the token and return contents.
func VerifyToken(token string, secretKey string) (*JwtCustomClaims, error) {
	if token == "" {
		return nil, repoErrors.Error{
			Code:    repoErrors.Code_MISSING_TOKEN,
			Details: "authorize token required",
		}
	}
	if secretKey == "" {
		return nil, repoErrors.Error{
			Code:    repoErrors.Code_FORBIDDEN,
			Details: "secret key not found",
		}
	}

	JWTToken, err := jwt.ParseWithClaims(token, &JwtCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, repoErrors.Error{
				Code:    repoErrors.Code_FORBIDDEN,
				Details: "invalid signing method",
			}
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		if e, ok := err.(*jwt.ValidationError); ok {
			if e.Errors == jwt.ValidationErrorExpired {
				return nil, repoErrors.Error{
					Code:    repoErrors.Code_TOKEN_EXPIRED,
					Details: err.Error(),
				}
			}
		}
		return nil, err
	}

	claims, ok := JWTToken.Claims.(*JwtCustomClaims)
	if !ok || !JWTToken.Valid {
		return nil, repoErrors.Error{
			Code:    repoErrors.Code_INVALID_TOKEN,
			Details: "token invalid",
		}
	}

	return claims, nil
}
