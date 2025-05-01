package helper

import (
	"context"

	"github.com/kevinyobeth/go-boilerplate/internal/books/domain/books"
	"github.com/kevinyobeth/go-boilerplate/internal/books/infrastructure/repository"
	"github.com/kevinyobeth/go-boilerplate/shared/errors"
	"github.com/kevinyobeth/go-boilerplate/shared/telemetry"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"
)

type GetBookRequest struct {
	ID uuid.UUID
}

type GetBookOpts struct {
	Params         GetBookRequest
	BookRepository repository.Repository
}

func GetBook(c context.Context, opts GetBookOpts) (*books.Book, error) {
	ctx, span := telemetry.NewCQHelperSpan(c)
	defer span.End()

	book, err := opts.BookRepository.GetBook(ctx, opts.Params.ID)
	if err != nil {
		span.RecordError(err, trace.WithStackTrace(true))
		return nil, errors.NewGenericError(err, "failed to get book")
	}

	if book == nil {
		span.RecordError(err, trace.WithStackTrace(true))
		return nil, errors.NewNotFoundError(nil, "book")
	}

	return book, nil
}
