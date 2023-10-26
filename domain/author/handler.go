package author

import (
	"library/shared"
	helper "library/shared/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func NewAuthorHandler(useCase UseCase) *Handler {
	return &Handler{
		UseCase: useCase,
	}
}

func (i *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	page, limit := helper.GetLimitPaginationFromQuery(r)

	authors, err := i.UseCase.Repo.GetAll(r.Context(), shared.LimitPagination{Page: page, Limit: limit})
	if err != nil {
		helper.ErrorJSON(w, err)
		return
	}

	err = helper.WriteJSON(w, http.StatusOK, shared.ResponseObject{
		Data: authors,
		Metadata: shared.ResponseMetadataObject{
			Pagination: &shared.LimitPagination{
				Page:       page,
				Limit:      limit,
				TotalItems: authors.Count,
			},
		},
		Message: "success get authors",
	})
	if err != nil {
		helper.ErrorJSON(w, err, http.StatusInternalServerError)
	}
}

func (i *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	uuid, err := uuid.Parse(id)
	if err != nil {
		helper.ErrorJSON(w, err)
		return
	}

	author, err := i.UseCase.GetById(r.Context(), uuid)
	if err != nil {
		helper.ErrorJSON(w, err)
		return
	}

	err = helper.WriteJSON(w, http.StatusOK, shared.ResponseObject{
		Data:    author,
		Message: "success get author",
	})
	if err != nil {
		helper.ErrorJSON(w, err, http.StatusInternalServerError)
	}
}

func (i *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var body UpsertAuthorEntity

	err := helper.ReadJSON(w, r, &body)
	if err != nil {
		helper.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	author, err := i.UseCase.Create(r.Context(), body)
	if err != nil {
		helper.ErrorJSON(w, err)
		return
	}

	err = helper.WriteJSON(w, http.StatusCreated, shared.ResponseObject{
		Data:     author,
		Message:  "success create author",
		Metadata: shared.ResponseMetadataObject{},
	})
	if err != nil {
		helper.ErrorJSON(w, err, http.StatusInternalServerError)
	}
}

func (i *Handler) UpdateById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var body UpsertAuthorEntity

	uuid, err := uuid.Parse(id)
	if err != nil {
		helper.ErrorJSON(w, err)
		return
	}

	err = helper.ReadJSON(w, r, &body)
	if err != nil {
		helper.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	author, err := i.UseCase.Update(r.Context(), uuid, body)
	if err != nil {
		helper.ErrorJSON(w, err)
		return
	}

	err = helper.WriteJSON(w, http.StatusOK, shared.ResponseObject{
		Data:    author,
		Message: "success update author",
	})
	if err != nil {
		helper.ErrorJSON(w, err, http.StatusInternalServerError)
	}
}

func (i *Handler) DeleteById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	uuid, err := uuid.Parse(id)
	if err != nil {
		helper.ErrorJSON(w, err)
		return
	}

	err = i.UseCase.DeleteById(r.Context(), uuid)
	if err != nil {
		helper.ErrorJSON(w, err)
		return
	}

	err = helper.WriteJSON(w, http.StatusOK, shared.ResponseObject{
		Message: "success delete author",
	})
	if err != nil {
		helper.ErrorJSON(w, err, http.StatusInternalServerError)
	}
}
