package query

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinyobeth/go-boilerplate/internal/link/domain/link"
	"github.com/kevinyobeth/go-boilerplate/internal/link/infrastructure/repository"
	"github.com/kevinyobeth/go-boilerplate/shared/decorator"
	"github.com/kevinyobeth/go-boilerplate/shared/errors"
	"github.com/kevinyobeth/go-boilerplate/shared/metrics"
	"go.uber.org/zap"
)

type GetLinksRequest struct {
	UserID uuid.UUID
}

type getLinksHandler struct {
	repository repository.Repository
}

type GetLinksHandler decorator.QueryHandler[*GetLinksRequest, []link.Link]

func (h getLinksHandler) Handle(c context.Context, params *GetLinksRequest) ([]link.Link, error) {
	links, err := h.repository.GetLinks(c, params.UserID)
	if err != nil {
		return nil, errors.NewGenericError(err, "failed to get links")
	}

	return links, nil
}

func NewGetLinksHandler(repository repository.Repository, logger *zap.SugaredLogger, metricsClient metrics.Client) GetLinksHandler {
	return decorator.ApplyQueryDecorators(
		getLinksHandler{
			repository: repository,
		}, logger, metricsClient,
	)
}
