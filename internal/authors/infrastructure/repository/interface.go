package repository

import (
	"context"

	"github.com/kevinyobeth/go-boilerplate/internal/authors/domain/authors"

	"github.com/google/uuid"
)

type Repository interface {
	GetAuthors(c context.Context, request authors.GetAuthorsDto) ([]authors.Author, error)
	GetAuthor(c context.Context, id uuid.UUID) (*authors.Author, error)
	GetAuthorByName(c context.Context, name string) (*authors.Author, error)

	CreateAuthor(c context.Context, request authors.CreateAuthorDto) error
	DeleteAuthor(c context.Context, id uuid.UUID) error
}
