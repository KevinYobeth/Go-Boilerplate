package services

import (
	"go-boilerplate/shared/database"
	"go-boilerplate/src/books/domain/authors"
	"go-boilerplate/src/books/infrastructure/repository"
	"go-boilerplate/src/books/services/command"
	"go-boilerplate/src/books/services/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateBook         command.CreateBookHandler
	UpdateBook         command.UpdateBookHandler
	DeleteBook         command.DeleteBookHandler
	DeleteBookByAuthor command.DeleteBookByAuthorHandler

	CreateAuthorBook command.CreateAuthorBookHandler
}

type Queries struct {
	GetBook          query.GetBookHandler
	GetBooks         query.GetBooksHandler
	GetBooksByAuthor query.GetBooksByAuthorHandler
}

func NewBookService(repository repository.Repository, cache repository.Cache, manager database.TransactionManager, authorService authors.AuthorService) Application {
	return Application{
		Commands: Commands{
			CreateBook:         command.NewCreateBookHandler(manager, repository, cache, authorService),
			UpdateBook:         command.NewUpdateBookHandler(repository, cache),
			DeleteBook:         command.NewDeleteBookHandler(manager, repository, cache),
			DeleteBookByAuthor: command.NewDeleteBookByAuthorHandler(manager, repository),
			CreateAuthorBook:   command.NewCreateAuthorBookHandler(repository),
		},
		Queries: Queries{
			GetBooks:         query.NewGetBooksHandler(repository, cache),
			GetBook:          query.NewGetBookHandler(repository),
			GetBooksByAuthor: query.NewGetBooksByAuthorHandler(repository),
		},
	}
}
