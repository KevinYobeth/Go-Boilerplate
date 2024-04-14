package repository

import (
	"context"
	"go-boilerplate/src/authors/domain/authors"
)

type Repository interface {
	GetAuthors(c context.Context, request authors.GetAuthorsDto) ([]authors.Author, error)
	GetAuthor(c context.Context, id string) (*authors.Author, error)

	CreateAuthor(c context.Context, request authors.CreateAuthorDto) error
}
