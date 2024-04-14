package command

import (
	"context"
	"go-boilerplate/shared/database"
	"go-boilerplate/src/books/domain/books"
	"go-boilerplate/src/books/infrastructure/repository"
)

type CreateBookHandler struct {
	manager    *database.TransactionManager
	repository repository.Repository
}

func (h CreateBookHandler) Execute(c context.Context, request books.CreateBookDto) error {
	dto := books.NewCreateBookDto(request.Title)

	return h.repository.CreateBook(c, dto)
}

func NewCreateBookHandler(manager *database.TransactionManager, repository repository.Repository) CreateBookHandler {
	return CreateBookHandler{manager, repository}
}
