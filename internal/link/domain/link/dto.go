package link

import (
	"github.com/google/uuid"
)

type LinkDTO struct {
	ID          uuid.UUID
	Slug        string
	URL         string
	Description string
	CreatedBy   uuid.UUID
}

func NewLinkDTO(slug, url, description string, createdBy uuid.UUID) *LinkDTO {
	return &LinkDTO{
		ID:          uuid.New(),
		Slug:        slug,
		URL:         url,
		Description: description,
		CreatedBy:   createdBy,
	}
}

func NewUpdateLinkDTO(slug, url, description string, createdBy uuid.UUID) *LinkDTO {
	return &LinkDTO{
		Slug:        slug,
		URL:         url,
		Description: description,
		CreatedBy:   createdBy,
	}
}

type LinkVisitEventDTO struct {
	ID        uuid.UUID
	LinkID    uuid.UUID
	Slug      string
	IPAddress string
	UserAgent string
}

func NewLinkVisitEventDTO(linkID uuid.UUID, slug, ipAddress, userAgent string) *LinkVisitEventDTO {
	return &LinkVisitEventDTO{
		ID:        uuid.New(),
		LinkID:    linkID,
		Slug:      slug,
		IPAddress: ipAddress,
		UserAgent: userAgent,
	}
}

type LinkVisitSnapshotDTO struct {
	ID     uuid.UUID
	LinkID uuid.UUID
	Total  int
}

func NewLinkVisitSnapshotDTO(linkID uuid.UUID) *LinkVisitSnapshotDTO {
	return &LinkVisitSnapshotDTO{
		ID:     uuid.New(),
		Total:  0,
		LinkID: linkID,
	}
}

type UpdateLinkVisitSnapshotDTO struct {
	LinkID    uuid.UUID
	NewVisits int
}

func NewUpdateLinkVisitSnapshotDTO(linkID uuid.UUID, newVisits int) *UpdateLinkVisitSnapshotDTO {
	return &UpdateLinkVisitSnapshotDTO{
		LinkID:    linkID,
		NewVisits: newVisits,
	}
}
