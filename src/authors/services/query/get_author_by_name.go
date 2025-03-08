package query

import (
	"context"
	"go-boilerplate/shared/decorator"
	"go-boilerplate/src/authors/domain/authors"
	"go-boilerplate/src/authors/infrastructure/repository"

	"github.com/ztrue/tracerr"
)

type GetAuthorByNameParams struct {
	Name string
}

type getAuthorByNameHandler struct {
	repository repository.Repository
}

type GetAuthorByNameHandler decorator.QueryHandler[GetAuthorByNameParams, *authors.Author]

func (h getAuthorByNameHandler) Handle(c context.Context, params GetAuthorByNameParams) (*authors.Author, error) {
	author, err := h.repository.GetAuthorByName(c, params.Name)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return author, nil
}

func NewGetAuthorByNameHandler(repository repository.Repository) GetAuthorByNameHandler {
	if repository == nil {
		panic("repository is required")
	}

	return decorator.ApplyQueryDecorators(
		getAuthorByNameHandler{
			repository: repository,
		},
	)
}
