package query

import (
	"go-boilerplate/src/books/domain/books"
	"go-boilerplate/src/books/infrastructure/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type GetBookHandler struct {
	repository repository.Repository
}

func (h GetBookHandler) Execute(c *fiber.Ctx, id uuid.UUID) (*books.Book, error) {
	book, err := h.repository.GetBook(c, id)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func NewGetBookHandler(repository repository.Repository) GetBookHandler {
	return GetBookHandler{repository}
}
