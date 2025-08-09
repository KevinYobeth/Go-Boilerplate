package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinyobeth/go-boilerplate/internal/link/domain/link"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/builder/pagination"
)

type Repository interface {
	CreateLink(c context.Context, request *link.LinkDTO) error
	GetLinks(c context.Context, userID uuid.UUID) ([]link.LinkModel, error)
	GetLink(c context.Context, id, userID uuid.UUID) (*link.Link, error)
	GetLinkBySlug(c context.Context, slug string) (*link.RedirectLink, error)
	DeleteLink(c context.Context, id uuid.UUID) error
	UpdateLink(c context.Context, id uuid.UUID, request *link.LinkDTO) error
	GetLinksPaginated(c context.Context, userID uuid.UUID, paginationConfig pagination.Config[link.LinkModel]) (pagination.Collection[link.LinkModel], error)

	GetNewVisitsCount(c context.Context) ([]link.NewVisitCountModel, error)
	GetLinksVisitSnapshot(c context.Context, linkIDs []uuid.UUID) ([]link.LinkVisitSnapshot, error)
	UpdateLinkVisitSnapshot(c context.Context, dto []link.UpdateLinkVisitSnapshotDTO) error
	CreateLinkVisit(c context.Context, dto *link.LinkVisitEventDTO) error
	CreateLinkVisitSnapshot(c context.Context, dto *link.LinkVisitSnapshotDTO) error
}

type Cache interface {
	GetRedirectLink(c context.Context, slug string) (*link.RedirectLink, error)
	SetRedirectLink(c context.Context, slug string, value link.RedirectLink) error
	ClearRedirectLink(c context.Context, slug string) error
}
