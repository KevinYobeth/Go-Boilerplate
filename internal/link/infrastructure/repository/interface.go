package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinyobeth/go-boilerplate/internal/link/domain/link"
)

type Repository interface {
	CreateLink(c context.Context, request *link.LinkDTO) error
	GetLinks(c context.Context, userID uuid.UUID) ([]link.LinkModel, error)
	GetLink(c context.Context, id, userID uuid.UUID) (*link.Link, error)
	GetLinkBySlug(c context.Context, slug string) (*link.RedirectLink, error)
	DeleteLink(c context.Context, id uuid.UUID) error
	UpdateLink(c context.Context, id uuid.UUID, request *link.LinkDTO) error

	GetNewVisitsCount(c context.Context) ([]link.NewVisitCountModel, error)
	GetLinksVisitSnapshot(c context.Context, linkIDs []uuid.UUID) ([]link.LinkVisitSnapshot, error)
	UpdateLinkVisitSnapshot(c context.Context, dto []link.UpdateLinkVisitSnapshotDTO) error
	CreateLinkVisit(c context.Context, dto *link.LinkVisitEventDTO) error
	CreateLinkVisitSnapshot(c context.Context, dto *link.LinkVisitSnapshotDTO) error
}
