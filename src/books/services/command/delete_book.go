package command

import (
	"context"
	"go-boilerplate/src/books/infrastructure/repository"

	"github.com/google/uuid"
)

type DeleteBookHandler struct {
	repository repository.Repository
}

func (h DeleteBookHandler) Execute(c context.Context, id uuid.UUID) error {
	return h.repository.DeleteBook(c, id)
}

func NewDeleteBookHandler(repository repository.Repository) DeleteBookHandler {
	return DeleteBookHandler{repository}
}
