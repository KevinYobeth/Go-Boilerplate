package shared

type ResponseObject struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
}

type LimitPagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
