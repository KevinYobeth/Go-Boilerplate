package book

import (
	"fmt"
	"library/domain/author"
	"library/shared"
	model "library/shared/models"
	helper "library/shared/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type BookInterface struct {
	Repo       Repo
	AuthorRepo author.Repo
}

func (i *BookInterface) Create(w http.ResponseWriter, r *http.Request) {
	var body UpsertBookEntity

	err := helper.ReadJSON(w, r, &body)
	if err != nil {
		fmt.Println("failed to read JSON", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	author, err := i.AuthorRepo.GetById(r.Context(), body.AuthorId.String())
	if err != nil {
		helper.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	book := model.Book{
		Id:        uuid.New(),
		Title:     body.Title,
		AuthorId:  author.Id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = i.Repo.Insert(r.Context(), book)
	if err != nil {
		fmt.Println("failed to insert: ", err)
		helper.WriteJSON(w, http.StatusInternalServerError, shared.ResponseObject{
			Message: "something went wrong",
			Data:    nil,
		})
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

func (i *BookInterface) List(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	var page int = 1
	var limit int = 10

	if query.Get("page") != "" {
		convertedPage, err := strconv.Atoi(query.Get("page"))
		if err != nil {
			panic(err)
		}
		page = convertedPage
	}

	if query.Get("limit") != "" {
		convertedLimit, err := strconv.Atoi(query.Get("limit"))
		if err != nil {
			panic(err)
		}
		limit = convertedLimit
	}

	books, err := i.Repo.GetAll(r.Context(), shared.LimitPagination{Page: page, Limit: limit})

	if err != nil {
		fmt.Println("something went wrong", err)
		helper.WriteJSON(w, http.StatusInternalServerError, shared.ResponseObject{
			Message: "something went wrong",
			Data:    nil,
		})
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

func (i *BookInterface) GetById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	book, err := i.Repo.GetById(r.Context(), id)
	if err != nil {
		fmt.Println("something went wrong", err)
		helper.WriteJSON(w, http.StatusInternalServerError, shared.ResponseObject{
			Message: "something went wrong",
			Data:    nil,
		})
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

func (i *BookInterface) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var body UpsertBookEntity

	bookFromDB, err := i.Repo.GetById(r.Context(), id)
	if err != nil {
		fmt.Println("something went wrong", err)
		helper.WriteJSON(w, http.StatusInternalServerError, shared.ResponseObject{
			Message: "something went wrong",
			Data:    nil,
		})
		return
	}

	err = helper.ReadJSON(w, r, &body)
	if err != nil {
		fmt.Println("failed to read JSON", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var authorId uuid.UUID

	if body.AuthorId != uuid.Nil {
		author, err := i.AuthorRepo.GetById(r.Context(), body.AuthorId.String())
		if err != nil {
			helper.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		authorId = author.Id
	} else {
		authorId = bookFromDB.AuthorId
	}

	book := model.Book{
		Id:        bookFromDB.Id,
		Title:     body.Title,
		AuthorId:  authorId,
		CreatedAt: bookFromDB.CreatedAt,
		UpdatedAt: time.Now(),
	}

	err = i.Repo.Update(r.Context(), id, book)
	if err != nil {
		fmt.Println("something went wrong", err)
		helper.WriteJSON(w, http.StatusInternalServerError, shared.ResponseObject{
			Message: "something went wrong",
			Data:    nil,
		})
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

func (i *BookInterface) DeleteById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := i.Repo.DeleteById(r.Context(), id)
	if err != nil {
		fmt.Println("something went wrong", err)
		helper.WriteJSON(w, http.StatusInternalServerError, shared.ResponseObject{
			Message: "something went wrong",
			Data:    nil,
		})
		return
	}

	err = helper.WriteJSON(w, http.StatusOK, shared.ResponseObject{
		Message: "success delete book",
	})
	if err != nil {
		helper.ErrorJSON(w, err, http.StatusInternalServerError)
	}
}
