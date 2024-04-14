package command

import (
	"context"
	"go-boilerplate/src/books/domain/books"
	"go-boilerplate/src/books/infrastructure/repository"
)

type UpdateBookHandler struct {
	repository repository.Repository
}

func (h UpdateBookHandler) Execute(c context.Context, request books.UpdateBookDto) error {
	return h.repository.UpdateBook(c, request)
}

func NewUpdateBookHandler(repository repository.Repository) UpdateBookHandler {
	return UpdateBookHandler{repository}
}
