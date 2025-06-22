package command

import (
	"bytes"
	"context"
	"text/template"

	"github.com/google/uuid"
	"github.com/kevinyobeth/go-boilerplate/config"
	mail_template "github.com/kevinyobeth/go-boilerplate/internal/shared/templates/mail"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/builder/notification"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/decorator"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/metrics"
	"github.com/ztrue/tracerr"
	"go.uber.org/zap"
)

type SendWelcomeNotificationRequest struct {
	UserID uuid.UUID `json:"user_id"`
	Name   string    `json:"name"`
	Email  string    `json:"email"`
}

type sendWelcomeNotificationHandler struct {
	Notification notification.Notification
}

type SendWelcomeNotificationHandler decorator.CommandHandler[*SendWelcomeNotificationRequest]

func (h sendWelcomeNotificationHandler) Handle(c context.Context, params *SendWelcomeNotificationRequest) error {
	tmpl, err := template.ParseFiles("internal/shared/templates/mail/welcome_mail.html")
	if err != nil {
		return tracerr.Wrap(err)
	}

	smtpConfig := config.LoadSMTPConfig()
	config := config.LoadAppConfig()

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, mail_template.WelcomeMail{
		Name:    params.Name,
		AppName: config.AppName,
	})
	if err != nil {
		return tracerr.Wrap(err)
	}

	err = h.Notification.
		To(params.Email).
		From(smtpConfig.SMTPEmailFrom).
		Subject("Welcome to " + config.AppName).
		BodyHTML(buf.String()).
		Send()
	if err != nil {
		return tracerr.Wrap(err)
	}

	return nil
}

func NewSendWelcomeNotificationHandler(notification notification.Notification, logger *zap.SugaredLogger, metricsClient metrics.Client) SendWelcomeNotificationHandler {
	if notification == nil {
		panic("notification cannot be nil")
	}

	return decorator.ApplyCommandDecorators(
		sendWelcomeNotificationHandler{
			Notification: notification,
		},
		logger, metricsClient,
	)
}
