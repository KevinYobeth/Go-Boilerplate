package authors

import "github.com/google/uuid"

type Author struct {
	ID   uuid.UUID
	Name string
}
