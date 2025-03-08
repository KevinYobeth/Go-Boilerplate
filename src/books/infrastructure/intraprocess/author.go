package intraprocess

import (
	"context"
	"go-boilerplate/src/books/domain/authors"
	"go-boilerplate/src/shared/interfaces"

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
