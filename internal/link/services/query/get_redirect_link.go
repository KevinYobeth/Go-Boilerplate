package query

import (
	"context"

	"github.com/kevinyobeth/go-boilerplate/internal/link/domain/link"
	"github.com/kevinyobeth/go-boilerplate/internal/link/infrastructure/repository"
	"github.com/kevinyobeth/go-boilerplate/shared/decorator"
	"github.com/kevinyobeth/go-boilerplate/shared/errors"
	"github.com/kevinyobeth/go-boilerplate/shared/metrics"
	"go.uber.org/zap"
)

type GetRedirectLinkRequest struct {
	Slug string
}

type getRedirectLinkHandler struct {
	repository repository.Repository
}

type GetRedirectLinkHandler decorator.QueryHandler[*GetRedirectLinkRequest, *link.RedirectLink]

func (h getRedirectLinkHandler) Handle(c context.Context, params *GetRedirectLinkRequest) (*link.RedirectLink, error) {
	link, err := h.repository.GetLinkBySlug(c, params.Slug)
	if err != nil {
		return nil, errors.NewGenericError(err, "failed to get link")
	}

	return link, nil
}

func NewGetRedirectLinkHandler(repository repository.Repository, logger *zap.SugaredLogger, metricsClient metrics.Client) GetRedirectLinkHandler {
	return decorator.ApplyQueryDecorators(
		getRedirectLinkHandler{
			repository: repository,
		}, logger, metricsClient,
	)
}
