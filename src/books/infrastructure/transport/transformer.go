package transport

import "go-boilerplate/src/books/domain/books"

func TransformToHTTPBook(bookObj *books.Book) Book {
	return Book{
		Id:   bookObj.ID,
		Name: bookObj.Title,
	}
}

func TransformToHTTPBooks(booksObj []books.Book) []Book {
	var books []Book
	for _, book := range booksObj {
		books = append(books, TransformToHTTPBook(&book))
	}
	return books
}
