package api

import (
	"time"

	"github.com/quocbang/data-flow-sync/server/internal/mailserver"
	rd_connection "github.com/quocbang/data-flow-sync/server/internal/redis_conn"
	"github.com/quocbang/data-flow-sync/server/internal/repositories"
	"github.com/quocbang/data-flow-sync/server/internal/services"
	a "github.com/quocbang/data-flow-sync/server/internal/services/account"
	stationService "github.com/quocbang/data-flow-sync/server/internal/services/station"
	"github.com/quocbang/data-flow-sync/server/swagger/restapi/operations"
	"github.com/quocbang/data-flow-sync/server/swagger/restapi/operations/account"
	"github.com/quocbang/data-flow-sync/server/swagger/restapi/operations/station"
	"github.com/quocbang/data-flow-sync/server/utils/roles"
)

type ServiceConfig struct {
	Repo          repositories.Repositories
	Smtp          mailserver.MailServer
	Redis         rd_connection.RedisConn
	TokenLifeTime time.Duration
	MRExpiryTime  int64
	SecretKey     string
}

func NewHandleService(s ServiceConfig) *services.Services {
	return services.RegisterService(
		a.NewAuthorization(s.Repo, s.TokenLifeTime, roles.HasPermission, s.Smtp, s.Redis, s.SecretKey),
		stationService.NewStation(s.Repo, s.MRExpiryTime, roles.HasPermission),
	)
}

func RegisterAPI(api *operations.DataFlowSyncAPI, serviceConfig ServiceConfig) {
	s := NewHandleService(serviceConfig)

	// account
	api.APIKeyAuth = s.Account.Auth
	api.AccountLoginHandler = account.LoginHandlerFunc(s.Account.Login)
	api.AccountLogoutHandler = account.LogoutHandlerFunc(s.Account.Logout)
	api.AccountSignUpHandler = account.SignUpHandlerFunc(s.Account.SignUp)
	api.AccountVerifyAccountHandler = account.VerifyAccountHandlerFunc(s.Account.VerifyAccount)
	api.AccountSendMailHandler = account.SendMailHandlerFunc(s.Account.SendMail)

	// limitary-hour

	// station
	api.StationCreateStationMergeRequestHandler = station.CreateStationMergeRequestHandlerFunc(s.Station.CreateStationMergeRequest)
	api.StationGetStationMergeRequestHandler = station.GetStationMergeRequestHandlerFunc(s.Station.GetStationMergeRequest)

	// station group
}
