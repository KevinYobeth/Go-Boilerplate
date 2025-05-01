package http

import (
	"github.com/kevinyobeth/go-boilerplate/internal/authors/domain/authors"
)

func TransformToHTTPAuthor(authorObj *authors.Author) Author {
	return Author{
		Id:   authorObj.ID,
		Name: authorObj.Name,
	}
}

func TransformToHTTPAuthors(authorsObj []authors.Author) []Author {
	var authors []Author = make([]Author, 0)
	for _, author := range authorsObj {
		authors = append(authors, TransformToHTTPAuthor(&author))
	}
	return authors
}
