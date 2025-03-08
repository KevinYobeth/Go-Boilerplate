package services

import (
	"go-boilerplate/shared/database"
	"go-boilerplate/shared/event"
	"go-boilerplate/src/authors/infrastructure/repository"
	"go-boilerplate/src/authors/services/command"
	"go-boilerplate/src/authors/services/query"
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

	repository := repository.NewAuthorsPostgresRepository(db)
	publisher := event.InitPublisher(event.PublisherOptions{
		Topic: "authors",
	})

	return Application{
		Commands: Commands{
			CreateAuthor: command.NewCreateAuthorHandler(repository),
			DeleteAuthor: command.NewDeleteAuthorHandler(repository, publisher),
		},
		Queries: Queries{
			GetAuthors: query.NewGetAuthorsHandler(repository),
			GetAuthor:  query.NewGetAuthorHandler(repository),
		},
	}
}
