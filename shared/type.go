package shared

type ResponseObject struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
}
