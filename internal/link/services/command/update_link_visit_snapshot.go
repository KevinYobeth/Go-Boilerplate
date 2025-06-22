package command

import (
	"context"

	"github.com/kevinyobeth/go-boilerplate/internal/link/domain/link"
	"github.com/kevinyobeth/go-boilerplate/internal/link/infrastructure/repository"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/decorator"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/metrics"
	"github.com/ztrue/tracerr"
	"go.uber.org/zap"
)

type UpdateLinkVisitSnapshotRequest struct {
}

type updateLinkVisitSnapshotHandler struct {
	repository repository.Repository
	logger     *zap.SugaredLogger
}

type UpdateLinkVisitSnapshot decorator.CommandHandler[*UpdateLinkVisitSnapshotRequest]

func (h updateLinkVisitSnapshotHandler) Handle(c context.Context, params *UpdateLinkVisitSnapshotRequest) error {
	count, err := h.repository.GetNewVisitsCount(c)
	if err != nil {
		return tracerr.Wrap(err)
	}

	if len(count) == 0 {
		h.logger.Info("No new visits to update")
		return nil
	}

	dto := make([]link.UpdateLinkVisitSnapshotDTO, 0, len(count))
	for _, v := range count {
		dto = append(dto, *link.NewUpdateLinkVisitSnapshotDTO(v.LinkID, v.NewVisits))
	}

	err = h.repository.UpdateLinkVisitSnapshot(c, dto)
	if err != nil {
		return tracerr.Wrap(err)
	}

	h.logger.Infof("Updated %d link visit snapshots", len(dto))
	return nil
}

func NewUpdateLinkVisitSnapshotHandler(repository repository.Repository, logger *zap.SugaredLogger, metricsClient metrics.Client) UpdateLinkVisitSnapshot {
	if repository == nil {
		panic("repository cannot be nil")
	}

	return decorator.ApplyCommandDecorators(
		updateLinkVisitSnapshotHandler{
			repository: repository,
			logger:     logger,
		}, logger, metricsClient,
	)
}
