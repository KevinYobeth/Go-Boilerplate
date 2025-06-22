package helper

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinyobeth/go-boilerplate/internal/link/domain/link"
	"github.com/kevinyobeth/go-boilerplate/internal/link/infrastructure/repository"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/errors"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/telemetry"
	"github.com/samber/lo"
)

type GetLinkVisitSnapshotRequest struct {
	LinkIDs []uuid.UUID
}

type GetLinkVisitSnapshotOpts struct {
	Params         GetLinkVisitSnapshotRequest
	LinkRepository repository.Repository
}

func GetLinkVisitSnapshot(c context.Context, opts GetLinkVisitSnapshotOpts) (map[uuid.UUID]link.LinkVisitSnapshot, error) {
	ctx, span := telemetry.NewCQHelperSpan(c)
	defer span.End()

	linkSnapshot, err := opts.LinkRepository.GetLinksVisitSnapshot(ctx, opts.Params.LinkIDs)
	if err != nil {
		return nil, errors.NewGenericError(err, "failed to get links visit snapshot")
	}
	linkSnapshotMap := lo.SliceToMap(linkSnapshot, func(snapshot link.LinkVisitSnapshot) (uuid.UUID, link.LinkVisitSnapshot) {
		return snapshot.LinkID, snapshot
	})

	return linkSnapshotMap, nil
}
