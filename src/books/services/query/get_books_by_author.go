package query

import (
	"context"
	"go-boilerplate/src/books/domain/books"
	"go-boilerplate/src/books/infrastructure/repository"

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
	booksObj, err := h.repository.GetBooksByAuthor(c, params.ID)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	if booksObj == nil {
		return []books.Book{}, nil
	}

	return booksObj, nil
}

func NewGetBooksByAuthorHandler(repository repository.Repository) GetBooksByAuthorHandler {
	return GetBooksByAuthorHandler{repository}
}
