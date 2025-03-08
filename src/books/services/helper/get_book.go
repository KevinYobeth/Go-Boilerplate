package helper

import (
	"context"
	"go-boilerplate/shared/errors"
	"go-boilerplate/src/books/domain/books"
	"go-boilerplate/src/books/infrastructure/repository"

	"github.com/google/uuid"
)

type GetBookRequest struct {
	ID uuid.UUID
}

type GetBookOpts struct {
	Params         GetBookRequest
	BookRepository repository.Repository
}

func GetBook(c context.Context, opts GetBookOpts) (*books.Book, error) {
	book, err := opts.BookRepository.GetBook(c, opts.Params.ID)
	if err != nil {
		return nil, errors.NewGenericError(err, "failed to get book")
	}

	if book == nil {
		return nil, errors.NewNotFoundError(nil, "book")

	}

	return book, nil
}
