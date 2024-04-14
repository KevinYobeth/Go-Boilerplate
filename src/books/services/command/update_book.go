package command

import (
	"go-boilerplate/src/books/domain/books"
	"go-boilerplate/src/books/infrastructure/repository"

	"github.com/gofiber/fiber/v2"
)

type UpdateBookHandler struct {
	repository repository.Repository
}

func (h UpdateBookHandler) Execute(c *fiber.Ctx, request books.UpdateBookDto) error {
	return h.repository.UpdateBook(c, request)
}

func NewUpdateBookHandler(repository repository.Repository) UpdateBookHandler {
	return UpdateBookHandler{repository}
}
