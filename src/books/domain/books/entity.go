package books

import "github.com/google/uuid"

type Book struct {
	ID    uuid.UUID
	Title string
}
