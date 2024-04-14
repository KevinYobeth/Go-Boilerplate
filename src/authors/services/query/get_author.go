package query

import (
	"context"
	"go-boilerplate/src/authors/domain/authors"
	"go-boilerplate/src/authors/infrastructure/repository"
)

type GetAuthorHandler struct {
	repository repository.Repository
}

func (h GetAuthorHandler) Execute(c context.Context, id string) (*authors.Author, error) {
	author, err := h.repository.GetAuthor(c, id)
	if err != nil {
		return nil, err
	}

	return author, nil
}

func NewGetAuthorHandler(repository repository.Repository) GetAuthorHandler {
	return GetAuthorHandler{repository}
}
