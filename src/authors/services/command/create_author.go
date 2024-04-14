package command

import (
	"context"
	"go-boilerplate/src/authors/domain/authors"
	"go-boilerplate/src/authors/infrastructure/repository"
)

type CreateAuthorParams struct {
	Name string
}

type CreateAuthorHandler struct {
	repository repository.Repository
}

func (h CreateAuthorHandler) Execute(c context.Context, params CreateAuthorParams) (*authors.Author, error) {
	dto := authors.NewCreateAuthorDto(params.Name)

	err := h.repository.CreateAuthor(c, dto)
	if err != nil {
		return nil, err
	}

	return &authors.Author{
		ID:   dto.ID,
		Name: dto.Name,
	}, nil
}

func NewCreateAuthorHandler(database repository.Repository) CreateAuthorHandler {
	return CreateAuthorHandler{database}
}
