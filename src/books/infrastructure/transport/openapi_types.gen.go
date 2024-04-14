// Package transport provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.15.0 DO NOT EDIT.
package transport

import (
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Book defines model for Book.
type Book struct {
	Id   openapi_types.UUID `json:"id"`
	Name string             `json:"name"`
}

// CreateBookRequest defines model for CreateBookRequest.
type CreateBookRequest struct {
	Name string `json:"name"`
}

// GetBookResponse defines model for GetBookResponse.
type GetBookResponse struct {
	Data    Book   `json:"data"`
	Message string `json:"message"`
}

// GetBooksResponse defines model for GetBooksResponse.
type GetBooksResponse struct {
	Data    []Book `json:"data"`
	Message string `json:"message"`
}

// MessageResponse defines model for MessageResponse.
type MessageResponse struct {
	Message string `json:"message"`
}

// CreateBookJSONRequestBody defines body for CreateBook for application/json ContentType.
type CreateBookJSONRequestBody = CreateBookRequest
