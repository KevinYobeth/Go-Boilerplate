package services

import (
	"github.com/kevinyobeth/go-boilerplate/internal/notification/services/command"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/builder/notification"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/log"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/metrics"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	SendWelcomeNotification command.SendWelcomeNotificationHandler
}

type Queries struct {
}

func NewNotificationService() Application {
	logger := log.InitLogger()
	metricsClient := metrics.InitClient()

	notificationStrategy := notification.NewEmailNotification()
	emailNotification := notification.NewNotification(notificationStrategy)

	return Application{
		Commands: Commands{
			SendWelcomeNotification: command.NewSendWelcomeNotificationHandler(emailNotification, logger, metricsClient),
		},
		Queries: Queries{},
	}
}
