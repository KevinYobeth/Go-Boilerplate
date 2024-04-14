package books

type GetBooksDto struct {
	Title *string
}

type CreateBookDto struct {
	Title string `json:"title"`
}
