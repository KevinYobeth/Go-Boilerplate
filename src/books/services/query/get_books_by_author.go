package query

import (
	"context"
	"go-boilerplate/src/books/domain/books"
	"go-boilerplate/src/books/infrastructure/repository"
	"go-boilerplate/src/books/services/helper"

	"github.com/google/uuid"
	"github.com/ztrue/tracerr"
)

type GetBooksByAuthorParams struct {
	ID uuid.UUID
}

type GetBooksByAuthorHandler struct {
	repository repository.Repository
}

func (h GetBooksByAuthorHandler) Execute(c context.Context, params GetBooksByAuthorParams) ([]books.Book, error) {
	booksObj, err := helper.GetBooksByAuthor(c, helper.GetBooksByAuthorOpts{
		Params: helper.GetBooksByAuthorRequest{
			ID: params.ID,
		},
		BookRepository: h.repository,
	})
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return booksObj, nil
}

func NewGetBooksByAuthorHandler(repository repository.Repository) GetBooksByAuthorHandler {
	return GetBooksByAuthorHandler{repository}
}
