package repository

import (
	"fmt"
	"go-boilerplate/src/books/domain/books"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PostgresBooksRepo struct {
}

func NewBooksPostgresRepository() Repository {
	return &PostgresBooksRepo{}
}

func (r PostgresBooksRepo) GetBooks(c *fiber.Ctx, request books.GetBooksDto) ([]books.Book, error) {
	return []books.Book{{
		ID:    uuid.MustParse("7edb2649-c637-4b7e-9f12-23ce6f35dd34"),
		Title: "The Hobbit",
	}}, nil
}

func (r PostgresBooksRepo) CreateBook(c *fiber.Ctx, request books.CreateBookDto) error {
	fmt.Println("Creating book", request)
	return nil
}
