package helper

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinyobeth/go-boilerplate/internal/link/domain/link"
	"github.com/kevinyobeth/go-boilerplate/internal/link/infrastructure/repository"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/errors"
)

type CreateLinkVisitSnapshotRequest struct {
	ID uuid.UUID
}

type CreateLinkVisitSnapshotOpts struct {
	Params         CreateLinkVisitSnapshotRequest
	LinkRepository repository.Repository
}

func CreateLinkVisitSnapshot(c context.Context, opts CreateLinkVisitSnapshotOpts) error {
	dto := link.NewLinkVisitSnapshotDTO(opts.Params.ID)

	err := opts.LinkRepository.CreateLinkVisitSnapshot(c, dto)
	if err != nil {
		return errors.NewGenericError(err, "failed to create link visit snapshot")
	}

	return nil
}
