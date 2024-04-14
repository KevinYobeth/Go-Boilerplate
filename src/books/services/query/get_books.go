package query

import (
	"go-boilerplate/src/books/domain/books"
	"go-boilerplate/src/books/infrastructure/repository"

	"github.com/gofiber/fiber/v2"
)

type GetBooksHandler struct {
	repository repository.Repository
}

func (h GetBooksHandler) Execute(c *fiber.Ctx, request books.GetBooksDto) ([]books.Book, error) {
	books, err := h.repository.GetBooks(c, request)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func NewGetBooksHandler(repository repository.Repository) GetBooksHandler {
	return GetBooksHandler{repository}
}
