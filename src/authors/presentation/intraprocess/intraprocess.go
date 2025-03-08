package intraprocess

import (
	"context"
	"go-boilerplate/src/authors/domain/authors"
	"go-boilerplate/src/authors/services"
	"go-boilerplate/src/authors/services/command"
	"go-boilerplate/src/authors/services/query"
	"go-boilerplate/src/shared/interfaces"

	"github.com/ztrue/tracerr"
)

type AuthorIntraprocessService struct {
	intraprocessInterface services.Application
}

func NewAuthorIntraprocessService(intraprocessInterface services.Application) interfaces.AuthorIntraprocess {
	return AuthorIntraprocessService{intraprocessInterface: intraprocessInterface}
}

func (i AuthorIntraprocessService) GetAuthors(c context.Context, name *string) ([]interfaces.Author, error) {
	authors, err := i.intraprocessInterface.Queries.GetAuthors.Execute(c, query.GetAuthorsParams{
		Name: name,
	})
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return transformDomainAuthorsToIntraprocessAuthors(authors), nil
}

func (i AuthorIntraprocessService) CreateAuthor(c context.Context, name string) (*interfaces.Author, error) {
	author, err := i.intraprocessInterface.Commands.CreateAuthor.Execute(c, command.CreateAuthorParams{
		Name: name,
	})
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return transformDomainAuthorToIntraprocessAuthor(author), err
}

func transformDomainAuthorToIntraprocessAuthor(domainAuthor *authors.Author) *interfaces.Author {
	if domainAuthor == nil {
		return nil
	}

	return &interfaces.Author{
		ID:   domainAuthor.ID,
		Name: domainAuthor.Name,
	}
}

func transformDomainAuthorsToIntraprocessAuthors(domainAuthors []authors.Author) []interfaces.Author {
	var intraprocessAuthors []interfaces.Author

	for _, domainAuthor := range domainAuthors {
		intraprocessAuthors = append(intraprocessAuthors, interfaces.Author{
			ID:   domainAuthor.ID,
			Name: domainAuthor.Name,
		})
	}

	return intraprocessAuthors
}
