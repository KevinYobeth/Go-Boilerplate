package repository

import (
	"go-boilerplate/src/books/domain/books"

	"github.com/gofiber/fiber/v2"
)

type Repository interface {
	CreateBook(c *fiber.Ctx, request books.CreateBookDto) error

	GetBooks(c *fiber.Ctx, request books.GetBooksDto) ([]books.Book, error)
}
