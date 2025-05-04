package http

import (
	"github.com/google/uuid"
	"github.com/kevinyobeth/go-boilerplate/internal/link/services"
	"github.com/kevinyobeth/go-boilerplate/internal/link/services/command"
	"github.com/kevinyobeth/go-boilerplate/internal/link/services/query"
	"github.com/kevinyobeth/go-boilerplate/shared/middlewares/http"
	"github.com/kevinyobeth/go-boilerplate/shared/response"
	"github.com/kevinyobeth/go-boilerplate/shared/types"
	"github.com/labstack/echo/v4"
)

type HTTPTransport struct {
	app *services.Application
}

func NewLinkHTTPServer(app *services.Application) HTTPTransport {
	return HTTPTransport{app: app}
}

func (h HTTPTransport) RegisterHTTPRoutes(r *echo.Group) {
	api := r.Group("/v1")
	RegisterHandlers(api, h)
}

// POST /links
func (h HTTPTransport) CreateLink(c echo.Context) error {
	claims, ctx, err := http.AuthenticatedMiddleware(c)
	if err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return err
	}

	var request CreateLinkRequest
	if err := c.Bind(&request); err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return err
	}

	err = h.app.Commands.ShortenLink.Handle(ctx, &command.ShortenLinkRequest{
		Slug:        request.Slug,
		URL:         request.Url,
		Description: request.Description,
		UserID:      uuid.MustParse(claims.Subject),
	})
	if err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return err
	}

	response.SendHTTP(c, &types.Response{
		Body: MessageResponse{
			Message: "success create link",
		},
	})
	return nil
}

// GET /links
func (h HTTPTransport) GetLinks(c echo.Context) error {
	claims, ctx, err := http.AuthenticatedMiddleware(c)
	if err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return err
	}

	links, err := h.app.Queries.GetLinks.Handle(ctx, &query.GetLinksRequest{
		UserID: uuid.MustParse(claims.Subject),
	})
	if err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return err
	}

	response.SendHTTP(c, &types.Response{
		Body: GetLinksResponse{
			Data:    TransformToHTTPLinks(links),
			Message: "success get links",
		},
	})
	return nil
}

// GET /links/:slug
func (h HTTPTransport) GetLink(c echo.Context, slug string) error {
	link, err := h.app.Queries.GetRedirectLink.Handle(c.Request().Context(), &query.GetRedirectLinkRequest{
		Slug: slug,
	})
	if err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return err
	}

	c.Redirect(302, link.URL)
	return nil
}
