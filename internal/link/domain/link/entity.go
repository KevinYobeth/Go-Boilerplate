package link

import (
	"time"

	"github.com/google/uuid"
	valueobjects "github.com/kevinyobeth/go-boilerplate/pkg/common/value_objects"
)

type Link struct {
	ID          uuid.UUID
	Slug        string
	URL         string
	Description string
	Total       int `json:"total"`

	valueobjects.AuditAuthor
	valueobjects.AuditTrail
}

type RedirectLink struct {
	ID   uuid.UUID `json:"id"`
	Slug string    `json:"slug"`
	URL  string    `json:"url"`
}

type LinkVisitSnapshot struct {
	ID             uuid.UUID  `json:"id"`
	LinkID         uuid.UUID  `json:"link_id"`
	Total          int        `json:"total"`
	LastSnapshotAt *time.Time `json:"last_snapshot_at"`
}
