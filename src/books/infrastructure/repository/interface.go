package repository

import (
	"context"
	"go-boilerplate/src/books/domain/books"

	"github.com/google/uuid"
)

type Repository interface {
	GetBooks(c context.Context, request books.GetBooksDto) ([]books.BookWithAuthor, error)
	GetBook(c context.Context, id uuid.UUID) (*books.Book, error)

	CreateBook(c context.Context, request books.CreateBookDto) error
	UpdateBook(c context.Context, request books.UpdateBookDto) error
	DeleteBook(c context.Context, id uuid.UUID) error

	CreateAuthorBook(c context.Context, request books.CreateAuthorBookDto) error
}

type Cache interface {
	GetBooks(c context.Context, request books.GetBooksDto) ([]books.BookWithAuthor, error)
	SetBooks(c context.Context, request books.GetBooksDto, books []books.BookWithAuthor) error
	ClearBooks(c context.Context) error
}
