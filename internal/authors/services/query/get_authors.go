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

type GetAuthorsRequest struct {
	Name *string
}

type getAuthorsHandler struct {
	repository repository.Repository
}

type GetAuthorsHandler decorator.QueryHandler[GetAuthorsRequest, []authors.Author]

func (h getAuthorsHandler) Handle(c context.Context, params GetAuthorsRequest) ([]authors.Author, error) {
	authorsObj, err := h.repository.GetAuthors(c, authors.GetAuthorsDto{Name: params.Name})
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return authorsObj, nil
}

func NewGetAuthorsHandler(repository repository.Repository, logger *zap.SugaredLogger, metricsClient metrics.Client) GetAuthorsHandler {
	if repository == nil {
		panic("repository is required")
	}

	return decorator.ApplyQueryDecorators(
		getAuthorsHandler{
			repository: repository,
		}, logger, metricsClient,
	)
}
