package services

import (
	"github.com/go-openapi/runtime/middleware"
	"gitlab.com/quocbang/data-flow-sync/server/swagger/models"
	"gitlab.com/quocbang/data-flow-sync/server/swagger/restapi/operations/account"
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
	Logout(params account.LogoutParams) middleware.Responder
}
