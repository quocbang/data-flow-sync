package account

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"

	"gitlab.com/quocbang/data-flow-sync/server/internal/repositories"
	s "gitlab.com/quocbang/data-flow-sync/server/internal/services"
	"gitlab.com/quocbang/data-flow-sync/server/swagger/models"
	"gitlab.com/quocbang/data-flow-sync/server/swagger/restapi/operations/account"
)

type Authorization struct {
	Repo repositories.Repositories
}

func NewAuthorization(repo repositories.Repositories) s.AccountServices {
	return Authorization{
		Repo: repo,
	}
}

func (a Authorization) Auth(token string) (*models.Principal, error) {
	return nil, nil
}

func (a Authorization) Login(params account.LoginParams) middleware.Responder {
	return account.NewLoginOK().WithPayload(&account.LoginOKBody{
		Token: uuid.NewString(),
	})
}

func (a Authorization) Logout(params account.LogoutParams) middleware.Responder {
	return account.NewLogoutOK()
}
