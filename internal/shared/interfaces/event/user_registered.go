package interfaces

import "github.com/google/uuid"

const UserRegisteredTopic = "user.registered"

type UserRegistered struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
}
