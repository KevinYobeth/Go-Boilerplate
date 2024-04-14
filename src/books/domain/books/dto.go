package books

import "github.com/google/uuid"

type GetBooksDto struct {
	Title *string
}

type CreateBookDto struct {
	ID    uuid.UUID
	Title string
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
