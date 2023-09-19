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
}

type StationGroupServices interface {
}

type AccountServices interface {
	DeleteAccount(context.Context, DeleteAccountRequest) (CommonUpdateAndDeleteReply, error)
	GetAccount(context.Context, string) (models.Account, error)
	SignIn(context.Context, SignInRequest) (SignInReply, error)
	SignOut(context.Context, string) error
	SignUp(context.Context, SignUpAccountRequest) error
	UpdateToUserRole(context.Context, string) (CommonUpdateAndDeleteReply, error)
	Authorization(context.Context, string) (*models.JwtCustomClaims, error)
	AddOTP(context.Context, string, string) error
	DelOTP(context.Context, string) error
	GetOTPByEmail(context.Context, string) (string, error)
}
