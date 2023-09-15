package repositories

import (
	"context"

	"github.com/quocbang/data-flow-sync/server/internal/repositories/orm/models"
)

type Services interface {
	Account() AccountServices
	Station() StationServices
	StationGroup() StationGroupServices
}

type StationServices interface {
	CreateMergeRequest(context.Context) (CreateStationMRReply, error)
	GetMergeRequest(context.Context, GetMRRequest) (GetMRReply, error)
}

type StationGroupServices interface {
}

type AccountServices interface {
	DeleteAccount(context.Context, DeleteAccountRequest) (CommonUpdateAndDeleteReply, error)
	GetAccount(context.Context, string) (models.Account, error)
	SignUp(context.Context, SignUpAccountRequest) error
	UpdateToUserRole(context.Context, string) (CommonUpdateAndDeleteReply, error)
}
