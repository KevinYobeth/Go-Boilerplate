package services

import (
	"go-boilerplate/internal/authors/infrastructure/repository"
	"go-boilerplate/internal/authors/services/command"
	"go-boilerplate/internal/authors/services/query"
	"go-boilerplate/shared/database"
	"go-boilerplate/shared/event"
	"go-boilerplate/shared/log"
	"go-boilerplate/shared/metrics"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateAuthor command.CreateAuthorHandler
	DeleteAuthor command.DeleteAuthorHandler
}

type Queries struct {
	GetAuthors query.GetAuthorsHandler
	GetAuthor  query.GetAuthorHandler
}

func NewAuthorService() Application {
	db := database.InitPostgres()
	logger := log.InitLogger()
	metricsClient := metrics.InitClient()

	repository := repository.NewAuthorsPostgresRepository(db)
	publisher := event.InitPublisher(event.PublisherOptions{
		Topic: "authors",
	})

	return Application{
		Commands: Commands{
			CreateAuthor: command.NewCreateAuthorHandler(repository, logger, metricsClient),
			DeleteAuthor: command.NewDeleteAuthorHandler(repository, publisher, logger, metricsClient),
		},
		Queries: Queries{
			GetAuthors: query.NewGetAuthorsHandler(repository, logger, metricsClient),
			GetAuthor:  query.NewGetAuthorHandler(repository, logger, metricsClient),
		},
	}
}
