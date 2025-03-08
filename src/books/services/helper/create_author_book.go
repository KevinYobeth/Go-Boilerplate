package helper

import (
	"context"
	"go-boilerplate/src/books/domain/books"
	"go-boilerplate/src/books/infrastructure/repository"

	"github.com/google/uuid"
	"github.com/ztrue/tracerr"
)

type CreateAuthorBookRequest struct {
	BookID   uuid.UUID
	AuthorID uuid.UUID
}

type CreateAuthorBookOpts struct {
	Params         CreateAuthorBookRequest
	BookRepository repository.Repository
}

func CreateAuthorBook(c context.Context, opts CreateAuthorBookOpts) error {
	err := opts.BookRepository.CreateAuthorBook(c, books.CreateAuthorBookDto{
		BookID:   opts.Params.BookID,
		AuthorID: opts.Params.AuthorID,
	})
	if err != nil {
		return tracerr.Wrap(err)
	}

	return nil
}
