package services

import (
	"go-boilerplate/internal/books/infrastructure/intraprocess"
	"go-boilerplate/internal/books/infrastructure/repository"
	"go-boilerplate/internal/books/services/command"
	"go-boilerplate/internal/books/services/query"
	"go-boilerplate/shared/cache"
	"go-boilerplate/shared/database"
	"go-boilerplate/shared/log"
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
	logger := log.InitLogger()

	repo := repository.NewBooksPostgresRepository(db)
	cache := repository.NewBooksRedisCache(lru)

	return Application{
		Commands: Commands{
			CreateBook:         command.NewCreateBookHandler(manager, repo, cache, authorService, logger),
			UpdateBook:         command.NewUpdateBookHandler(repo, cache, logger),
			DeleteBook:         command.NewDeleteBookHandler(manager, repo, cache, logger),
			DeleteBookByAuthor: command.NewDeleteBookByAuthorHandler(manager, repo, logger),
			CreateAuthorBook:   command.NewCreateAuthorBookHandler(repo, logger),
		},
		Queries: Queries{
			GetBooks:         query.NewGetBooksHandler(repo, cache, logger),
			GetBook:          query.NewGetBookHandler(repo, logger),
			GetBooksByAuthor: query.NewGetBooksByAuthorHandler(repo, logger),
		},
	}
}
