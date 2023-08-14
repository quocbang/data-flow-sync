package repositories

import "context"

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
	CreateAccount(context.Context, CreateAccountRequest) (CreateAccountReply, error)
	DeleteAccount(context.Context, DeleteAccountRequest) (CommonUpdateAndDeleteReply, error)
	SignIn(context.Context, SignInRequest) (SignInReply, error)
	SignOut(context.Context) error
}
