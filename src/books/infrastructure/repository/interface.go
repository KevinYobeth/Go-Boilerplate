package repository

import (
	"go-boilerplate/src/books/domain/books"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Repository interface {
	GetBooks(c *fiber.Ctx, request books.GetBooksDto) ([]books.Book, error)
	GetBook(c *fiber.Ctx, id uuid.UUID) (books.Book, error)

	CreateBook(c *fiber.Ctx, request books.CreateBookDto) error
	UpdateBook(c *fiber.Ctx, request books.UpdateBookDto) error
	DeleteBook(c *fiber.Ctx, id uuid.UUID) error
}
