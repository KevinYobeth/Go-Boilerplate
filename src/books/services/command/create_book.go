package command

import (
	"context"
	"go-boilerplate/shared/database"
	"go-boilerplate/src/books/domain/authors"
	"go-boilerplate/src/books/domain/books"
	"go-boilerplate/src/books/infrastructure/intraprocess"
	"go-boilerplate/src/books/infrastructure/repository"
	"go-boilerplate/src/books/services/helper"

	"github.com/ztrue/tracerr"
)

type CreateBookParams struct {
	Title  string
	Author string
}

type CreateBookHandler struct {
	manager    database.TransactionManager
	repository repository.Repository
	cache      repository.Cache

	authorService intraprocess.BookAuthorIntraprocess
}

func (h CreateBookHandler) Execute(c context.Context, params CreateBookParams) error {
	return tracerr.Wrap(h.manager.RunInTransaction(c, func(c context.Context) error {
		var authorObj *authors.Author

		authorObj, err := h.authorService.GetAuthorByName(c, params.Author)
		if err != nil {
			return tracerr.Wrap(err)
		}
		if authorObj == nil {
			authorObj, err = h.authorService.CreateAuthor(c, params.Author)
			if err != nil {
				return tracerr.Wrap(err)
			}
		}

		dto := books.NewCreateBookDto(params.Title)

		err = h.repository.CreateBook(c, dto)
		if err != nil {
			return tracerr.Wrap(err)
		}

		err = helper.CreateAuthorBook(c, helper.CreateAuthorBookOpts{
			Params: helper.CreateAuthorBookRequest{
				BookID:   dto.ID,
				AuthorID: authorObj.ID,
			},
		})
		if err != nil {
			return tracerr.Wrap(err)
		}

		return h.cache.ClearBooks(c)
	}))
}

func NewCreateBookHandler(manager database.TransactionManager, database repository.Repository, cache repository.Cache, authorService intraprocess.BookAuthorIntraprocess) CreateBookHandler {
	return CreateBookHandler{manager, database, cache, authorService}
}
