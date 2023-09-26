package services

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/quocbang/data-flow-sync/server/swagger/models"
	"github.com/quocbang/data-flow-sync/server/swagger/restapi/operations/account"
	"github.com/quocbang/data-flow-sync/server/swagger/restapi/operations/station"
)

type Services struct {
	Account      AccountServices
	Station      StationServices
	StationGroup StationGroupServices
}

func RegisterService(account AccountServices, station StationServices) *Services {
	return &Services{
		Account: account,
		Station: station,
	}
}

type StationServices interface {
	CreateStationMergeRequest(params station.CreateStationMergeRequestParams, principal *models.Principal) middleware.Responder
	GetStationMergeRequest(params station.GetStationMergeRequestParams, principal *models.Principal) middleware.Responder
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
	SignUp(params account.SignUpParams) middleware.Responder
	VerifyAccount(params account.VerifyAccountParams, principal *models.Principal) middleware.Responder
	SendMail(params account.SendMailParams, principal *models.Principal) middleware.Responder
}
