package http

import (
	"go-boilerplate/shared/response"
	"go-boilerplate/shared/types"
	"go-boilerplate/shared/utils"
	"go-boilerplate/src/authors/domain/authors"
	"go-boilerplate/src/authors/services"
	"go-boilerplate/src/authors/services/command"
	"go-boilerplate/src/authors/services/query"
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
	authorsObj, err := h.app.Queries.GetAuthors.Handle(c.Request().Context(), query.GetAuthorsParams{})
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

	author, err := h.app.Queries.GetAuthor.Handle(c.Request().Context(), query.GetAuthorParams{ID: parsedUUID})
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

	err = h.app.Commands.DeleteAuthor.Handle(c.Request().Context(), command.DeleteAuthorParams{ID: parsedUUID})
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

	err := h.app.Commands.CreateAuthor.Handle(c.Request().Context(), command.CreateAuthorParams{Name: request.Name})
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
