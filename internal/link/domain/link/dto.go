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

type LinkVisitEventDTO struct {
	ID          uuid.UUID
	LinkID      uuid.UUID
	IPAddress   string
	UserAgent   string
	Referer     string
	CountryCode string
	DeviceType  string
	Browser     string
}

func NewLinkVisitEventDTO(linkID uuid.UUID, ipAddress, userAgent, referer, countryCode, deviceType, browser string) *LinkVisitEventDTO {

	return &LinkVisitEventDTO{
		ID:          uuid.New(),
		LinkID:      linkID,
		IPAddress:   ipAddress,
		UserAgent:   userAgent,
		Referer:     referer,
		CountryCode: countryCode,
		DeviceType:  deviceType,
		Browser:     browser,
	}
}
