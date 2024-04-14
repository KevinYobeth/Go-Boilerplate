package query

import (
	"context"
	"go-boilerplate/src/authors/domain/authors"
	"go-boilerplate/src/authors/infrastructure/repository"

	"github.com/ztrue/tracerr"
)

type GetAuthorByNameParams struct {
	Name string
}

type GetAuthorByNameHandler struct {
	repository repository.Repository
}

func (h GetAuthorByNameHandler) Execute(c context.Context, params GetAuthorByNameParams) (*authors.Author, error) {
	author, err := h.repository.GetAuthorByName(c, params.Name)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return author, nil
}

func NewGetAuthorByNameHandler(repository repository.Repository) GetAuthorByNameHandler {
	return GetAuthorByNameHandler{repository}
}
