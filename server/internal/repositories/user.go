package repositories

type SignInRequest struct {
	UserID   string
	Password string
}

type SignInReply commonWithToken

type CreateAccountRequest struct {
	UserID   string  `validate:"required"`
	Password string  `validate:"required"`
	Roles    []int64 `validate:"required,dive,gte=0"`
}

type CreateAccountReply struct {
	RowsAffected RowsAffected
}

type DeleteAccountRequest struct {
	UserID string `validate:"required"`
}

type VerifyAccountRequest commonWithToken

type VerifyAccountReply commonWithToken
