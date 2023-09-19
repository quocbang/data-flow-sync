package account

import (
	"context"
	"fmt"
	"net/http"
	"time"

	apiErrors "github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	"github.com/quocbang/data-flow-sync/server/internal/mailserver"
	"github.com/quocbang/data-flow-sync/server/internal/repositories"
	e "github.com/quocbang/data-flow-sync/server/internal/repositories/errors"
	s "github.com/quocbang/data-flow-sync/server/internal/services"
	"github.com/quocbang/data-flow-sync/server/swagger/models"
	"github.com/quocbang/data-flow-sync/server/swagger/restapi/operations/account"
	"github.com/quocbang/data-flow-sync/server/utils"
	"github.com/quocbang/data-flow-sync/server/utils/function"
	"github.com/quocbang/data-flow-sync/server/utils/roles"
)

type key string

const (
	AuthorizationKey key = "x-data-flow-sync-auth-key"
	SecretAccessKey  key = "secret-access-key"
)

type Authorization struct {
	smtp          mailserver.MailServer
	repo          repositories.Repositories
	tokenLifeTime time.Duration
	hasPermission func(function.FuncName, roles.Roles) bool
	secretKey     string
}

func NewAuthorization(repo repositories.Repositories,
	tokenLifeTime time.Duration,
	hasPermission func(function.FuncName, roles.Roles) bool,
	smtp mailserver.MailServer,
	secretKey string,
) s.AccountServices {
	return Authorization{
		smtp:          smtp,
		repo:          repo,
		tokenLifeTime: tokenLifeTime,
		hasPermission: hasPermission,
		secretKey:     secretKey,
	}
}

func (a Authorization) Auth(token string) (*models.Principal, error) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, SecretAccessKey, a.secretKey)
	claims, err := a.repo.Account().Authorization(ctx, token)
	if err != nil {
		if e, ok := e.As(err); ok {
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

	ctx := params.HTTPRequest.Context()
	ctx = context.WithValue(ctx, SecretAccessKey, a.secretKey)
	signInReply, err := a.repo.Account().SignIn(ctx, signInRequest)
	if err != nil {
		return utils.ParseError(ctx, account.NewLoginDefault(0), err)
	}

	return account.NewLoginOK().WithPayload(&models.Token{
		Token: signInReply.Token,
	})
}

func (a Authorization) Logout(params account.LogoutParams, principal *models.Principal) middleware.Responder {
	token := params.HTTPRequest.Header.Get(string(AuthorizationKey))

	ctx := params.HTTPRequest.Context()
	ctx = context.WithValue(ctx, SecretAccessKey, a.secretKey)
	if err := a.repo.Account().SignOut(ctx, token); err != nil {
		return utils.ParseError(ctx, account.NewLogoutDefault(0), err)
	}
	return account.NewLogoutOK()
}

func (a Authorization) SignUp(params account.SignUpParams) middleware.Responder {
	ctx := params.HTTPRequest.Context() // transfer the logger in middleware to repositories layer.
	ctx = context.WithValue(ctx, SecretAccessKey, a.secretKey)
	signUpRequest := repositories.SignUpAccountRequest{
		CreateAccountRequest: repositories.CreateAccountRequest{
			UserID:   params.SignUp.Name,
			Email:    params.SignUp.Email,
			Password: params.SignUp.Password,
		},
	}

	tx, err := a.repo.Begin(ctx)
	if err != nil {
		return utils.ParseError(ctx, account.NewSignUpDefault(http.StatusInternalServerError), err)
	}

	good := false
	defer func() {
		tx.RollBack()
		if !good {
			tx.Account().DelOTP(ctx, params.SignUp.Email)
		}
	}()

	// add new account
	if err := tx.Account().SignUp(ctx, signUpRequest); err != nil {
		return utils.ParseError(ctx, account.NewSignUpDefault(http.StatusInternalServerError), err)
	}

	// send verify mail
	otp, err := a.smtp.SendAccountVerification(ctx, mailserver.MailVerifyRequest{Recipient: params.SignUp.Email})
	if err != nil {
		return utils.ParseError(ctx, account.NewSignUpDefault(http.StatusInternalServerError), err)
	}

	// add OTP to cache
	err = tx.Account().AddOTP(ctx, params.SignUp.Email, otp)
	if err != nil {
		return utils.ParseError(ctx, account.NewSignUpDefault(http.StatusInternalServerError), err)
	}

	// all good
	tx.Commit()
	good = true

	// login with to get token.
	reply, err := a.repo.Account().SignIn(ctx, repositories.SignInRequest{
		Identifier: params.SignUp.Name,
		Password:   params.SignUp.Password,
		Options: repositories.Option{
			TokenLifeTime: a.tokenLifeTime,
		},
	})
	if err != nil {
		return utils.ParseError(ctx, account.NewSignUpDefault(0), err)
	}

	return account.NewLoginOK().WithPayload(&models.Token{
		Token: reply.Token,
	})
}

func (a Authorization) VerifyAccount(params account.VerifyAccountParams, principal *models.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	if !principal.IsUnspecifiedUser {
		return utils.ParseError(ctx, account.NewVerifyAccountDefault(http.StatusBadRequest), fmt.Errorf("user been verified"))
	}

	// get OTP in store.
	StoredOTP, err := a.repo.Account().GetOTPByEmail(ctx, principal.Email)
	if err != nil {
		return utils.ParseError(ctx, account.NewVerifyAccountDefault(0), err)
	}

	// compare OTP
	OTPProvidedByUser := params.AccountVerify.Otp
	if OTPProvidedByUser != StoredOTP {
		return account.NewVerifyAccountBadRequest().WithPayload(&models.ErrorResponse{
			Code:    int64(e.Code_WRONG_OPT),
			Details: "wrong OTP",
		})
	}

	// upgrade to user role.
	if _, err := a.repo.Account().UpdateToUserRole(ctx, principal.Email); err != nil {
		return utils.ParseError(ctx, account.NewVerifyAccountDefault(0), err)
	}

	// get user info with id in principal.
	acc, err := a.repo.Account().GetAccount(ctx, principal.ID)
	if err != nil {
		return utils.ParseError(ctx, account.NewVerifyAccountDefault(0), err)
	}

	// generate new token with new role.
	token, err := acc.GenerateJWT(ctx, a.tokenLifeTime, a.secretKey)
	if err != nil {
		return utils.ParseError(ctx, account.NewVerifyAccountDefault(0), err)
	}

	return account.NewVerifyAccountOK().WithPayload(&models.Token{
		Token: token,
	})
}

func (a Authorization) SendMail(params account.SendMailParams, principal *models.Principal) middleware.Responder {
	ctx := context.Background()
	if !principal.IsUnspecifiedUser {
		return utils.ParseError(ctx, account.NewSendMailDefault(http.StatusBadRequest), fmt.Errorf("user been verified"))
	}

	sendMailRequest := mailserver.MailVerifyRequest{
		Recipient: principal.Email,
	}

	otp, err := a.smtp.SendAccountVerification(ctx, sendMailRequest)
	if err != nil {
		return utils.ParseError(ctx, account.NewSendMailDefault(http.StatusInternalServerError), err)
	}

	err = a.repo.Account().AddOTP(ctx, principal.Email, otp)
	if err != nil {
		return utils.ParseError(ctx, account.NewSendMailDefault(http.StatusInternalServerError), err)
	}

	return account.NewSendMailOK()
}
