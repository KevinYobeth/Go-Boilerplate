package http

import (
	"context"
	"go-boilerplate/shared/database"
	respond "go-boilerplate/shared/response"
	"go-boilerplate/shared/utils"
	"go-boilerplate/src/authors/domain/authors"
	"go-boilerplate/src/authors/services"
	"go-boilerplate/src/authors/services/command"
	"go-boilerplate/src/authors/services/query"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HTTPTransport struct {
	app     *services.Application
	manager database.TransactionManager
}

func NewAuthorsHTTPServer(app *services.Application) HTTPTransport {
	return HTTPTransport{app: app}
}

func (h HTTPTransport) TransactionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		return h.manager.RunInTransaction(echoCtx.Request().Context(), func(c context.Context) error {
			ctx := echoCtx.Echo().NewContext(echoCtx.Request().WithContext(c), echoCtx.Response())
			ctx.SetParamNames(echoCtx.ParamNames()...)
			ctx.SetParamValues(echoCtx.ParamValues()...)

			return next(ctx)
		})
	}
}

func (h HTTPTransport) RegisterHTTPRoutes(r *echo.Group) {
	api := r.Group("/v1")
	RegisterHandlers(api, h)
}

// GET /authors
func (h HTTPTransport) GetAuthors(c echo.Context) error {
	authorsObj, err := h.app.Queries.GetAuthors.Execute(c.Request().Context(), query.GetAuthorsParams{})
	if err != nil {
		respond.SendHTTP(c, err)
		return err
	}

	c.JSON(http.StatusOK, GetAuthorsResponse{
		Data:    TransformToHTTPAuthors(authorsObj),
		Message: "success get authors",
	})
	return nil
}

// GET /authors/:id
func (h HTTPTransport) GetAuthor(c echo.Context, id string) error {
	parsedUUID, err := utils.ParseUUID(id)
	if err != nil {
		respond.SendHTTP(c, err)
		return err
	}

	author, err := h.app.Queries.GetAuthor.Execute(c.Request().Context(), query.GetAuthorParams{ID: parsedUUID})
	if err != nil {
		respond.SendHTTP(c, err)
		return err
	}

	c.JSON(http.StatusOK, GetAuthorResponse{
		Data:    TransformToHTTPAuthor(author),
		Message: "success get author",
	})
	return nil
}

// DELETE /authors/:id
func (h HTTPTransport) DeleteAuthor(c echo.Context, id string) error {
	parsedUUID, err := utils.ParseUUID(id)
	if err != nil {
		respond.SendHTTP(c, err)
		return err
	}

	err = h.app.Commands.DeleteAuthor.Execute(c.Request().Context(), command.DeleteAuthorParams{ID: parsedUUID})
	if err != nil {
		respond.SendHTTP(c, err)
		return err
	}

	c.JSON(http.StatusOK, MessageResponse{
		Message: "success delete author",
	})
	return nil
}

// POST /authors
func (h HTTPTransport) CreateAuthor(c echo.Context) error {
	var request authors.CreateAuthorDto
	if err := c.Bind(&request); err != nil {
		respond.SendHTTP(c, err)
		return err
	}

	_, err := h.app.Commands.CreateAuthor.Execute(c.Request().Context(), command.CreateAuthorParams{Name: request.Name})
	if err != nil {
		respond.SendHTTP(c, err)
		return err
	}

	c.JSON(http.StatusCreated, MessageResponse{
		Message: "success create author",
	})
	return nil
}
