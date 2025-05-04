package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinyobeth/go-boilerplate/internal/link/domain/link"
)

type Repository interface {
	CreateLink(c context.Context, request *link.LinkDTO) error
	GetLinks(c context.Context, userID uuid.UUID) ([]link.Link, error)
	GetLinkBySlug(c context.Context, slug string) (*link.RedirectLink, error)
}
