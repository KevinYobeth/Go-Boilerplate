package query

import (
	"context"
	"go-boilerplate/shared/decorator"
	"go-boilerplate/src/books/domain/books"
	"go-boilerplate/src/books/infrastructure/repository"
	"go-boilerplate/src/books/services/helper"

	"github.com/google/uuid"
	"github.com/ztrue/tracerr"
)

type GetBookParams struct {
	ID uuid.UUID
}

type getBookHandler struct {
	repository repository.Repository
}

type GetBookHandler decorator.QueryHandler[GetBookParams, *books.Book]

func (h getBookHandler) Handle(c context.Context, params GetBookParams) (*books.Book, error) {
	book, err := helper.GetBook(c, helper.GetBookOpts{
		Params: helper.GetBookRequest{
			ID: params.ID,
		},
		BookRepository: h.repository,
	})
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return book, nil
}

func NewGetBookHandler(repository repository.Repository) getBookHandler {
	if repository == nil {
		panic("repository is required")
	}

	return getBookHandler{
		repository: repository,
	}
}
