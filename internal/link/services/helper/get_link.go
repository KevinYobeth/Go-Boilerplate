package helper

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinyobeth/go-boilerplate/internal/link/domain/link"
	"github.com/kevinyobeth/go-boilerplate/internal/link/infrastructure/repository"
	"github.com/kevinyobeth/go-boilerplate/shared/errors"
	"github.com/kevinyobeth/go-boilerplate/shared/telemetry"
	"go.opentelemetry.io/otel/trace"
)

type GetLinkRequest struct {
	UserID uuid.UUID
	ID     uuid.UUID
}

type GetLinkOpts struct {
	Params         GetLinkRequest
	LinkRepository repository.Repository
}

func GetLink(c context.Context, opts GetLinkOpts) (*link.Link, error) {
	ctx, span := telemetry.NewCQHelperSpan(c)
	defer span.End()

	link, err := opts.LinkRepository.GetLink(ctx, opts.Params.ID, opts.Params.UserID)
	if err != nil {
		span.RecordError(err, trace.WithStackTrace(true))
		return nil, errors.NewGenericError(err, "failed to get link")
	}

	if link == nil {
		span.RecordError(err, trace.WithStackTrace(true))
		return nil, errors.NewNotFoundError(nil, "link")
	}

	return link, nil
}
