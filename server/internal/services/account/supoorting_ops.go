package account

import (
	"context"
	"time"

	"github.com/quocbang/data-flow-sync/server/internal/repositories"
	e "github.com/quocbang/data-flow-sync/server/internal/repositories/errors"
	"github.com/quocbang/data-flow-sync/server/internal/repositories/orm/models"
	"golang.org/x/crypto/bcrypt"
)

// Authorization verify token and parse it to check auth.
func (a Authorization) authorization(ctx context.Context, token string) (*models.JwtCustomClaims, error) {
	// check black list
	dataCount, err := a.rd.GetBlackList(ctx, token)
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
	secretKey := ctx.Value(SecretAccessKey).(string)
	claims, err := models.VerifyToken(token, secretKey)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

// SignOut is logout the system and save token to black list of
// the time haven't expired
func (a Authorization) signOut(ctx context.Context, token string) error {
	secretKey := ctx.Value(SecretAccessKey).(string)
	claims, err := models.VerifyToken(token, secretKey)
	if err != nil {
		return err
	}

	timeRemaining := time.Until(time.Unix(claims.ExpiresAt, 0))

	return a.rd.AddToBackList(ctx, token, timeRemaining)
}

// SignIn is access system with user and password
func (a Authorization) signIn(ctx context.Context, username string, password string) (repositories.SignInReply, error) {
	userInfo, err := a.repo.Account().GetAccount(ctx, username)
	if err != nil {
		return repositories.SignInReply{}, err
	}

	// compare given password and stored password.
	if err := bcrypt.CompareHashAndPassword(userInfo.Password, []byte(password)); err != nil {
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
	secretKey := ctx.Value(SecretAccessKey).(string)
	token, err := userInfo.GenerateJWT(ctx, a.tokenLifeTime, secretKey)
	if err != nil {
		return repositories.SignInReply{}, err
	}

	return repositories.SignInReply{Token: token}, nil
}
