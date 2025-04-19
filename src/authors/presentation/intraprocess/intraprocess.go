package intraprocess

import (
	"context"
	"go-boilerplate/src/authors/domain/authors"
	"go-boilerplate/src/authors/services"
	"go-boilerplate/src/authors/services/command"
	"go-boilerplate/src/authors/services/query"
	"go-boilerplate/src/shared/interfaces"

	"github.com/google/uuid"
	"github.com/ztrue/tracerr"
)

type AuthorIntraprocessService struct {
	intraprocessInterface services.Application
}

func NewAuthorIntraprocessService(intraprocessInterface services.Application) interfaces.AuthorIntraprocess {
	return AuthorIntraprocessService{intraprocessInterface: intraprocessInterface}
}

func (i AuthorIntraprocessService) GetAuthors(c context.Context, name *string) ([]interfaces.Author, error) {
	authors, err := i.intraprocessInterface.Queries.GetAuthors.Handle(c, query.GetAuthorsRequest{
		Name: name,
	})
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return transformDomainAuthorsToIntraprocessAuthors(authors), nil
}

func (i AuthorIntraprocessService) CreateAuthor(c context.Context, name string) (*interfaces.Author, error) {
	ID := uuid.New()
	err := i.intraprocessInterface.Commands.CreateAuthor.Handle(c, command.CreateAuthorRequest{
		ID:   &ID,
		Name: name,
	})
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return transformDomainAuthorToIntraprocessAuthor(&authors.Author{
		ID:   ID,
		Name: name,
	}), err
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
