package http

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinyobeth/go-boilerplate/internal/link/services"
	"github.com/kevinyobeth/go-boilerplate/internal/link/services/command"
	"github.com/kevinyobeth/go-boilerplate/internal/link/services/query"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/middlewares/http"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/response"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/types"
	"github.com/labstack/echo/v4"
)

type HTTPTransport struct {
	app *services.Application
}

func NewLinkHTTPServer(app *services.Application) HTTPTransport {
	return HTTPTransport{app: app}
}

func (h HTTPTransport) RegisterHTTPRoutes(r *echo.Group, root *echo.Echo) {
	api := r.Group("/v1")

	RegisterHandlers(api, h)

	root.GET("/:slug", func(c echo.Context) error {
		slug := c.Param("slug")
		return h.GetRedirectLink(c, slug)
	})
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

	err = http.TransactionMiddleware(ctx, func(ctx context.Context) error {
		return h.app.Commands.ShortenLink.Handle(ctx, &command.ShortenLinkRequest{
			Slug:        request.Slug,
			URL:         request.Url,
			Description: request.Description,
			UserID:      uuid.MustParse(claims.Subject),
		})
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
func (h HTTPTransport) GetLinks(c echo.Context, params GetLinksParams) error {
	claims, ctx, err := http.AuthenticatedMiddleware(c)
	if err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return err
	}

	next := params.Next
	prev := params.Prev
	if next == nil {
		next = &uuid.Nil
	}
	if prev == nil {
		prev = &uuid.Nil
	}
	limit := uint64(10)
	if params.Limit != nil {
		limit = *params.Limit
	}

	links, err := h.app.Queries.GetLinks.Handle(ctx, &query.GetLinksRequest{
		UserID: uuid.MustParse(claims.Subject),
		Next:   *next,
		Prev:   *prev,
		Limit:  limit,
	})
	if err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return err
	}

	response.SendHTTP(c, &types.Response{
		Body: GetLinksResponse{
			Data:     TransformToHTTPLinks(links.Data),
			Metadata: TransformToHTTPMetadata(links.Metadata),
			Message:  "success get links",
		},
	})
	return nil
}

// GET /links/:id
func (h HTTPTransport) GetLink(c echo.Context, id uuid.UUID) error {
	claims, ctx, err := http.AuthenticatedMiddleware(c)
	if err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return err
	}

	link, err := h.app.Queries.GetLink.Handle(ctx, &query.GetLinkRequest{
		ID:     id,
		UserID: uuid.MustParse(claims.Subject),
	})
	if err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return err
	}

	response.SendHTTP(c, &types.Response{
		Body: GetLinkResponse{
			Data:    TransformToHTTPLink(link),
			Message: "success get link",
		},
	})
	return nil
}

// PUT /links/:id
func (h HTTPTransport) UpdateLink(c echo.Context, id uuid.UUID) error {
	claims, ctx, err := http.AuthenticatedMiddleware(c)
	if err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return err
	}

	var request UpdateLinkRequest
	if err := c.Bind(&request); err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return err
	}

	err = h.app.Commands.UpdateLink.Handle(ctx, &command.UpdateLinkRequest{
		ID:          id,
		UserID:      uuid.MustParse(claims.Subject),
		Slug:        request.Slug,
		URL:         request.Url,
		Description: request.Description,
	})
	if err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return err
	}

	response.SendHTTP(c, &types.Response{
		Body: MessageResponse{
			Message: "success update link",
		},
	})
	return nil
}

// DELETE /links/:id
func (h HTTPTransport) DeleteLink(c echo.Context, id uuid.UUID) error {
	claims, ctx, err := http.AuthenticatedMiddleware(c)
	if err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return err
	}

	err = h.app.Commands.DeleteLink.Handle(ctx, &command.DeleteLinkRequest{
		UserID: uuid.MustParse(claims.Subject),
		ID:     id,
	})
	if err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return err
	}

	response.SendHTTP(c, &types.Response{
		Body: MessageResponse{
			Message: "success delete link",
		},
	})
	return nil
}

// GET /:slug
func (h HTTPTransport) GetRedirectLink(c echo.Context, slug string) error {
	link, err := h.app.Queries.GetRedirectLink.Handle(c.Request().Context(), &query.GetRedirectLinkRequest{
		Slug: slug,
		Metadata: query.LinkVisitEventMetadata{
			IPAddress:   c.RealIP(),
			UserAgent:   c.Request().UserAgent(),
			Referer:     c.Request().Header.Get("Referer"),
			CountryCode: c.Request().Header.Get("X-Country-Code"),
			DeviceType:  c.Request().Header.Get("X-Device-Type"),
			Browser:     c.Request().Header.Get("X-Browser"),
		},
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
