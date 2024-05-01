package authors

import "github.com/google/uuid"

type DeleteAuthorEvent struct {
	ID uuid.UUID
}
