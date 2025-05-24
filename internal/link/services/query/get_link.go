package query

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinyobeth/go-boilerplate/internal/link/domain/link"
	"github.com/kevinyobeth/go-boilerplate/internal/link/infrastructure/repository"
	"github.com/kevinyobeth/go-boilerplate/internal/link/services/helper"
	"github.com/kevinyobeth/go-boilerplate/shared/decorator"
	"github.com/kevinyobeth/go-boilerplate/shared/metrics"
	"github.com/ztrue/tracerr"
	"go.uber.org/zap"
)

type GetLinkRequest struct {
	UserID uuid.UUID
	ID     uuid.UUID
}

type getLinkHandler struct {
	repository repository.Repository
}

type GetLinkHandler decorator.QueryHandler[*GetLinkRequest, *link.Link]

func (h getLinkHandler) Handle(c context.Context, params *GetLinkRequest) (*link.Link, error) {
	link, err := helper.GetLink(c, helper.GetLinkOpts{
		Params: helper.GetLinkRequest{
			UserID: params.UserID,
			ID:     params.ID,
		},
		LinkRepository: h.repository,
	})
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	linkSnapshotMap, err := helper.GetLinkVisitSnapshot(c, helper.GetLinkVisitSnapshotOpts{
		Params: helper.GetLinkVisitSnapshotRequest{
			LinkIDs: []uuid.UUID{link.ID},
		},
		LinkRepository: h.repository,
	})
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	snapshot, ok := linkSnapshotMap[link.ID]
	if ok {
		link.Total = snapshot.Total
	}

	return link, nil
}

func NewGetLinkHandler(repository repository.Repository, logger *zap.SugaredLogger, metricsClient metrics.Client) GetLinkHandler {
	if repository == nil {
		panic("repository cannot be nil")
	}

	return decorator.ApplyQueryDecorators(
		getLinkHandler{
			repository: repository,
		}, logger, metricsClient,
	)
}
