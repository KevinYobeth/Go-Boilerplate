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
	Slug     string
	Metadata LinkVisitEventMetadata
}

type LinkVisitEventMetadata struct {
	IPAddress   string
	UserAgent   string
	Referer     string
	CountryCode string
	DeviceType  string
	Browser     string
}

type getRedirectLinkHandler struct {
	repository repository.Repository
	logger     *zap.SugaredLogger
}

type GetRedirectLinkHandler decorator.QueryHandler[*GetRedirectLinkRequest, *link.RedirectLink]

func (h getRedirectLinkHandler) Handle(c context.Context, params *GetRedirectLinkRequest) (*link.RedirectLink, error) {
	redirectLink, err := h.repository.GetLinkBySlug(c, params.Slug)
	if err != nil {
		return nil, errors.NewGenericError(err, "failed to get link")
	}

	dto := link.NewLinkVisitEventDTO(redirectLink.ID,
		params.Metadata.IPAddress,
		params.Metadata.UserAgent,
		params.Metadata.Referer,
		params.Metadata.CountryCode,
		params.Metadata.DeviceType,
		params.Metadata.Browser)
	err = h.repository.CreateLinkVisitEvent(c, dto)
	if err != nil {
		h.logger.Error("failed to create link visit event", zap.Error(err))
	}

	return redirectLink, nil
}

func NewGetRedirectLinkHandler(repository repository.Repository, logger *zap.SugaredLogger, metricsClient metrics.Client) GetRedirectLinkHandler {
	if repository == nil {
		panic("repository cannot be nil")
	}

	return decorator.ApplyQueryDecorators(
		getRedirectLinkHandler{
			repository: repository,
			logger:     logger,
		}, logger, metricsClient,
	)
}
