package command

import (
	"context"
	"go-boilerplate/internal/authors/domain/authors"
	"go-boilerplate/internal/authors/infrastructure/repository"
	"go-boilerplate/shared/decorator"

	"github.com/google/uuid"
	"github.com/ztrue/tracerr"
	"go.uber.org/zap"
)

type CreateAuthorRequest struct {
	ID   *uuid.UUID
	Name string
}

type createAuthorHandler struct {
	repository repository.Repository
}

type CreateAuthorHandler decorator.CommandHandler[CreateAuthorRequest]

func (h createAuthorHandler) Handle(c context.Context, params CreateAuthorRequest) error {
	dto := authors.NewCreateAuthorDto(params.Name, params.ID)

	err := h.repository.CreateAuthor(c, dto)
	if err != nil {
		return tracerr.Wrap(err)
	}

	return nil
}

func NewCreateAuthorHandler(repository repository.Repository, logger *zap.SugaredLogger) CreateAuthorHandler {
	if repository == nil {
		panic("repository is required")
	}

	return decorator.ApplyCommandDecorators(
		createAuthorHandler{
			repository: repository,
		}, logger,
	)
}
