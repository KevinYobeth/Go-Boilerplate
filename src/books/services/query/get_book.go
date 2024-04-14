package query

import (
	"context"
	"go-boilerplate/shared/errors"
	"go-boilerplate/src/books/domain/books"
	"go-boilerplate/src/books/infrastructure/repository"

	"github.com/google/uuid"
)

type GetBookHandler struct {
	repository repository.Repository
}

func (h GetBookHandler) Execute(c context.Context, id uuid.UUID) (*books.Book, error) {
	book, err := h.repository.GetBook(c, id)
	if err != nil {
		return nil, errors.NewGenericError(err, "failed to get book")
	}

	if book == nil {
		return nil, errors.NewNotFoundError(nil, "book")

	}

	return book, nil
}

func NewGetBookHandler(repository repository.Repository) GetBookHandler {
	return GetBookHandler{repository}
}
