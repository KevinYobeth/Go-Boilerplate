package query

import (
	"context"
	"go-boilerplate/internal/authors/domain/authors"
	"go-boilerplate/internal/authors/infrastructure/repository"
	"go-boilerplate/shared/decorator"
	"go-boilerplate/shared/errors"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type GetAuthorRequest struct {
	ID uuid.UUID
}

type getAuthorHandler struct {
	repository repository.Repository
}

type GetAuthorHandler decorator.QueryHandler[GetAuthorRequest, *authors.Author]

func (h getAuthorHandler) Handle(c context.Context, params GetAuthorRequest) (*authors.Author, error) {
	author, err := h.repository.GetAuthor(c, params.ID)
	if err != nil {
		return nil, err
	}
	if author == nil {
		return nil, errors.NewNotFoundError(nil, "author")
	}

	return author, nil
}

func NewGetAuthorHandler(repository repository.Repository, logger *zap.SugaredLogger) GetAuthorHandler {
	if repository == nil {
		panic("repository is required")
	}

	return decorator.ApplyQueryDecorators(
		getAuthorHandler{
			repository: repository,
		}, logger,
	)
}
