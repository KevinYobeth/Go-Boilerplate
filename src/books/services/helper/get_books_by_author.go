package helper

import (
	"context"
	"go-boilerplate/shared/telemetry"
	"go-boilerplate/src/books/domain/books"
	"go-boilerplate/src/books/infrastructure/repository"

	"github.com/google/uuid"
	"github.com/ztrue/tracerr"
	"go.opentelemetry.io/otel/trace"
)

type GetBooksByAuthorRequest struct {
	ID uuid.UUID
}

type GetBooksByAuthorOpts struct {
	Params         GetBooksByAuthorRequest
	BookRepository repository.Repository
}

func GetBooksByAuthor(c context.Context, opts GetBooksByAuthorOpts) ([]books.Book, error) {
	ctx, span := telemetry.NewCQHelperSpan(c)
	defer span.End()

	booksObj, err := opts.BookRepository.GetBooksByAuthor(ctx, opts.Params.ID)
	if err != nil {
		span.RecordError(err, trace.WithStackTrace(true))
		return nil, tracerr.Wrap(err)
	}
	if booksObj == nil {
		span.RecordError(err, trace.WithStackTrace(true))
		return []books.Book{}, nil
	}

	return booksObj, nil
}
