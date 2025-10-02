package valueobjects

import (
	"time"
)

type AuditAuthor struct {
	CreatedBy string  `json:"createdBy" db:"created_by"`
	UpdatedBy *string `json:"updatedBy" db:"updated_by"`
}

type AuditTrail struct {
	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" db:"updated_at"`
}

type SoftDelete struct {
	DeletedAt *time.Time `json:"deletedAt" db:"deleted_at"`
	DeletedBy *string    `json:"deletedBy" db:"deleted_by"`
}

type SoftDeleteTrail struct {
	DeletedAt *time.Time `json:"deletedAt" db:"deleted_at"`
}
