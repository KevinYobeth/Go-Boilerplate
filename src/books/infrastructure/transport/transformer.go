package transport

import "go-boilerplate/src/books/domain/books"

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
