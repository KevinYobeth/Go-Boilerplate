package query

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinyobeth/go-boilerplate/internal/link/domain/link"
	"github.com/kevinyobeth/go-boilerplate/internal/link/infrastructure/repository"
	"github.com/kevinyobeth/go-boilerplate/internal/link/services/helper"
	"github.com/kevinyobeth/go-boilerplate/shared/builder/pagination"
	"github.com/kevinyobeth/go-boilerplate/shared/decorator"
	"github.com/kevinyobeth/go-boilerplate/shared/errors"
	"github.com/kevinyobeth/go-boilerplate/shared/metrics"
	"github.com/samber/lo"
	"github.com/ztrue/tracerr"
	"go.uber.org/zap"
)

type GetLinksRequest struct {
	UserID uuid.UUID
	Next   uuid.UUID
	Prev   uuid.UUID
	Page   int
	Limit  int
}

type getLinksHandler struct {
	repository repository.Repository
}

type GetLinksHandler decorator.QueryHandler[*GetLinksRequest, *pagination.Collection[link.Link]]

func (h getLinksHandler) Handle(c context.Context, params *GetLinksRequest) (*pagination.Collection[link.Link], error) {
	paginationConfig := pagination.NewLimitPagination(pagination.LimitPaginationRequest[link.LinkModel]{
		Page:  params.Page,
		Limit: params.Limit,
	})

	// paginationConfig := pagination.NewCursorPagination(pagination.CursorPaginationRequest[link.LinkModel]{
	// 	Limit: params.Limit,
	// 	Next:  params.Next,
	// 	Prev:  params.Prev,
	// })

	collection, err := h.repository.GetLinksPaginated(c, params.UserID, paginationConfig)
	if err != nil {
		return nil, errors.NewGenericError(err, "failed to get links")
	}

	linkIDs := lo.Map(collection.Data, func(link link.LinkModel, _ int) uuid.UUID {
		return link.ID
	})
	linkSnapshotMap, err := helper.GetLinkVisitSnapshot(c, helper.GetLinkVisitSnapshotOpts{
		Params: helper.GetLinkVisitSnapshotRequest{
			LinkIDs: linkIDs,
		},
		LinkRepository: h.repository,
	})
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	linksResult := lo.Map(collection.Data, func(model link.LinkModel, _ int) link.Link {
		total := 0
		snapshot, ok := linkSnapshotMap[model.ID]
		if ok {
			total = snapshot.Total
		}

		return link.Link{
			ID:          model.ID,
			Slug:        model.Slug,
			URL:         model.URL,
			Description: model.Description,
			Total:       total,
			AuditAuthor: model.AuditAuthor,
			AuditTrail:  model.AuditTrail,
		}
	})

	return &pagination.Collection[link.Link]{
		Data:     linksResult,
		Metadata: collection.Metadata,
	}, nil
}

func NewGetLinksHandler(repository repository.Repository, logger *zap.SugaredLogger, metricsClient metrics.Client) GetLinksHandler {
	return decorator.ApplyQueryDecorators(
		getLinksHandler{
			repository: repository,
		}, logger, metricsClient,
	)
}
