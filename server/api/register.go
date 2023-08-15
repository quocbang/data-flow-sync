package api

import (
	"time"

	"github.com/quocbang/data-flow-sync/server/internal/repositories"
	"github.com/quocbang/data-flow-sync/server/internal/services"
	a "github.com/quocbang/data-flow-sync/server/internal/services/account"
	"github.com/quocbang/data-flow-sync/server/swagger/restapi/operations"
	"github.com/quocbang/data-flow-sync/server/swagger/restapi/operations/account"
	"github.com/quocbang/data-flow-sync/server/utils/roles"
)

type ServiceConfig struct {
	Repo          repositories.Repositories
	TokenLifeTime time.Duration
	MRExpiryTime  int64
}

func newHandleService(s ServiceConfig) *services.Services {
	return services.RegisterService(
		a.NewAuthorization(s.Repo, s.TokenLifeTime, roles.HasPermission),
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
