package query

import (
	"context"

	"github.com/kevinyobeth/go-boilerplate/internal/authors/domain/authors"
	"github.com/kevinyobeth/go-boilerplate/internal/authors/infrastructure/repository"
	"github.com/kevinyobeth/go-boilerplate/shared/decorator"
	"github.com/kevinyobeth/go-boilerplate/shared/errors"
	"github.com/kevinyobeth/go-boilerplate/shared/metrics"

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

func NewGetAuthorHandler(repository repository.Repository, logger *zap.SugaredLogger, metricsClient metrics.Client) GetAuthorHandler {
	if repository == nil {
		panic("repository is required")
	}

	return decorator.ApplyQueryDecorators(
		getAuthorHandler{
			repository: repository,
		}, logger, metricsClient,
	)
}
