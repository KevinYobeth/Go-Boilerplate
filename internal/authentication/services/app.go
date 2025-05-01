package services

import (
	"github.com/kevinyobeth/go-boilerplate/internal/authentication/infrastructure/repository"
	"github.com/kevinyobeth/go-boilerplate/internal/authentication/services/command"
	"github.com/kevinyobeth/go-boilerplate/internal/authentication/services/query"
	"github.com/kevinyobeth/go-boilerplate/shared/database"
	"github.com/kevinyobeth/go-boilerplate/shared/log"
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
