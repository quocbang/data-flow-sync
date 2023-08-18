package account

import (
	"context"
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
		IsUnspecifiedUser: claims.IsUnspecifiedUser,
	}, nil
}

func (a Authorization) Login(params account.LoginParams) middleware.Responder {
	signInRequest := repositories.SignInRequest{
		UserID:   *params.Login.Username,
		Password: *params.Login.Password,
		Options: repositories.Option{
			TokenLifeTime: a.tokenLifeTime,
		},
	}

	ctx := context.Background()
	signInReply, err := a.repo.Account().SignIn(ctx, signInRequest)
	if err != nil {
		return utils.ParseError(ctx, account.NewLoginDefault(0), err)
	}

	return account.NewLoginOK().WithPayload(&account.LoginOKBody{
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
