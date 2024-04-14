package command

import (
	"context"
	"go-boilerplate/src/books/domain/books"
	"go-boilerplate/src/books/infrastructure/repository"

	"github.com/google/uuid"
)

type CreateAuthorBookParams struct {
	BookID   uuid.UUID
	AuthorID uuid.UUID
}

type CreateAuthorBookHandler struct {
	repository repository.Repository
}

func (h CreateAuthorBookHandler) Execute(c context.Context, params CreateAuthorBookParams) error {
	return h.repository.CreateAuthorBook(c, books.CreateAuthorBookDto{BookID: params.BookID, AuthorID: params.AuthorID})
}

func NewCreateAuthorBookHandler(database repository.Repository) CreateAuthorBookHandler {
	return CreateAuthorBookHandler{database}
}
