package notification

import (
	"errors"

	"github.com/kevinyobeth/go-boilerplate/config"
	gomail "gopkg.in/mail.v2"
)

type EmailNotification struct {
	Dialer  *gomail.Dialer
	Message *gomail.Message
}

func NewEmailNotification() Notification {
	config := config.LoadSMTPConfig()

	dialer := gomail.NewDialer(
		config.SMTPHost,
		config.SMTPPort,
		config.SMTPUsername,
		config.SMTPPassword,
	)

	message := gomail.NewMessage()

	return &EmailNotification{
		Dialer:  dialer,
		Message: message,
	}
}

func (n *EmailNotification) To(to string) Notification {
	n.Message.SetHeader("To", to)
	return n
}

func (n *EmailNotification) From(from string) Notification {
	n.Message.SetHeader("From", from)
	return n
}

func (n *EmailNotification) Subject(subject string) Notification {
	n.Message.SetHeader("Subject", subject)
	return n
}

func (n *EmailNotification) Body(body string) Notification {
	n.Message.SetBody("text/plain", body)
	return n
}

func (n *EmailNotification) BodyHTML(htmlBody string) Notification {
	n.Message.SetBody("text/html", htmlBody)
	return n
}

func (n *EmailNotification) Cc(cc ...string) Notification {
	n.Message.SetHeader("Cc", cc...)
	return n
}

func (n *EmailNotification) Bcc(bcc ...string) Notification {
	n.Message.SetHeader("Bcc", bcc...)
	return n
}

func (n *EmailNotification) Attachments(attachments ...string) Notification {
	for _, attachment := range attachments {
		n.Message.Attach(attachment)
	}
	return n
}
func (n *EmailNotification) Send() error {
	if n.Dialer == nil {
		return errors.New("email dialer is not initialized")
	}
	if n.Message == nil {
		return errors.New("email message is not initialized")
	}

	if err := n.Dialer.DialAndSend(n.Message); err != nil {
		return err
	}

	return nil
}
