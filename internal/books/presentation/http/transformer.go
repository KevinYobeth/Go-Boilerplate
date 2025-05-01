package http

import (
	"github.com/kevinyobeth/go-boilerplate/internal/books/domain/books"
)

func TransformToHTTPBook(bookObj *books.Book) Book {
	return Book{
		Id:    bookObj.ID,
		Title: bookObj.Title,
	}
}

func TransformToHTTPBooks(booksObj []books.Book) []Book {
	var books []Book = make([]Book, 0)
	for _, book := range booksObj {
		books = append(books, TransformToHTTPBook(&book))
	}
	return books
}

func TransformToHTTPBookWithAuthor(bookObj *books.BookWithAuthor) BookWithAuthor {
	return BookWithAuthor{
		Id:    bookObj.ID,
		Title: bookObj.Title,
		Author: &Author{
			Id:   bookObj.Author.ID,
			Name: bookObj.Author.Name,
		},
	}
}

func TransformToHTTPBooksWithAuthor(booksObj []books.BookWithAuthor) []BookWithAuthor {
	var books []BookWithAuthor = make([]BookWithAuthor, 0)
	for _, book := range booksObj {
		books = append(books, TransformToHTTPBookWithAuthor(&book))
	}
	return books
}
