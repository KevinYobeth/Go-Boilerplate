package services

import (
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

func NewAuthorService(repository repository.Repository, publisher event.PublisherInterface) Application {
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
