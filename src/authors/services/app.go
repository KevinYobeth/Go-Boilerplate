package services

import (
	"go-boilerplate/shared/database"
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
}

type Queries struct {
	GetAuthors query.GetAuthorsHandler
	GetAuthor  query.GetAuthorHandler
}

func NewAuthorService() Application {
	db := database.InitPostgres()
	repo := repository.NewAuthorsPostgresRepository(db)

	return Application{
		Commands: Commands{
			CreateAuthor: command.NewCreateAuthorHandler(repo),
		},
		Queries: Queries{
			GetAuthors: query.NewGetAuthorsHandler(repo),
			GetAuthor:  query.NewGetAuthorHandler(repo),
		},
	}
}
