package helper

import (
	"context"
	"go-boilerplate/src/books/domain/books"
	"go-boilerplate/src/books/infrastructure/repository"

	"github.com/google/uuid"
	"github.com/ztrue/tracerr"
)

type GetBooksByAuthorRequest struct {
	ID uuid.UUID
}

type GetBooksByAuthorOpts struct {
	Params         GetBooksByAuthorRequest
	BookRepository repository.Repository
}

func GetBooksByAuthor(c context.Context, opts GetBooksByAuthorOpts) ([]books.Book, error) {
	booksObj, err := opts.BookRepository.GetBooksByAuthor(c, opts.Params.ID)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	if booksObj == nil {
		return []books.Book{}, nil
	}

	return booksObj, nil
}
