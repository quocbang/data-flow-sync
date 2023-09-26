package repositories

import (
	"context"

	"github.com/quocbang/data-flow-sync/server/internal/repositories/orm/models"
)

type Services interface {
	Account() AccountServices
	Station() StationServices
	StationGroup() StationGroupServices
	File() FileServices
	MergeRequest() MergeRequestServices
}

type StationServices interface {
}

type StationGroupServices interface {
}

type FileServices interface {
	GetFile(context.Context, GetFileRequest) (GetFileReply, error)
}

type AccountServices interface {
	DeleteAccount(context.Context, DeleteAccountRequest) (CommonUpdateAndDeleteReply, error)
	GetAccount(context.Context, string) (models.Account, error)
	SignUp(context.Context, SignUpAccountRequest) error
	UpdateToUserRole(context.Context, string) (CommonUpdateAndDeleteReply, error)
}

type MergeRequestServices interface {
	CreateMergeRequest(context.Context, CreateMRRequest) (CreateMRReply, error)
	GetMergeRequest(context.Context, GetMRRequest) (GetMRReply, error)
	GetMergeRequestOpeningByFileID(context.Context, string) (GetMergeRequestOpeningByFileIDReply, error)
}
