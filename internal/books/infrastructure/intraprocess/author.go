package intraprocess

import (
	"context"

	"github.com/kevinyobeth/go-boilerplate/internal/books/domain/authors"
	"github.com/kevinyobeth/go-boilerplate/internal/shared/interfaces"
	"github.com/kevinyobeth/go-boilerplate/shared/telemetry"

	"github.com/ztrue/tracerr"
)

type BookAuthorIntraprocessService struct {
	intraprocessInterface interfaces.AuthorIntraprocess
}

type BookAuthorIntraprocess interface {
	GetAuthorByName(ctx context.Context, name string) (*authors.Author, error)
	CreateAuthor(ctx context.Context, name string) (*authors.Author, error)
}

func NewBookAuthorIntraprocessService(intraprocessInterface interfaces.AuthorIntraprocess) BookAuthorIntraprocess {
	return BookAuthorIntraprocessService{intraprocessInterface: intraprocessInterface}
}

func (i BookAuthorIntraprocessService) GetAuthorByName(ctx context.Context, name string) (*authors.Author, error) {
	ctx, span := telemetry.NewIntraprocessSpan(ctx)
	defer span.End()

	authors, err := i.intraprocessInterface.GetAuthors(ctx, &name)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	if len(authors) == 0 {
		return nil, nil
	}

	return transformIntraprocessAuthorToDomainAuthor(&authors[0]), nil
}

func (i BookAuthorIntraprocessService) CreateAuthor(ctx context.Context, name string) (*authors.Author, error) {
	ctx, span := telemetry.NewIntraprocessSpan(ctx)
	defer span.End()

	author, err := i.intraprocessInterface.CreateAuthor(ctx, name)
	if err != nil {
		return nil, err
	}

	return transformIntraprocessAuthorToDomainAuthor(author), nil
}

func transformIntraprocessAuthorToDomainAuthor(intraprocessAuthor *interfaces.Author) *authors.Author {
	if intraprocessAuthor == nil {
		return nil
	}

	return &authors.Author{
		ID:   intraprocessAuthor.ID,
		Name: intraprocessAuthor.Name,
	}
}
