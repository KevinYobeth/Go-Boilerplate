package link

import (
	"time"

	"github.com/google/uuid"
	valueobjects "github.com/kevinyobeth/go-boilerplate/shared/value_objects"
)

type LinkModel struct {
	ID          uuid.UUID `db:"id"`
	Slug        string    `db:"slug"`
	URL         string    `db:"url"`
	Description string    `db:"description"`

	valueobjects.AuditAuthor
	valueobjects.AuditTrail
}

func (l LinkModel) UniqueColumn() string {
	return "id"
}

type NewVisitCountModel struct {
	LinkID      uuid.UUID
	NewVisits   int
	LatestVisit time.Time
}
