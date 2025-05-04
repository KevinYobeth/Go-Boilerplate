package link

import "github.com/google/uuid"

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
