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

func (h CreateAuthorHandler) Execute(c context.Context, params CreateAuthorParams) error {
	dto := authors.NewCreateAuthorDto(params.Name)

	err := h.repository.CreateAuthor(c, dto)
	if err != nil {
		return err
	}

	return nil
}

func NewCreateAuthorHandler(database repository.Repository) CreateAuthorHandler {
	return CreateAuthorHandler{database}
}
