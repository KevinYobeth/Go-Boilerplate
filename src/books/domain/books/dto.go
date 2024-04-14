package books

import "github.com/google/uuid"

type GetBooksDto struct {
	Title *string
}

type CreateBookDto struct {
	ID     uuid.UUID
	Title  string
	Author uuid.UUID
}

func NewCreateBookDto(title string) CreateBookDto {
	return CreateBookDto{
		ID:    uuid.New(),
		Title: title,
	}
}

type UpdateBookDto struct {
	ID    uuid.UUID
	Title string
}

type CreateAuthorBookDto struct {
	BookID   uuid.UUID
	AuthorID uuid.UUID
}
