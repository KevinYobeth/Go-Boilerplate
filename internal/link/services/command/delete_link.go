package command

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinyobeth/go-boilerplate/internal/link/infrastructure/repository"
	"github.com/kevinyobeth/go-boilerplate/internal/link/services/helper"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/decorator"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/errors"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/metrics"
	"github.com/ztrue/tracerr"
	"go.uber.org/zap"
)

type DeleteLinkRequest struct {
	UserID uuid.UUID
	ID     uuid.UUID
}

type deleteLinkHandler struct {
	repository repository.Repository
}

type DeleteLinkHandler decorator.CommandHandler[*DeleteLinkRequest]

func (h deleteLinkHandler) Handle(c context.Context, params *DeleteLinkRequest) error {
	link, err := helper.GetLink(c, helper.GetLinkOpts{
		Params: helper.GetLinkRequest{
			UserID: params.UserID,
			ID:     params.ID,
		},
		LinkRepository: h.repository,
	})
	if err != nil {
		return tracerr.Wrap(err)
	}

	err = h.repository.DeleteLink(c, link.ID)
	if err != nil {
		return errors.NewGenericError(err, "failed to delete link")
	}

	return nil
}

func NewDeleteLinkHandler(repository repository.Repository, logger *zap.SugaredLogger, metricsClient metrics.Client) DeleteLinkHandler {
	if repository == nil {
		panic("repository cannot be nil")
	}

	return decorator.ApplyCommandDecorators(
		deleteLinkHandler{
			repository: repository,
		}, logger, metricsClient,
	)
}
