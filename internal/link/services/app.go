package services

import (
	"github.com/kevinyobeth/go-boilerplate/internal/link/infrastructure/repository"
	"github.com/kevinyobeth/go-boilerplate/internal/link/services/command"
	"github.com/kevinyobeth/go-boilerplate/internal/link/services/query"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/cache"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/database"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/log"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/metrics"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	ShortenLink             command.ShortenLinkHandler
	UpdateLink              command.UpdateLinkHandler
	DeleteLink              command.DeleteLinkHandler
	UpdateLinkVisitSnapshot command.UpdateLinkVisitSnapshot
}

type Queries struct {
	GetLinks        query.GetLinksHandler
	GetLink         query.GetLinkHandler
	GetRedirectLink query.GetRedirectLinkHandler
}

func NewLinkService() Application {
	db := database.InitPostgres()
	cache := cache.InitRedis()

	logger := log.InitLogger()
	metricsClient := metrics.InitClient()

	cacheRepository := repository.NewLinkRedisCache(cache)
	repository := repository.NewLinkPostgresRepository(db)

	return Application{
		Commands: Commands{
			ShortenLink:             command.NewShortenLinkHandler(repository, cacheRepository, logger, metricsClient),
			UpdateLink:              command.NewUpdateLinkHandler(repository, cacheRepository, logger, metricsClient),
			DeleteLink:              command.NewDeleteLinkHandler(repository, logger, metricsClient),
			UpdateLinkVisitSnapshot: command.NewUpdateLinkVisitSnapshotHandler(repository, logger, metricsClient),
		},
		Queries: Queries{
			GetLinks:        query.NewGetLinksHandler(repository, logger, metricsClient),
			GetLink:         query.NewGetLinkHandler(repository, logger, metricsClient),
			GetRedirectLink: query.NewGetRedirectLinkHandler(repository, cacheRepository, logger, metricsClient),
		},
	}
}
