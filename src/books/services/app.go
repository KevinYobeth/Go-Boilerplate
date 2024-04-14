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
	UpdateBook command.UpdateBookHandler
	DeleteBook command.DeleteBookHandler
}

type Queries struct {
	GetBook  query.GetBookHandler
	GetBooks query.GetBooksHandler
}

func NewBookService() Application {
	repository := repository.NewBooksPostgresRepository()

	return Application{
		Commands: Commands{
			CreateBook: command.NewCreateBookHandler(repository),
			UpdateBook: command.NewUpdateBookHandler(repository),
			DeleteBook: command.NewDeleteBookHandler(repository),
		},
		Queries: Queries{
			GetBooks: query.NewGetBooksHandler(repository),
			GetBook:  query.NewGetBookHandler(repository),
		},
	}
}
