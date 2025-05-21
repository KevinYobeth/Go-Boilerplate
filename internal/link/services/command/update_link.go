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

	dto := link.NewUpdateLinkDTO(
		params.Slug,
		params.URL,
		params.Description,
		params.UserID,
	)
	err = h.repository.UpdateLink(c, dbLink.ID, dto)
	if err != nil {
		return errors.NewGenericError(err, "failed to create link")
	}

	return nil
}

func NewUpdateLinkHandler(repository repository.Repository, logger *zap.SugaredLogger, metricsClient metrics.Client) UpdateLinkHandler {
	if repository == nil {
		panic("repository is required")
	}

	return decorator.ApplyCommandDecorators(
		updateLinkHandler{
			repository: repository,
		}, logger, metricsClient,
	)
}
