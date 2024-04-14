package infrastructure

import books "go-boilerplate/internal/book/domain"

type Repository interface {
	GetBooks() ([]books.Book, error)
}
