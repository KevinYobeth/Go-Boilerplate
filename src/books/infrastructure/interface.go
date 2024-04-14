package infrastructure

import "go-boilerplate/src/books/domain/books"

type Repository interface {
	GetBooks() ([]books.Book, error)
}
