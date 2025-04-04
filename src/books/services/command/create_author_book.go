package command

import (
	"context"
	"go-boilerplate/shared/decorator"
	"go-boilerplate/src/books/infrastructure/repository"
	"go-boilerplate/src/books/services/helper"

	"github.com/google/uuid"
	"github.com/ztrue/tracerr"
	"go.uber.org/zap"
)

type CreateAuthorBookParams struct {
	BookID   uuid.UUID
	AuthorID uuid.UUID
}

type createAuthorBookHandler struct {
	repository repository.Repository
}

type CreateAuthorBookHandler decorator.CommandHandler[CreateAuthorBookParams]

func (h createAuthorBookHandler) Handle(c context.Context, params CreateAuthorBookParams) error {
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

func NewCreateAuthorBookHandler(repository repository.Repository, logger *zap.SugaredLogger) CreateAuthorBookHandler {
	if repository == nil {
		panic("nil repository")
	}

	return decorator.ApplyCommandDecorators(
		createAuthorBookHandler{
			repository: repository,
		}, logger,
	)
}
