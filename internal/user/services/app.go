package services

import (
	"github.com/kevinyobeth/go-boilerplate/internal/user/infrastructure/repository"
	"github.com/kevinyobeth/go-boilerplate/internal/user/services/query"
	"github.com/kevinyobeth/go-boilerplate/shared/database"
	"github.com/kevinyobeth/go-boilerplate/shared/log"
	"github.com/kevinyobeth/go-boilerplate/shared/metrics"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
}

type Queries struct {
	GetUser query.GetUserHandler
}

func NewUserService() Application {
	db := database.InitPostgres()
	logger := log.InitLogger()
	metricsClient := metrics.InitClient()

	repository := repository.NewUserPostgresRepository(db)

	return Application{
		Commands: Commands{},
		Queries: Queries{
			GetUser: query.NewGetUserHandler(repository, logger, metricsClient),
		},
	}
}
