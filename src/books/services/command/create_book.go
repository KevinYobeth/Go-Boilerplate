package command

import (
	"context"
	"go-boilerplate/src/books/domain/books"
	"go-boilerplate/src/books/infrastructure/repository"
)

type CreateBookHandler struct {
	repository repository.Repository
}

func (h CreateBookHandler) Execute(c context.Context, request books.CreateBookDto) error {
	dto := books.NewCreateBookDto(request.Title)

	return h.repository.CreateBook(c, dto)
}

func NewCreateBookHandler(repository repository.Repository) CreateBookHandler {
	return CreateBookHandler{repository}
}
