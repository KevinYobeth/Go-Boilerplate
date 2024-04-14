package query

import (
	"context"
	"go-boilerplate/src/authors/domain/authors"
	"go-boilerplate/src/authors/infrastructure/repository"

	"github.com/google/uuid"
)

type GetAuthorParams struct {
	ID uuid.UUID
}

type GetAuthorHandler struct {
	repository repository.Repository
}

func (h GetAuthorHandler) Execute(c context.Context, params GetAuthorParams) (*authors.Author, error) {
	author, err := h.repository.GetAuthor(c, params.ID)
	if err != nil {
		return nil, err
	}

	return author, nil
}

func NewGetAuthorHandler(repository repository.Repository) GetAuthorHandler {
	return GetAuthorHandler{repository}
}
