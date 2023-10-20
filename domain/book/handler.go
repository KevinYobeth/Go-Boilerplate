package book

import "net/http"

type BookInterface struct {
	Repo Repo
}

func (h *BookInterface) Create(w http.ResponseWriter, r *http.Request) {
}
