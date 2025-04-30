package query

import (
	"context"
	"go-boilerplate/internal/authors/domain/authors"
	"go-boilerplate/internal/authors/infrastructure/repository"
	"go-boilerplate/shared/decorator"
	"go-boilerplate/shared/metrics"

	"github.com/ztrue/tracerr"
	"go.uber.org/zap"
)

type GetAuthorByNameRequest struct {
	Name string
}

type getAuthorByNameHandler struct {
	repository repository.Repository
}

type GetAuthorByNameHandler decorator.QueryHandler[GetAuthorByNameRequest, *authors.Author]

func (h getAuthorByNameHandler) Handle(c context.Context, params GetAuthorByNameRequest) (*authors.Author, error) {
	author, err := h.repository.GetAuthorByName(c, params.Name)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return author, nil
}

func NewGetAuthorByNameHandler(repository repository.Repository, logger *zap.SugaredLogger, metricsClient metrics.Client) GetAuthorByNameHandler {
	if repository == nil {
		panic("repository is required")
	}

	return decorator.ApplyQueryDecorators(
		getAuthorByNameHandler{
			repository: repository,
		}, logger, metricsClient,
	)
}
