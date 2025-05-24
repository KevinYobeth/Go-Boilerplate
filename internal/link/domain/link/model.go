package link

import (
	"time"

	"github.com/google/uuid"
	valueobjects "github.com/kevinyobeth/go-boilerplate/shared/value_objects"
)

type LinkModel struct {
	ID          uuid.UUID
	Slug        string
	URL         string
	Description string

	valueobjects.AuditAuthor
	valueobjects.AuditTrail
}

type NewVisitCountModel struct {
	LinkID      uuid.UUID
	NewVisits   int
	LatestVisit time.Time
}
