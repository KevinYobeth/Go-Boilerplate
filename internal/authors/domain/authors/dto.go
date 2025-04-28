package authors

import "github.com/google/uuid"

type GetAuthorsDto struct {
	Name *string
}

type CreateAuthorDto struct {
	ID   uuid.UUID
	Name string
}

func NewCreateAuthorDto(name string, ID *uuid.UUID) CreateAuthorDto {
	id := uuid.New()
	if ID != nil {
		id = *ID
	}

	return CreateAuthorDto{
		ID:   id,
		Name: name,
	}
}
