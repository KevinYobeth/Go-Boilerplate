package query

import (
	"context"
	"go-boilerplate/src/authors/domain/authors"
	"go-boilerplate/src/authors/infrastructure/repository"

	"github.com/ztrue/tracerr"
)

type GetAuthorsHandler struct {
	repository repository.Repository
}

func (h GetAuthorsHandler) Execute(c context.Context, request authors.GetAuthorsDto) ([]authors.Author, error) {
	authorsObj, err := h.repository.GetAuthors(c, request)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return authorsObj, nil
}

func NewGetAuthorsHandler(repository repository.Repository) GetAuthorsHandler {
	return GetAuthorsHandler{repository}
}
