package book

import (
	"fmt"
	"library/shared"
	helper "library/shared/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func NewBookHandler(useCase UseCase) *Handler {
	return &Handler{
		UseCase: useCase,
	}
}

func (i *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	page, limit := helper.GetLimitPaginationFromQuery(r)

	books, err := i.UseCase.GetAll(r.Context(), shared.LimitPagination{Page: page, Limit: limit})
	if err != nil {
		helper.ErrorJSON(w, err)
		return
	}

	err = helper.WriteJSON(w, http.StatusOK, shared.ResponseObject{
		Data: books,
		Metadata: shared.ResponseMetadataObject{
			Pagination: &shared.LimitPagination{
				Page:  page,
				Limit: limit,
			},
		},
		Message: "success get books",
	})
	if err != nil {
		helper.ErrorJSON(w, err, http.StatusInternalServerError)
	}
}

func (i *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	book, err := i.UseCase.GetById(r.Context(), uuid.MustParse(id))
	if err != nil {
		helper.ErrorJSON(w, err)
		return
	}

	err = helper.WriteJSON(w, http.StatusOK, shared.ResponseObject{
		Data:    book,
		Message: "success get book",
	})
	if err != nil {
		helper.ErrorJSON(w, err, http.StatusInternalServerError)
	}
}

func (i *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var body UpsertBookEntity

	err := helper.ReadJSON(w, r, &body)
	if err != nil {
		helper.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	book, err := i.UseCase.Create(r.Context(), body)
	if err != nil {
		helper.ErrorJSON(w, err)
		return
	}

	err = helper.WriteJSON(w, http.StatusCreated, shared.ResponseObject{
		Data:     book,
		Message:  "success create book",
		Metadata: shared.ResponseMetadataObject{},
	})
	if err != nil {
		helper.ErrorJSON(w, err, http.StatusInternalServerError)
	}
}

func (i *Handler) UpdateById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var body UpsertBookEntity

	err := helper.ReadJSON(w, r, &body)
	if err != nil {
		fmt.Println("failed to read JSON", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	book, err := i.UseCase.UpdateById(r.Context(), uuid.MustParse(id), body)
	if err != nil {
		helper.ErrorJSON(w, err)
		return
	}

	err = helper.WriteJSON(w, http.StatusOK, shared.ResponseObject{
		Data:    book,
		Message: "success update book",
	})
	if err != nil {
		helper.ErrorJSON(w, err, http.StatusInternalServerError)
	}
}

func (i *Handler) DeleteById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := i.UseCase.DeleteById(r.Context(), uuid.MustParse(id))
	if err != nil {
		helper.ErrorJSON(w, err)
		return
	}

	err = helper.WriteJSON(w, http.StatusOK, shared.ResponseObject{
		Message: "success delete book",
	})
	if err != nil {
		helper.ErrorJSON(w, err, http.StatusInternalServerError)
	}
}
