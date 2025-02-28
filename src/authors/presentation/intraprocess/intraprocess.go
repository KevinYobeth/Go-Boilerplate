package intraprocess

import (
	"context"
	"go-boilerplate/src/authors/domain/authors"
	"go-boilerplate/src/authors/services"
	"go-boilerplate/src/authors/services/command"
	"go-boilerplate/src/authors/services/query"

	"github.com/ztrue/tracerr"
)

type AuthorIntraprocessService struct {
	intraprocessInterface services.Application
}

func NewAuthorIntraprocessService(intraprocessInterface services.Application) AuthorIntraprocessService {
	return AuthorIntraprocessService{intraprocessInterface: intraprocessInterface}
}

func (i AuthorIntraprocessService) GetAuthors(c context.Context, name *string) ([]authors.Author, error) {
	authors, err := i.intraprocessInterface.Queries.GetAuthors.Execute(c, query.GetAuthorsParams{
		Name: name,
	})
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return authors, nil
}

func (i AuthorIntraprocessService) CreateAuthor(c context.Context, name string) (*authors.Author, error) {
	author, err := i.intraprocessInterface.Commands.CreateAuthor.Execute(c, command.CreateAuthorParams{
		Name: name,
	})
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return author, err
}
