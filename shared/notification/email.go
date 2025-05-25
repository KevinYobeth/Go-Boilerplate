package notification

import (
	"github.com/kevinyobeth/go-boilerplate/config"
	"github.com/ztrue/tracerr"
	gomail "gopkg.in/mail.v2"
)

type EmailNotification struct {
	Dialer *gomail.Dialer
}

func NewEmailNotificationStrategy() (Notification, error) {
	config := config.LoadSMTPConfig()

	dialer := gomail.NewDialer(
		config.SMTPHost,
		config.SMTPPort,
		config.SMTPUsername,
		config.SMTPPassword,
	)

	dial, err := dialer.Dial()
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	defer dial.Close()

	return &EmailNotification{
		Dialer: dialer,
	}, nil
}

func (e *EmailNotification) Send(from string, to []string, subject string, content string) error {
	message := gomail.NewMessage()

	message.SetHeader("From", from)
	message.SetHeader("To", to...)
	message.SetHeader("Subject", subject)

	message.SetBody("text/plain", content)

	err := e.Dialer.DialAndSend(message)
	if err != nil {
		return tracerr.Wrap(err)
	}

	return nil
}
