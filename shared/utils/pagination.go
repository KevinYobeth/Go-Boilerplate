package helper

func CalculateLimitPaginationOffset(limit int, page int) int {
	return limit*page - limit
}
