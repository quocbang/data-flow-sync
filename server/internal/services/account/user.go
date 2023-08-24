package account

import (
	"context"
	"fmt"
	"net/http"
	"time"

	apiErrors "github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	"github.com/quocbang/data-flow-sync/server/internal/repositories"
	repoErrors "github.com/quocbang/data-flow-sync/server/internal/repositories/errors"
	s "github.com/quocbang/data-flow-sync/server/internal/services"
	"github.com/quocbang/data-flow-sync/server/swagger/models"
	"github.com/quocbang/data-flow-sync/server/swagger/restapi/operations/account"
	"github.com/quocbang/data-flow-sync/server/utils"
	"github.com/quocbang/data-flow-sync/server/utils/function"
	"github.com/quocbang/data-flow-sync/server/utils/roles"
)

const AuthorizationKey string = "x-data-flow-sync-auth-key"

type Authorization struct {
	repo          repositories.Repositories
	tokenLifeTime time.Duration
	hasPermission func(function.FuncName, roles.Roles) bool
}

func NewAuthorization(repo repositories.Repositories,
	tokenLifeTime time.Duration,
	hasPermission func(function.FuncName, roles.Roles) bool,
) s.AccountServices {
	return Authorization{
		repo:          repo,
		tokenLifeTime: tokenLifeTime,
		hasPermission: hasPermission,
	}
}

func (a Authorization) Auth(token string) (*models.Principal, error) {
	ctx := context.Background()
	claims, err := a.repo.Account().Authorization(ctx, token)
	if err != nil {
		if e, ok := repoErrors.As(err); ok {
			return nil, apiErrors.New(http.StatusUnauthorized, e.Error())
		}
		return nil, err
	}
	return &models.Principal{
		ID:                claims.UserID,
		Role:              int64(claims.Role),
		Email:             claims.Email,
		IsUnspecifiedUser: claims.IsUnspecifiedUser,
	}, nil
}

func (a Authorization) Login(params account.LoginParams) middleware.Responder {
	signInRequest := repositories.SignInRequest{
		Identifier: *params.Login.Username,
		Password:   *params.Login.Password,
		Options: repositories.Option{
			TokenLifeTime: a.tokenLifeTime,
		},
	}

	ctx := context.Background()
	signInReply, err := a.repo.Account().SignIn(ctx, signInRequest)
	if err != nil {
		return utils.ParseError(ctx, account.NewLoginDefault(0), err)
	}

	return account.NewLoginOK().WithPayload(&models.Token{
		Token: signInReply.Token,
	})
}

func (a Authorization) Logout(params account.LogoutParams, principal *models.Principal) middleware.Responder {
	token := params.HTTPRequest.Header.Get(AuthorizationKey)

	ctx := context.Background()
	if err := a.repo.Account().SignOut(ctx, token); err != nil {
		return utils.ParseError(ctx, account.NewLogoutDefault(0), err)
	}
	return account.NewLogoutOK()
}

func (a Authorization) SignUp(params account.SignupParams) middleware.Responder {
	signUpRequest := repositories.SignUpAccountRequest{
		CreateAccountRequest: repositories.CreateAccountRequest{
			UserID:   params.Signup.Name,
			Email:    params.Signup.Email,
			Password: params.Signup.Password,
		},
		Option: repositories.Option{
			TokenLifeTime: 5 * time.Minute,
		},
	}
	ctx := context.Background()

	reply, err := a.repo.Account().SignUp(ctx, signUpRequest)
	if err != nil {
		return utils.ParseError(ctx, account.NewSignupDefault(http.StatusInternalServerError), err)
	}
	return account.NewLoginOK().WithPayload(&models.Token{
		Token: reply.Token,
	})
}

func (a Authorization) VerifyAccount(params account.VerifyAccountParams, principal *models.Principal) middleware.Responder {
	ctx := context.Background()
	if !principal.IsUnspecifiedUser {
		return utils.ParseError(ctx, account.NewVerifyAccountDefault(http.StatusBadRequest), fmt.Errorf("user been verified"))
	}
	verifyRequest := repositories.VerifyAccountRequest{
		Otp:   params.AccountVerify.Otp,
		Email: principal.Email,
		Option: repositories.Option{
			TokenLifeTime: 5 * time.Minute,
		},
	}

	reply, err := a.repo.Account().VerifyAccount(ctx, verifyRequest)
	if err != nil {
		return utils.ParseError(ctx, account.NewVerifyAccountDefault(http.StatusInternalServerError), err)
	}

	return account.NewVerifyAccountOK().WithPayload(&models.Token{
		Token: reply.Token,
	})
}

func (a Authorization) SendMail(params account.SendMailParams, principal *models.Principal) middleware.Responder {
	ctx := context.Background()
	if !principal.IsUnspecifiedUser {
		return utils.ParseError(ctx, account.NewVerifyAccountDefault(http.StatusBadRequest), fmt.Errorf("user been verified"))
	}

	sendMailRequest := repositories.SendMailRequest{
		Email: principal.Email,
	}

	err := a.repo.Account().SendMail(ctx, sendMailRequest)
	if err != nil {
		return utils.ParseError(ctx, account.NewSendMailDefault(http.StatusInternalServerError), err)
	}

	return account.NewSendMailOK()
}
