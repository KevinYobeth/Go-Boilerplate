package command

import (
	"context"
	"go-boilerplate/src/authors/domain/authors"
	"go-boilerplate/src/authors/infrastructure/repository"
)

type CreateAuthorHandler struct {
	repository repository.Repository
}

func (h CreateAuthorHandler) Execute(c context.Context, request authors.CreateAuthorDto) error {
	dto := authors.NewCreateAuthorDto(request.Name)

	err := h.repository.CreateAuthor(c, dto)
	if err != nil {
		return err
	}

	return nil
}

func NewCreateAuthorHandler(database repository.Repository) CreateAuthorHandler {
	return CreateAuthorHandler{database}
}
