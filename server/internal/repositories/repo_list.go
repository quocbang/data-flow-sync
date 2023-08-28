package repositories

import (
	"context"

	"github.com/go-openapi/runtime/middleware"
	"github.com/quocbang/data-flow-sync/server/internal/repositories/orm/models"
	mod "github.com/quocbang/data-flow-sync/server/swagger/models"
	"github.com/quocbang/data-flow-sync/server/swagger/restapi/operations/station"
)

type Services interface {
	Account() AccountServices
	Station() StationServices
	StationGroup() StationGroupServices
}

type StationServices interface {
	UpsertStationDataStorage(params station.CreateStationDataStorageParams, principal *mod.Principal) middleware.Responder
}

type StationGroupServices interface {
}

type AccountServices interface {
	Authorization(context.Context, string) (*models.JwtCustomClaims, error)
	DeleteAccount(context.Context, DeleteAccountRequest) (CommonUpdateAndDeleteReply, error)
	SignIn(context.Context, SignInRequest) (SignInReply, error)
	SignOut(context.Context, string) error
}
