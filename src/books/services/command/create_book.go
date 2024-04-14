package command

import (
	"context"
	"go-boilerplate/shared/database"
	"go-boilerplate/src/authors/domain/authors"
	"go-boilerplate/src/authors/services/command"
	authorCommand "go-boilerplate/src/authors/services/command"
	authorQuery "go-boilerplate/src/authors/services/query"
	"go-boilerplate/src/books/domain/books"
	"go-boilerplate/src/books/infrastructure/repository"

	"github.com/ztrue/tracerr"
)

type CreateBookParams struct {
	Title  string
	Author string
}

type AuthorService struct {
	GetAuthorByName authorQuery.GetAuthorByNameHandler
	CreateAuthor    authorCommand.CreateAuthorHandler
}

type CreateBookService struct {
	CreateAuthorBook CreateAuthorBookHandler
}

type CreateBookHandler struct {
	manager       database.TransactionManager
	repository    repository.Repository
	cache         repository.Cache
	service       CreateBookService
	authorService AuthorService
}

func (h CreateBookHandler) Execute(c context.Context, params CreateBookParams) error {
	return tracerr.Wrap(h.manager.RunInTransaction(c, func(c context.Context) error {
		var authorObj *authors.Author

		authorObj, err := h.authorService.GetAuthorByName.Execute(c, authorQuery.GetAuthorByNameParams{Name: params.Author})
		if err != nil {
			return tracerr.Wrap(err)
		}
		if authorObj == nil {
			authorObj, err = h.authorService.CreateAuthor.Execute(c, command.CreateAuthorParams{Name: params.Author})
			if err != nil {
				return tracerr.Wrap(err)
			}
		}

		dto := books.NewCreateBookDto(params.Title)

		err = h.repository.CreateBook(c, dto)
		if err != nil {
			return tracerr.Wrap(err)
		}

		err = h.service.CreateAuthorBook.Execute(c, CreateAuthorBookParams{BookID: dto.ID, AuthorID: authorObj.ID})
		if err != nil {
			return tracerr.Wrap(err)
		}

		return h.cache.ClearBooks(c)
	}))
}

func NewCreateBookHandler(manager database.TransactionManager, database repository.Repository, cache repository.Cache, service CreateBookService, authorService AuthorService) CreateBookHandler {
	return CreateBookHandler{manager, database, cache, service, authorService}
}
