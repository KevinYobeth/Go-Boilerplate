package user

import (
	"time"

	valueobjects "github.com/kevinyobeth/go-boilerplate/shared/value_objects"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`

	valueobjects.AuditAuthor
	valueobjects.AuditTrail
}

type VerificationToken struct {
	ID     uuid.UUID  `json:"id"`
	UserID uuid.UUID  `json:"user_id"`
	Token  string     `json:"token"`
	UsedAt *time.Time `json:"used_at"`

	valueobjects.AuditAuthor
	valueobjects.AuditTrail
}
