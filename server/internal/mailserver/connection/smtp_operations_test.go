package mail_connection

import (
	"context"
	"fmt"
	"io"
	"testing"

	"bou.ke/monkey"
	"github.com/go-openapi/swag"
	"github.com/jessevdk/go-flags"
	"github.com/quocbang/data-flow-sync/server/config"
	mailserver "github.com/quocbang/data-flow-sync/server/internal/mailserver"
	"github.com/stretchr/testify/assert"
)

var smtpTest struct {
	SmtpServer  string `long:"smtp-server" description:"the smtp server" env:"SMTP_SERVER_TEST"`
	SmtpPort    int    `long:"smtp-port" description:"the smtp port" env:"SMTP_PORT_TEST"`
	SenderEmail string `long:"smtp-sender" description:"sender email" env:"SMTP_SENDER_TEST"`
	Password    string `long:"smtp-sender-password" description:"sender's password" env:"SMTP_PASSWORD_TEST"`
}

func parseFlags() error {
	configuration := []swag.CommandLineOptionsGroup{
		{
			ShortDescription: "smtp configuration",
			LongDescription:  "smtp configuration",
			Options:          &smtpTest,
		},
	}

	parse := flags.NewParser(nil, flags.IgnoreUnknown)
	for _, opt := range configuration {
		if _, err := parse.AddGroup(opt.LongDescription, opt.LongDescription, opt.Options); err != nil {
			return err
		}
	}

	if _, err := parse.Parse(); err != nil {
		return fmt.Errorf("failed to parse postgres flags")
	}

	return nil
}

func Test_SendAccountVerification(t *testing.T) {
	assertion := assert.New(t)
	ctx := context.Background()
	// Arrange

	// patch send mail method to reduce the real mail sending
	monkey.Patch(fmt.Fprintln, func(w io.Writer, a ...any) (int, error) {
		return 0, nil
	})
	monkey.Patch(OptCreator, func() string {
		return "111111"
	})
	defer monkey.UnpatchAll()

	assertion.NoError(parseFlags())

	mail_s, err := NewSMTP(config.SmtpConfig{
		SmtpServer:  smtpTest.SmtpServer,
		SmtpPort:    smtpTest.SmtpPort,
		SenderEmail: smtpTest.SenderEmail,
		Password:    smtpTest.Password,
	})
	assertion.NoError(err)

	// Act
	otp, err := mail_s.SendAccountVerification(ctx, mailserver.MailVerifyRequest{
		Recipient: "mori@kenda.com.tw",
	})

	// Assert
	assertion.NoError(err)
	assertion.Equal("111111", otp)
}
