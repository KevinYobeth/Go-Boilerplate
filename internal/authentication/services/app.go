package services

import (
	"github.com/kevinyobeth/go-boilerplate/internal/authentication/infrastructure/repository"
	"github.com/kevinyobeth/go-boilerplate/internal/authentication/services/command"
	"github.com/kevinyobeth/go-boilerplate/internal/authentication/services/query"
	"github.com/kevinyobeth/go-boilerplate/internal/shared/interfaces"
	"github.com/kevinyobeth/go-boilerplate/internal/shared/queue"
	"github.com/kevinyobeth/go-boilerplate/internal/shared/topic"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/database"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/log"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/metrics"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/rabbitmq"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	Register command.RegisterHandler
}

type Queries struct {
	Login        query.LoginHandler
	RefreshToken query.RefreshTokenHandler
}

func NewAuthenticationService(userService interfaces.UserIntraprocess) Application {
	db := database.InitPostgres()
	logger := log.InitLogger()
	metricsClient := metrics.InitClient()
	publish := rabbitmq.InitPublisher(rabbitmq.PublisherOptions{
		Topic: topic.AuthenticationTopic,
		Queue: queue.AuthenticationQueue,
	})

	publisher := repository.NewRabbitMQAuthenticationPublisher(publish)
	repository := repository.NewAuthenticationPostgresRepository(db)

	return Application{
		Commands: Commands{
			Register: command.NewRegisterHandler(repository, userService, publisher, logger, metricsClient),
		},
		Queries: Queries{
			Login:        query.NewLoginHandler(repository, userService, logger, metricsClient),
			RefreshToken: query.NewRefreshTokenHandler(repository, userService, logger, metricsClient),
		},
	}
}
