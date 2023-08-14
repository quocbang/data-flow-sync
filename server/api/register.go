package api

import (
	"gitlab.com/quocbang/data-flow-sync/server/internal/repositories"
	"gitlab.com/quocbang/data-flow-sync/server/internal/services"
	a "gitlab.com/quocbang/data-flow-sync/server/internal/services/account"
	"gitlab.com/quocbang/data-flow-sync/server/swagger/restapi/operations"
	"gitlab.com/quocbang/data-flow-sync/server/swagger/restapi/operations/account"
)

type ServiceConfig struct {
	Repo         repositories.Repositories
	MRExpiryTime int64
}

func newHandleService(s ServiceConfig) *services.Services {
	return services.RegisterService(
		a.NewAuthorization(s.Repo),
	)
}

func RegisterAPI(api *operations.DataFlowSyncAPI, serviceConfig ServiceConfig) {
	s := newHandleService(serviceConfig)

	// account
	api.APIKeyAuth = s.Account.Auth
	api.AccountLoginHandler = account.LoginHandlerFunc(s.Account.Login)
	api.AccountLogoutHandler = account.LogoutHandlerFunc(s.Account.Logout)

	// limitary-hour

	// station

	// station group

}
