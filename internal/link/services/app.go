package services

import (
	"github.com/kevinyobeth/go-boilerplate/internal/link/infrastructure/repository"
	"github.com/kevinyobeth/go-boilerplate/internal/link/services/command"
	"github.com/kevinyobeth/go-boilerplate/internal/link/services/query"
	"github.com/kevinyobeth/go-boilerplate/shared/database"
	"github.com/kevinyobeth/go-boilerplate/shared/log"
	"github.com/kevinyobeth/go-boilerplate/shared/metrics"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	ShortenLink command.ShortenLinkHandler
}

type Queries struct {
	GetLinks        query.GetLinksHandler
	GetRedirectLink query.GetRedirectLinkHandler
}

func NewLinkService() Application {
	db := database.InitPostgres()
	logger := log.InitLogger()
	metricsClient := metrics.InitClient()

	repository := repository.NewLinkPostgresRepository(db)

	return Application{
		Commands: Commands{
			ShortenLink: command.NewShortenLinkHandler(repository, logger, metricsClient),
		},
		Queries: Queries{
			GetLinks:        query.NewGetLinksHandler(repository, logger, metricsClient),
			GetRedirectLink: query.NewGetRedirectLinkHandler(repository, logger, metricsClient),
		},
	}
}
