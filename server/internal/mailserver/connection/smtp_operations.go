package mail_connection

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	mailserver "github.com/quocbang/data-flow-sync/server/internal/mailserver"
)

func (s *SMTP) Close() error {
	if err := s.smtp.Close(); err != nil {
		return err
	}

	return nil
}

func OptCreator() string {
	// Initialize the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random 6-digit number
	randOTP := rand.Intn(999999)

	stringOTP := fmt.Sprintf("%v", randOTP)

	for i := 0; i < 6-len(stringOTP); i++ {
		stringOTP = "0" + stringOTP
	}

	return stringOTP
}

func (s SMTP) SendAccountVerification(ctx context.Context, req mailserver.MailVerifyRequest) (string, error) {
	// set sender
	if err := s.smtp.Mail(s.sender); err != nil {
		return "", err
	}

	if err := s.smtp.Rcpt(req.Recipient); err != nil {
		return "", err
	}

	// generate otp
	otp := OptCreator()

	// Compose the HTML email message
	// html active form
	message := "To: " + req.Recipient + "\n" +
		"Subject: OTP verifier\n" +
		"MIME-Version: 1.0\n" +
		"Content-Type: text/html; charset=\"utf-8\"\n" +
		"\n" +
		"<html><body><div class='active-form' style='display: flex; justify-content: center'>" +
		"<div style='background-color: rgba(228, 241, 254, 1); height: 600px; width: 600px;text-align: center; font-size: large'>" +
		"<h1>Data Flow Sync active mail</h1><p>you just registered an account at data flow sync, here is verify code:</p><h1>" + otp + "</h1>" +
		"<p><strong>*Note:</strong> this code will be available within two minute </p></div></div></body></html>"

	wc, err := s.smtp.Data()
	if err != nil {
		return "", err
	}
	defer wc.Close()

	_, err = fmt.Fprintf(wc, message)
	if err != nil {
		return "", err
	}

	return otp, nil
}
