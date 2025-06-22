package http

import (
	"github.com/google/uuid"
	"github.com/kevinyobeth/go-boilerplate/internal/user/services"
	"github.com/kevinyobeth/go-boilerplate/internal/user/services/query"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/middlewares/http"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/response"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/types"
	"github.com/labstack/echo/v4"
)

type HTTPTransport struct {
	app *services.Application
}

func NewUserHTTPServer(app *services.Application) HTTPTransport {
	return HTTPTransport{app: app}
}

func (h HTTPTransport) RegisterHTTPRoutes(r *echo.Group) {
	api := r.Group("/v1")
	RegisterHandlers(api, h)
}

// GET /profile
func (h HTTPTransport) GetProfile(c echo.Context) error {
	claims, ctx, err := http.AuthenticatedMiddleware(c)
	if err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return err
	}

	user, err := h.app.Queries.GetUser.Handle(ctx, &query.GetUserRequest{
		ID: uuid.MustParse(claims.Subject),
	})
	if err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return err
	}

	response.SendHTTP(c, &types.Response{
		Body: ProfileResponse{
			Data:    TransformToHTTPUser(user),
			Message: "success get profile",
		},
	})
	return nil
}
