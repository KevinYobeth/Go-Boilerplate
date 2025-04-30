package services

import (
	"go-boilerplate/internal/authentication/infrastructure/repository"
	"go-boilerplate/internal/authentication/services/command"
	"go-boilerplate/internal/authentication/services/query"
	"go-boilerplate/shared/database"
	"go-boilerplate/shared/log"
	"go-boilerplate/shared/metrics"
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
	GetUser      query.GetUserHandler
}

func NewAuthenticationService() Application {
	db := database.InitPostgres()
	logger := log.InitLogger()
	metricsClient := metrics.InitClient()

	repository := repository.NewAuthenticationPostgresRepository(db)

	return Application{
		Commands: Commands{
			Register: command.NewRegisterHandler(repository, logger, metricsClient),
		},
		Queries: Queries{
			Login:        query.NewLoginHandler(repository, logger, metricsClient),
			RefreshToken: query.NewRefreshTokenHandler(repository, logger, metricsClient),
			GetUser:      query.NewGetUserHandler(repository, logger, metricsClient),
		},
	}
}
