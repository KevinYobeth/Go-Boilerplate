package command

import (
	"go-boilerplate/src/books/infrastructure/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type DeleteBookHandler struct {
	repository repository.Repository
}

func (h DeleteBookHandler) Execute(c *fiber.Ctx, id uuid.UUID) error {
	return h.repository.DeleteBook(c, id)
}

func NewDeleteBookHandler(repository repository.Repository) DeleteBookHandler {
	return DeleteBookHandler{repository}
}
