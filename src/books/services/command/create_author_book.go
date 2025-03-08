package command

import (
	"context"
	"go-boilerplate/src/books/infrastructure/repository"
	"go-boilerplate/src/books/services/helper"

	"github.com/google/uuid"
	"github.com/ztrue/tracerr"
)

type CreateAuthorBookParams struct {
	BookID   uuid.UUID
	AuthorID uuid.UUID
}

type CreateAuthorBookHandler struct {
	repository repository.Repository
}

func (h CreateAuthorBookHandler) Execute(c context.Context, params CreateAuthorBookParams) error {
	err := helper.CreateAuthorBook(c, helper.CreateAuthorBookOpts{
		Params: helper.CreateAuthorBookRequest{
			BookID:   params.BookID,
			AuthorID: params.AuthorID,
		},
	})
	if err != nil {
		return tracerr.Wrap(err)
	}

	return nil
}

func NewCreateAuthorBookHandler(database repository.Repository) CreateAuthorBookHandler {
	return CreateAuthorBookHandler{database}
}
