package http

import (
	"go-boilerplate/internal/authors/domain/authors"
	"go-boilerplate/internal/authors/services"
	"go-boilerplate/internal/authors/services/command"
	"go-boilerplate/internal/authors/services/query"
	"go-boilerplate/shared/response"
	"go-boilerplate/shared/types"
	"go-boilerplate/shared/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HTTPTransport struct {
	app *services.Application
}

func NewAuthorsHTTPServer(app *services.Application) HTTPTransport {
	return HTTPTransport{app: app}
}

func (h HTTPTransport) RegisterHTTPRoutes(r *echo.Group) {
	api := r.Group("/v1")
	RegisterHandlers(api, h)
}

// GET /authors
func (h HTTPTransport) GetAuthors(c echo.Context) error {
	authorsObj, err := h.app.Queries.GetAuthors.Handle(c.Request().Context(), query.GetAuthorsRequest{})
	if err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return err
	}

	response.SendHTTP(c, &types.Response{
		Body: GetAuthorsResponse{
			Data:    TransformToHTTPAuthors(authorsObj),
			Message: "success get authors",
		},
		StatusCode: http.StatusOK,
	})
	return nil
}

// GET /authors/:id
func (h HTTPTransport) GetAuthor(c echo.Context, id string) error {
	parsedUUID, err := utils.ParseUUID(id)
	if err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return err
	}

	author, err := h.app.Queries.GetAuthor.Handle(c.Request().Context(), query.GetAuthorRequest{ID: parsedUUID})
	if err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return err
	}

	response.SendHTTP(c, &types.Response{
		Body: GetAuthorResponse{
			Data:    TransformToHTTPAuthor(author),
			Message: "success get author",
		},
	})
	return nil
}

// DELETE /authors/:id
func (h HTTPTransport) DeleteAuthor(c echo.Context, id string) error {
	parsedUUID, err := utils.ParseUUID(id)
	if err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return err
	}

	err = h.app.Commands.DeleteAuthor.Handle(c.Request().Context(), command.DeleteAuthorRequest{ID: parsedUUID})
	if err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return err
	}

	response.SendHTTP(c, &types.Response{
		Body: MessageResponse{
			Message: "success delete author",
		},
	})
	return nil
}

// POST /authors
func (h HTTPTransport) CreateAuthor(c echo.Context) error {
	var request authors.CreateAuthorDto
	if err := c.Bind(&request); err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return err
	}

	err := h.app.Commands.CreateAuthor.Handle(c.Request().Context(), command.CreateAuthorRequest{Name: request.Name})
	if err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return err
	}

	response.SendHTTP(c, &types.Response{
		StatusCode: http.StatusCreated,
		Body: MessageResponse{
			Message: "success create author",
		},
	})
	return nil
}
