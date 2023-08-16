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
	DeleteAccount(context.Context, DeleteAccountRequest) (CommonUpdateAndDeleteReply, error)
	SignIn(context.Context, SignInRequest) (SignInReply, error)
	SignOut(context.Context) error
	SignUp(context.Context, CreateAccountRequest) (SignInReply, error)
	VerifyAccount(context.Context, VerifyAccountRequest) (VerifyAccountReply, error)
	SendMail(context.Context, string) error
}
