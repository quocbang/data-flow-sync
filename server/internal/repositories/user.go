package repositories

import "time"

type SignInRequest struct {
	Identifier string
	Password   string
	Options    Option
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
}

type CreateAccountReply struct {
	RowsAffected RowsAffected
}

type DeleteAccountRequest struct {
	UserID string
}

type SendMailRequest struct {
	Email string
}
