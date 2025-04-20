package valueobjects

import (
	"time"
)

type AuditAuthor struct {
	CreatedBy string `json:"createdBy"`
	UpdatedBy *string `json:"updatedBy"`
}

type AuditTrail struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

type SoftDelete struct {
	DeletedAt *time.Time `json:"deletedAt"`
	DeletedBy *string    `json:"deletedBy"`
}

type SoftDeleteTrail struct {
	DeletedAt *time.Time `json:"deletedAt"`
}
