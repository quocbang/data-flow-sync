package mail_connection

import (
	"crypto/tls"
	"fmt"
	"net/smtp"

	"github.com/quocbang/data-flow-sync/server/config"
)

type SMTP struct {
	smtp   *smtp.Client
	sender string
}

func NewSMTP(config config.SmtpConfig) (*SMTP, error) {
	smtp, err := NewSMTPConnection(config)
	return &SMTP{
		smtp: smtp,
	}, err
}

func NewSMTPConnection(config config.SmtpConfig) (*smtp.Client, error) {
	// Create an authentication mechanism
	auth := smtp.PlainAuth("", config.SenderEmail, config.Password, config.SmtpServer)

	// Create a TLS configuration
	tlsConfig := &tls.Config{
		InsecureSkipVerify: false, // You might want to set this to false in production
		ServerName:         config.SmtpServer,
	}

	// Connect to the SMTP server with a TLS connection
	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", config.SmtpServer, config.SmtpPort), tlsConfig)
	if err != nil {
		return nil, err
	}

	// Connect to the SMTP server
	// Establish the SMTP client
	client, err := smtp.NewClient(conn, config.SmtpServer)
	if err != nil {
		return nil, err
	}

	// Authenticate
	if err := client.Auth(auth); err != nil {
		return nil, err
	}

	return client, nil
}
