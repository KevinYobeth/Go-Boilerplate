package command

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinyobeth/go-boilerplate/internal/link/domain/link"
	"github.com/kevinyobeth/go-boilerplate/internal/link/infrastructure/repository"
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
}

type ShortenLinkHandler decorator.CommandHandler[*ShortenLinkRequest]

func (h shortenLinkHandler) Handle(c context.Context, params *ShortenLinkRequest) error {
	if err := validator.ValidateStruct(params); err != nil {
		return tracerr.Wrap(err)
	}

	dto := link.NewLinkDTO(
		params.Slug,
		params.URL,
		params.Description,
		params.UserID,
	)
	err := h.repository.CreateLink(c, dto)
	if err != nil {
		return errors.NewGenericError(err, "failed to create link")
	}

	return nil
}

func NewShortenLinkHandler(repository repository.Repository, logger *zap.SugaredLogger, metricsClient metrics.Client) ShortenLinkHandler {
	if repository == nil {
		panic("repository is required")
	}

	return decorator.ApplyCommandDecorators(
		shortenLinkHandler{
			repository: repository,
		}, logger, metricsClient,
	)
}
