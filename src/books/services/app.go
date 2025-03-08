package services

import (
	"go-boilerplate/shared/cache"
	"go-boilerplate/shared/database"
	"go-boilerplate/src/books/infrastructure/intraprocess"
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

func NewBookService(authorService intraprocess.BookAuthorIntraprocess) Application {
	lru := cache.InitRedis()
	db := database.InitPostgres()
	manager := database.NewTransactionManager(db)

	repo := repository.NewBooksPostgresRepository(db)
	cache := repository.NewBooksRedisCache(lru)

	return Application{
		Commands: Commands{
			CreateBook:         command.NewCreateBookHandler(manager, repo, cache, authorService),
			UpdateBook:         command.NewUpdateBookHandler(repo, cache),
			DeleteBook:         command.NewDeleteBookHandler(manager, repo, cache),
			DeleteBookByAuthor: command.NewDeleteBookByAuthorHandler(manager, repo),
			CreateAuthorBook:   command.NewCreateAuthorBookHandler(repo),
		},
		Queries: Queries{
			GetBooks:         query.NewGetBooksHandler(repo, cache),
			GetBook:          query.NewGetBookHandler(repo),
			GetBooksByAuthor: query.NewGetBooksByAuthorHandler(repo),
		},
	}
}
