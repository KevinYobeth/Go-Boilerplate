package authors

import "github.com/google/uuid"

type Author struct {
	ID   uuid.UUID `json:"author_id"`
	Name string    `json:"author_name"`
}
