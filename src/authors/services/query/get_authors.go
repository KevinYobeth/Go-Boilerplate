package query

import (
	"context"
	"go-boilerplate/src/authors/domain/authors"
	"go-boilerplate/src/authors/infrastructure/repository"

	"github.com/ztrue/tracerr"
)

type GetAuthorsParams struct {
	Name *string
}

type GetAuthorsHandler struct {
	repository repository.Repository
}

func (h GetAuthorsHandler) Execute(c context.Context, params GetAuthorsParams) ([]authors.Author, error) {
	// event.InitPublisher()

	authorsObj, err := h.repository.GetAuthors(c, authors.GetAuthorsDto{Name: params.Name})
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return authorsObj, nil
}

func NewGetAuthorsHandler(repository repository.Repository) GetAuthorsHandler {
	return GetAuthorsHandler{repository}
}
