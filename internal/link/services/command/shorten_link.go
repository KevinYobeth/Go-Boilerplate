package command

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinyobeth/go-boilerplate/internal/link/domain/link"
	"github.com/kevinyobeth/go-boilerplate/internal/link/infrastructure/repository"
	"github.com/kevinyobeth/go-boilerplate/internal/link/services/helper"
	"github.com/kevinyobeth/go-boilerplate/shared/decorator"
	"github.com/kevinyobeth/go-boilerplate/shared/errors"
	"github.com/kevinyobeth/go-boilerplate/shared/metrics"
	"github.com/kevinyobeth/go-boilerplate/shared/validator"
	"github.com/ztrue/tracerr"
	"go.uber.org/zap"
)

type ShortenLinkRequest struct {
	Slug        string    `conform:"trim" validate:"required,min=3,max=255"`
	URL         string    `conform:"trim" validate:"required,url,min=3,max=255"`
	Description string    `conform:"trim" validate:"required,min=3,max=255"`
	UserID      uuid.UUID `validate:"required,uuid4"`
}

type shortenLinkHandler struct {
	repository repository.Repository
	cache      repository.Cache
	logger     *zap.SugaredLogger
}

type ShortenLinkHandler decorator.CommandHandler[*ShortenLinkRequest]

func (h shortenLinkHandler) Handle(c context.Context, params *ShortenLinkRequest) error {
	if err := validator.ValidateStruct(params); err != nil {
		return tracerr.Wrap(err)
	}

	dbLink, err := helper.GetLinkBySlug(c, helper.GetLinkBySlugOpts{
		Params: helper.GetLinkBySlugRequest{
			Slug: params.Slug,
		},
		SilentNotFound: true,
		LinkRepository: h.repository,
	})
	if err != nil {
		return tracerr.Wrap(err)
	}
	if dbLink != nil {
		return errors.NewGenericError(nil, "slug already exists")
	}

	dto := link.NewLinkDTO(
		params.Slug,
		params.URL,
		params.Description,
		params.UserID,
	)
	err = h.repository.CreateLink(c, dto)
	if err != nil {
		return errors.NewGenericError(err, "failed to create link")
	}

	err = h.cache.SetRedirectLink(c, dto.Slug, link.RedirectLink{
		ID:   dto.ID,
		Slug: dto.Slug,
		URL:  dto.URL,
	})
	if err != nil {
		h.logger.Errorw("failed to set redirect link in cache", "slug", dto.Slug, "error", err)
	}

	err = helper.CreateLinkVisitSnapshot(c, helper.CreateLinkVisitSnapshotOpts{
		Params: helper.CreateLinkVisitSnapshotRequest{
			ID: dto.ID,
		},
		LinkRepository: h.repository,
	})
	if err != nil {
		return tracerr.Wrap(err)
	}

	return nil
}

func NewShortenLinkHandler(repository repository.Repository, cache repository.Cache, logger *zap.SugaredLogger, metricsClient metrics.Client) ShortenLinkHandler {
	if repository == nil {
		panic("repository is required")
	}

	if cache == nil {
		panic("cache is required")
	}

	return decorator.ApplyCommandDecorators(
		shortenLinkHandler{
			repository: repository,
			cache:      cache,
			logger:     logger,
		}, logger, metricsClient,
	)
}
