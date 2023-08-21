package services

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/quocbang/data-flow-sync/server/swagger/models"
	"github.com/quocbang/data-flow-sync/server/swagger/restapi/operations/account"
)

type Services struct {
	Account      AccountServices
	Station      StationServices
	StationGroup StationGroupServices
}

func RegisterService(account AccountServices) *Services {
	return &Services{
		Account: account,
	}
}

type StationServices interface {
}

type StationGroupServices interface {
}

type LimitaryHourService interface {
	UploadLimitaryHour()
}

type AccountServices interface {
	Auth(string) (*models.Principal, error)
	Login(params account.LoginParams) middleware.Responder
	Logout(params account.LogoutParams, principal *models.Principal) middleware.Responder
	SignUp(params account.SignupParams) middleware.Responder
	VerifyAccount(params account.VerifyAccountParams) middleware.Responder
	SendMail(params account.SendMailParams, principal *models.Principal) middleware.Responder
}
