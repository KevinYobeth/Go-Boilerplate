package books

import (
	"github.com/kevinyobeth/go-boilerplate/internal/books/domain/authors"

	"github.com/google/uuid"
)

type Book struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
}

type BookWithAuthor struct {
	Book
	Author authors.Author
}
