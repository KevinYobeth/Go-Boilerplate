package shared

type ResponseObject struct {
	Data     any                    `json:"data"`
	Message  string                 `json:"message"`
	Metadata ResponseMetadataObject `json:"metadata"`
}

type ResponseMetadataObject struct {
	IsError    bool             `json:"isError"`
	Pagination *LimitPagination `json:"pagination,omitempty"`
}

type LimitPagination struct {
	Page  int `json:"page,omitempty"`
	Limit int `json:"limit,omitempty"`
}
