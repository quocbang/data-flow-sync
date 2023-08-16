package account

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"

	"github.com/quocbang/data-flow-sync/server/internal/repositories"
	s "github.com/quocbang/data-flow-sync/server/internal/services"
	"github.com/quocbang/data-flow-sync/server/swagger/models"
	"github.com/quocbang/data-flow-sync/server/swagger/restapi/operations/account"
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

	return account.NewLoginOK().WithPayload(&models.Token{
		Token: uuid.NewString(),
	})
}

func (a Authorization) Logout(params account.LogoutParams) middleware.Responder {
	return account.NewLogoutOK()
}

func (a Authorization) SignUp(params account.SignupParams) middleware.Responder {

	return nil
}

func (a Authorization) VerifyAccount(parmas account.VerifyAccountParams) middleware.Responder {
	return nil
}

func (a Authorization) SendMail(params account.SendMailParams) middleware.Responder {
	return nil
}
