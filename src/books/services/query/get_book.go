package query

import (
	"context"
	"go-boilerplate/src/books/domain/books"
	"go-boilerplate/src/books/infrastructure/repository"
	"go-boilerplate/src/books/services/helper"

	"github.com/google/uuid"
	"github.com/ztrue/tracerr"
)

type GetBookParams struct {
	ID uuid.UUID
}

type GetBookHandler struct {
	repository repository.Repository
}

func (h GetBookHandler) Execute(c context.Context, params GetBookParams) (*books.Book, error) {
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

func NewGetBookHandler(repository repository.Repository) GetBookHandler {
	return GetBookHandler{repository}
}
