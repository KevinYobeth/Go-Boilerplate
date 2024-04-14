package services

import (
	"go-boilerplate/shared/cache"
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
	lru := cache.InitRedis()

	repo := repository.NewBooksPostgresRepository(db)
	cache := repository.NewBooksRedisCache(lru)
	manager := database.NewTransactionManager(db)

	return Application{
		Commands: Commands{
			CreateBook: command.NewCreateBookHandler(repo, cache),
			UpdateBook: command.NewUpdateBookHandler(repo, cache),
			DeleteBook: command.NewDeleteBookHandler(manager, repo, cache,
				command.DeleteBookService{
					GetBook: query.NewGetBookHandler(repo),
				}),
		},
		Queries: Queries{
			GetBooks: query.NewGetBooksHandler(repo, cache),
			GetBook:  query.NewGetBookHandler(repo),
		},
	}
}
