package interfaces

import "github.com/google/uuid"

const UserRegisteredEvent = "user.registered"

type UserRegistered struct {
	UserID uuid.UUID `json:"user_id"`
	Name   string    `json:"name"`
	Email  string    `json:"email"`
}
