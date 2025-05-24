package helper

import (
	"context"

	"github.com/kevinyobeth/go-boilerplate/internal/link/domain/link"
	"github.com/kevinyobeth/go-boilerplate/internal/link/infrastructure/repository"
	"github.com/kevinyobeth/go-boilerplate/shared/errors"
	"github.com/kevinyobeth/go-boilerplate/shared/telemetry"
	"github.com/ztrue/tracerr"
	"go.opentelemetry.io/otel/trace"
)

type GetLinkBySlugRequest struct {
	Slug string
}

type GetLinkBySlugOpts struct {
	Params         GetLinkBySlugRequest
	SilentNotFound bool
	LinkRepository repository.Repository
}

func GetLinkBySlug(c context.Context, opts GetLinkBySlugOpts) (*link.RedirectLink, error) {
	ctx, span := telemetry.NewCQHelperSpan(c)
	defer span.End()

	link, err := opts.LinkRepository.GetLinkBySlug(ctx, opts.Params.Slug)
	if !opts.SilentNotFound && link == nil {
		span.RecordError(err, trace.WithStackTrace(true))
		return nil, tracerr.Wrap(errors.NewNotFoundError(nil, "link"))
	}

	if err != nil {
		span.RecordError(err, trace.WithStackTrace(true))
		return nil, tracerr.Wrap(errors.NewGenericError(err, "failed to get link"))
	}

	return link, nil
}
