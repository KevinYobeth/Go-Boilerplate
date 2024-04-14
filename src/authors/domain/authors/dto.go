package authors

import "github.com/google/uuid"

type GetAuthorsDto struct {
	Name *string
}

type CreateAuthorDto struct {
	ID   uuid.UUID
	Name string
}

func NewCreateAuthorDto(name string) CreateAuthorDto {
	return CreateAuthorDto{
		ID:   uuid.New(),
		Name: name,
	}
}
