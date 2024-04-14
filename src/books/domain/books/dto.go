package books

import "github.com/google/uuid"

type GetBooksDto struct {
	Title *string
}

type CreateBookDto struct {
	Title string
}

type UpdateBookDto struct {
	ID    uuid.UUID
	Title string
}
