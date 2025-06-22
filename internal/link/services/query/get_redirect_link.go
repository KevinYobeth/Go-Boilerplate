package query

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinyobeth/go-boilerplate/internal/link/domain/link"
	"github.com/kevinyobeth/go-boilerplate/internal/link/infrastructure/repository"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/decorator"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/errors"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/metrics"
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
	cache      repository.Cache
	logger     *zap.SugaredLogger
}

type GetRedirectLinkHandler decorator.QueryHandler[*GetRedirectLinkRequest, *link.RedirectLink]

func (h getRedirectLinkHandler) Handle(c context.Context, params *GetRedirectLinkRequest) (*link.RedirectLink, error) {
	redirectLink, err := h.cache.GetRedirectLink(c, params.Slug)
	if err != nil {
		h.logger.Error("failed to get redirect link from cache", zap.Error(err))
	}
	if redirectLink == nil {
		redirectLink, err = h.repository.GetLinkBySlug(c, params.Slug)
		if err != nil {
			return nil, errors.NewGenericError(err, "failed to get link")
		}
		if redirectLink == nil {
			return nil, errors.NewNotFoundError(nil, "redirect link")
		}
	}

	h.logVisit(c, redirectLink.ID, params.Slug, params.Metadata.IPAddress, params.Metadata.UserAgent)
	return redirectLink, nil
}

func NewGetRedirectLinkHandler(repository repository.Repository, cache repository.Cache, logger *zap.SugaredLogger, metricsClient metrics.Client) GetRedirectLinkHandler {
	if repository == nil {
		panic("repository cannot be nil")
	}
	if cache == nil {
		panic("cache cannot be nil")
	}

	return decorator.ApplyQueryDecorators(
		getRedirectLinkHandler{
			repository: repository,
			cache:      cache,
			logger:     logger,
		}, logger, metricsClient,
	)
}

func (h getRedirectLinkHandler) logVisit(c context.Context, linkID uuid.UUID, slug, ip, ua string) {
	dto := link.NewLinkVisitEventDTO(linkID, slug, ip, ua)
	if err := h.repository.CreateLinkVisit(c, dto); err != nil {
		h.logger.Error("failed to create link visit event", zap.Error(err))
	}
}
