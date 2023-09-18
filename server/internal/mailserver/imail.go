package mailserver

import "context"

type MailServer interface {
	Close() error
	SendAccountVerification(context.Context, MailVerifyRequest) (string, error)
}

type MailVerifyRequest struct {
	Recipient string
}
