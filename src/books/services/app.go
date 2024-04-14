package services

import (
	"go-boilerplate/src/books/infrastructure/repository"
	"go-boilerplate/src/books/services/command"
	"go-boilerplate/src/books/services/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateBook command.CreateBookHandler
}

type Queries struct {
	GetBooks query.GetBooksHandler
}

func NewBookService() Application {
	repository := repository.NewBooksPostgresRepository()

	return Application{
		Commands: Commands{
			CreateBook: command.NewCreateBookHandler(repository),
		},
		Queries: Queries{
			GetBooks: query.NewGetBooksHandler(repository),
		},
	}
}
