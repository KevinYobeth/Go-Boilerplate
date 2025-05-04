package link

import (
	"github.com/google/uuid"
	valueobjects "github.com/kevinyobeth/go-boilerplate/shared/value_objects"
)

type Link struct {
	ID          uuid.UUID `json:"id"`
	Slug        string    `json:"slug"`
	URL         string    `json:"url"`
	Description string    `json:"description"`

	valueobjects.AuditAuthor
	valueobjects.AuditTrail
}

type RedirectLink struct {
	ID   uuid.UUID `json:"id"`
	Slug string    `json:"slug"`
	URL  string    `json:"url"`
}
