package author

import (
	"fmt"
	"library/shared"
	helper "library/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type AuthorInterface struct {
	Repo Repo
}

func (i *AuthorInterface) Create(w http.ResponseWriter, r *http.Request) {
	var body CreateAuthorEntity

	err := helper.ReadJSON(w, r, &body)
	if err != nil {
		fmt.Println("failed to read JSON", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	author := Author{
		Id:        uuid.New(),
		Name:      body.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = i.Repo.Insert(r.Context(), author)
	if err != nil {
		fmt.Println("failed to insert: ", err)
		helper.WriteJSON(w, http.StatusInternalServerError, shared.ResponseObject{
			Message: "something went wrong",
			Data:    nil,
		})
	}

	err = helper.WriteJSON(w, http.StatusCreated, author)
	if err != nil {
		fmt.Println("something went wrong", err)
		helper.WriteJSON(w, http.StatusInternalServerError, shared.ResponseObject{
			Message: "something went wrong",
			Data:    nil,
		})
		return
	}
}

func (i *AuthorInterface) List(w http.ResponseWriter, r *http.Request) {
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

	authors, err := i.Repo.GetAll(r.Context(), helper.LimitPagination{Page: page, Limit: limit})

	if err != nil {
		fmt.Println("something went wrong", err)
		helper.WriteJSON(w, http.StatusInternalServerError, shared.ResponseObject{
			Message: "something went wrong",
			Data:    nil,
		})
		return
	}

	helper.WriteJSON(w, http.StatusOK, authors)
}
