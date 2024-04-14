package services

import (
	"go-boilerplate/shared/database"
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
	db := database.InitPostgres()
	repository := repository.NewBooksPostgresRepository(db)
	manager := database.NewTransactionManager(db)

	return Application{
		Commands: Commands{
			CreateBook: command.NewCreateBookHandler(manager, repository),
			UpdateBook: command.NewUpdateBookHandler(repository),
			DeleteBook: command.NewDeleteBookHandler(repository, command.DeleteBookService{
				GetBook: query.NewGetBookHandler(repository),
			}),
		},
		Queries: Queries{
			GetBooks: query.NewGetBooksHandler(repository),
			GetBook:  query.NewGetBookHandler(repository),
		},
	}
}
