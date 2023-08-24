package mail

import (
	"os"

	"gopkg.in/yaml.v3"
)

var (
	emailsPath = "default_mails.yaml"
)

var emails struct {
	otpSender string `yaml:"otp-mail-sender"`
}

func parseConfig(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, &emails); err != nil {
		return err
	}

	return nil
}

func GetOtpMailSender() string {
	if emails.otpSender == "" {
		parseConfig(emailsPath)
	}
	return emails.otpSender
}
