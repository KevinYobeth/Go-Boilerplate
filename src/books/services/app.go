package services

import (
	"go-boilerplate/shared/cache"
	"go-boilerplate/shared/database"
	authorRepo "go-boilerplate/src/authors/infrastructure/repository"
	authorCommand "go-boilerplate/src/authors/services/command"
	authorQuery "go-boilerplate/src/authors/services/query"
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

func NewBookService() Application {
	db := database.InitPostgres()
	lru := cache.InitRedis()

	repo := repository.NewBooksPostgresRepository(db)
	cache := repository.NewBooksRedisCache(lru)
	manager := database.NewTransactionManager(db)

	authorRepo := authorRepo.NewAuthorsPostgresRepository(db)

	return Application{
		Commands: Commands{
			CreateBook: command.NewCreateBookHandler(manager, repo, cache,
				command.CreateBookService{
					CreateAuthorBook: command.NewCreateAuthorBookHandler(repo),
				},
				command.AuthorService{
					GetAuthorByName: authorQuery.NewGetAuthorByNameHandler(authorRepo),
					CreateAuthor:    authorCommand.NewCreateAuthorHandler(authorRepo),
				}),
			UpdateBook: command.NewUpdateBookHandler(repo, cache),
			DeleteBook: command.NewDeleteBookHandler(manager, repo, cache,
				command.DeleteBookService{
					GetBook: query.NewGetBookHandler(repo),
				}),
			DeleteBookByAuthor: command.NewDeleteBookByAuthorHandler(manager, repo,
				command.DeleteBookByAuthorService{
					GetBooksByAuthor: query.NewGetBooksByAuthorHandler(repo),
				}),

			CreateAuthorBook: command.NewCreateAuthorBookHandler(repo),
		},
		Queries: Queries{
			GetBooks:         query.NewGetBooksHandler(repo, cache),
			GetBook:          query.NewGetBookHandler(repo),
			GetBooksByAuthor: query.NewGetBooksByAuthorHandler(repo),
		},
	}
}
