package link

import (
	"time"

	"github.com/google/uuid"
)

type NewVisitCountModel struct {
	LinkID      uuid.UUID
	NewVisits   int
	LatestVisit time.Time
}
