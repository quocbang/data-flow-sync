package repositories

import "time"

type SignInRequest struct {
	UserID   string
	Password string
	Options  Option
}

type SignInReply commonWithToken

type Option struct {
	TokenLifeTime time.Duration
}

type CreateAccountRequest struct {
	UserID   string
	Email    string
	Password string
}

type SignUpAccountRequest struct {
	CreateAccountRequest
	Option
}

type CreateAccountReply struct {
	RowsAffected RowsAffected
}

type DeleteAccountRequest struct {
	UserID string
}

type VerifyAccountRequest struct {
	Otp    string
	UserID string
	Option
}

type VerifyAccountReply commonWithToken

type SendMailRequest struct {
	UserID string
	Email  string
}
