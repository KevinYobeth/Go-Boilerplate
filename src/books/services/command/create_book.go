package command

import (
	"go-boilerplate/src/books/domain/books"
	"go-boilerplate/src/books/infrastructure/repository"

	"github.com/gofiber/fiber/v2"
)

type CreateBookHandler struct {
	repository repository.Repository
}

func (h CreateBookHandler) Execute(c *fiber.Ctx, request books.CreateBookDto) error {
	return h.repository.CreateBook(c, request)
}

func NewCreateBookHandler(repository repository.Repository) CreateBookHandler {
	return CreateBookHandler{repository}
}
