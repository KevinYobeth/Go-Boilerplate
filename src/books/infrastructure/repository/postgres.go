package repository

import (
	"database/sql"
	"fmt"
	"go-boilerplate/src/books/domain/books"

	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

type PostgresBooksRepo struct {
	db *sql.DB
}

func NewBooksPostgresRepository(db *sql.DB) Repository {
	return &PostgresBooksRepo{db}
}

func (r PostgresBooksRepo) GetBooks(c *fiber.Ctx, request books.GetBooksDto) ([]books.Book, error) {
	query, args, err := psql.Select("id", "title").From("books").ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(c.Context(), query, args...)
	if err != nil {
		fmt.Println("ERROR DISINI", err)
		return nil, err
	}
	defer rows.Close()

	var booksResult []books.Book
	for rows.Next() {
		var book books.Book
		rows.Scan(&book.ID, &book.Title)

		booksResult = append(booksResult, book)
	}

	return booksResult, nil
}

func (r PostgresBooksRepo) GetBook(c *fiber.Ctx, id uuid.UUID) (books.Book, error) {
	return books.Book{
		ID:    uuid.MustParse("7edb2649-c637-4b7e-9f12-23ce6f35dd34"),
		Title: "The Hobbit",
	}, nil
}

func (r PostgresBooksRepo) CreateBook(c *fiber.Ctx, request books.CreateBookDto) error {
	fmt.Println("Creating book", request)
	return nil
}

func (r PostgresBooksRepo) UpdateBook(c *fiber.Ctx, request books.UpdateBookDto) error {
	return nil
}

func (r PostgresBooksRepo) DeleteBook(c *fiber.Ctx, id uuid.UUID) error {
	return nil
}
