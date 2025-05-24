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

type UpdateLinkRequest struct {
	ID          uuid.UUID `validate:"required,uuid4"`
	Slug        string    `conform:"trim" validate:"required,min=3,max=255"`
	URL         string    `conform:"trim" validate:"required,url,min=3,max=255"`
	Description string    `conform:"trim" validate:"required,min=3,max=255"`
	UserID      uuid.UUID `validate:"required,uuid4"`
}

type updateLinkHandler struct {
	repository repository.Repository
	cache      repository.Cache
	logger     *zap.SugaredLogger
}

type UpdateLinkHandler decorator.CommandHandler[*UpdateLinkRequest]

func (h updateLinkHandler) Handle(c context.Context, params *UpdateLinkRequest) error {
	if err := validator.ValidateStruct(params); err != nil {
		return tracerr.Wrap(err)
	}

	dbLink, err := helper.GetLink(c, helper.GetLinkOpts{
		Params: helper.GetLinkRequest{
			UserID: params.UserID,
			ID:     params.ID,
		},
		LinkRepository: h.repository,
	})
	if err != nil {
		return tracerr.Wrap(err)
	}

	if params.Slug != dbLink.Slug {
		link, err := helper.GetLinkBySlug(c, helper.GetLinkBySlugOpts{
			Params: helper.GetLinkBySlugRequest{
				Slug: params.Slug,
			},
			SilentNotFound: true,
			LinkRepository: h.repository,
		})
		if err != nil {
			return tracerr.Wrap(err)
		}

		if link != nil {
			return errors.NewIncorrectInputError(nil, "slug already exists")
		}

		err = h.cache.ClearRedirectLink(c, dbLink.Slug)
		if err != nil {
			h.logger.Errorw("failed to clear redirect link in cache", "slug", dbLink.Slug, "error", err)
		}
	}

	dto := link.NewUpdateLinkDTO(
		params.Slug,
		params.URL,
		params.Description,
		params.UserID,
	)
	err = h.repository.UpdateLink(c, dbLink.ID, dto)
	if err != nil {
		return errors.NewGenericError(err, "failed to update link")
	}

	err = h.cache.SetRedirectLink(c, dto.Slug, link.RedirectLink{
		ID:   dbLink.ID,
		Slug: dto.Slug,
		URL:  dto.URL,
	})
	if err != nil {
		h.logger.Errorw("failed to set redirect link in cache", "slug", dto.Slug, "error", err)
	}

	return nil
}

func NewUpdateLinkHandler(repository repository.Repository, cache repository.Cache, logger *zap.SugaredLogger, metricsClient metrics.Client) UpdateLinkHandler {
	if repository == nil {
		panic("repository is required")
	}

	if cache == nil {
		panic("cache is required")
	}

	return decorator.ApplyCommandDecorators(
		updateLinkHandler{
			repository: repository,
			cache:      cache,
			logger:     logger,
		}, logger, metricsClient,
	)
}
