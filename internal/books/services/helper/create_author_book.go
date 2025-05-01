package helper

import (
	"context"

	"github.com/kevinyobeth/go-boilerplate/internal/books/domain/books"
	"github.com/kevinyobeth/go-boilerplate/internal/books/infrastructure/repository"
	"github.com/kevinyobeth/go-boilerplate/shared/telemetry"

	"github.com/google/uuid"
	"github.com/ztrue/tracerr"
	"go.opentelemetry.io/otel/trace"
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
	ctx, span := telemetry.NewCQHelperSpan(c)
	defer span.End()

	err := opts.BookRepository.CreateAuthorBook(ctx, books.CreateAuthorBookDto{
		BookID:   opts.Params.BookID,
		AuthorID: opts.Params.AuthorID,
	})
	if err != nil {
		span.RecordError(err, trace.WithStackTrace(true))
		return tracerr.Wrap(err)
	}

	return nil
}
