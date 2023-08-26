package repositories

import (
	"context"

	"github.com/quocbang/data-flow-sync/server/internal/repositories/orm/models"
	model "github.com/quocbang/data-flow-sync/server/swagger/models"
)

type Services interface {
	Account() AccountServices
	Station() StationServices
	StationGroup() StationGroupServices
}

type StationServices interface {
	UpsertStationDataStorage(context.Context, *model.CreateStationDataStorage) (CreateStationDataStorageReply, error)
}

type StationGroupServices interface {
}

type AccountServices interface {
	Authorization(context.Context, string) (*models.JwtCustomClaims, error)
	DeleteAccount(context.Context, DeleteAccountRequest) (CommonUpdateAndDeleteReply, error)
	SignIn(context.Context, SignInRequest) (SignInReply, error)
	SignOut(context.Context, string) error
}
