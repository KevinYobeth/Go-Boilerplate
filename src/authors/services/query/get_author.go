package query

import (
	"context"
	"go-boilerplate/shared/decorator"
	"go-boilerplate/shared/errors"
	"go-boilerplate/src/authors/domain/authors"
	"go-boilerplate/src/authors/infrastructure/repository"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type GetAuthorParams struct {
	ID uuid.UUID
}

type getAuthorHandler struct {
	repository repository.Repository
}

type GetAuthorHandler decorator.QueryHandler[GetAuthorParams, *authors.Author]

func (h getAuthorHandler) Handle(c context.Context, params GetAuthorParams) (*authors.Author, error) {
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
