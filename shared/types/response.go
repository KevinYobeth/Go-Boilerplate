package types

type ResponseBody struct {
	Data     interface{} `json:"data,omitempty"`
	Message  string      `json:"message"`
	Metadata interface{} `json:"metadata,omitempty"`
}
