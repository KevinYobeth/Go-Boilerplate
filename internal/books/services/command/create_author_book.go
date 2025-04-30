package command

import (
	"context"
	"go-boilerplate/internal/books/infrastructure/repository"
	"go-boilerplate/internal/books/services/helper"
	"go-boilerplate/shared/decorator"
	"go-boilerplate/shared/metrics"

	"github.com/google/uuid"
	"github.com/ztrue/tracerr"
	"go.uber.org/zap"
)

type CreateAuthorBookRequest struct {
	BookID   uuid.UUID
	AuthorID uuid.UUID
}

type createAuthorBookHandler struct {
	repository repository.Repository
}

type CreateAuthorBookHandler decorator.CommandHandler[CreateAuthorBookRequest]

func (h createAuthorBookHandler) Handle(c context.Context, params CreateAuthorBookRequest) error {
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

func NewCreateAuthorBookHandler(repository repository.Repository, logger *zap.SugaredLogger, metricsClient metrics.Client) CreateAuthorBookHandler {
	if repository == nil {
		panic("nil repository")
	}

	return decorator.ApplyCommandDecorators(
		createAuthorBookHandler{
			repository: repository,
		}, logger, metricsClient,
	)
}
