package helper

import (
	"net/http"
	"strconv"
)

func CalculateLimitPaginationOffset(limit int, page int) int {
	return limit*page - limit
}

func GetLimitPaginationFromQuery(r *http.Request) (int, int) {
	var page int = 1
	var limit int = 10

	query := r.URL.Query()

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

	return page, limit
}
