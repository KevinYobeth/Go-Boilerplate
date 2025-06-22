package services

import (
	"github.com/kevinyobeth/go-boilerplate/internal/user/infrastructure/repository"
	"github.com/kevinyobeth/go-boilerplate/internal/user/services/query"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/database"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/log"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/metrics"
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
