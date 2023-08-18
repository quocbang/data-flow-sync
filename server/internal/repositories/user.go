package repositories

import "time"

type SignInRequest struct {
	UserID   string
	Password string
	Options  Option
}

type SignInReply struct {
	Token string
}

type Option struct {
	TokenLifeTime time.Duration
}

type CreateAccountRequest struct {
	UserID   string
	Password string
}

type CreateAccountReply struct {
	RowsAffected RowsAffected
}

type DeleteAccountRequest struct {
	UserID string
}
