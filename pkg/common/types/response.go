package types

type Response struct {
	Body       interface{}
	StatusCode int
	Error      error
}

type ResponseBody struct {
	Data     interface{} `json:"data,omitempty"`
	Message  string      `json:"message"`
	Metadata interface{} `json:"metadata,omitempty"`
	TraceID  string      `json:"trace_id,omitempty"`
	Error    string      `json:"error,omitempty"`
}
