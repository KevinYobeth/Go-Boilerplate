package http

import (
	"context"

	"github.com/kevinyobeth/go-boilerplate/internal/authentication/services"
	"github.com/kevinyobeth/go-boilerplate/internal/authentication/services/command"
	"github.com/kevinyobeth/go-boilerplate/internal/authentication/services/query"
	"github.com/kevinyobeth/go-boilerplate/shared/middlewares/http"
	"github.com/kevinyobeth/go-boilerplate/shared/response"
	"github.com/kevinyobeth/go-boilerplate/shared/types"

	"github.com/labstack/echo/v4"
)

type HTTPTransport struct {
	app *services.Application
}

func NewAuthenticationHTTPServer(app *services.Application) HTTPTransport {
	return HTTPTransport{app: app}
}

func (h HTTPTransport) RegisterHTTPRoutes(r *echo.Group) {
	api := r.Group("/v1")
	RegisterHandlers(api, h)
}

// POST /login
func (h HTTPTransport) Login(c echo.Context) error {
	var request LoginRequest
	if err := c.Bind(&request); err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return err
	}

	token, err := h.app.Queries.Login.Handle(c.Request().Context(), &query.LoginRequest{
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return err
	}

	response.SendHTTP(c, &types.Response{
		Body: LoginResponse{
			Data:    TransformToHTTPToken(token),
			Message: "success login",
		},
	})
	return nil
}

// POST /register
func (h HTTPTransport) Register(c echo.Context) error {
	var request RegisterRequest
	if err := c.Bind(&request); err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return err
	}

	err := http.TransactionMiddleware(c.Request().Context(), func(ctx context.Context) error {
		return h.app.Commands.Register.Handle(ctx, &command.RegisterRequest{
			FirstName: request.FirstName,
			LastName:  request.LastName,
			Email:     request.Email,
			Password:  request.Password,
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
			Message: "success register",
		},
	})
	return nil
}

func (h HTTPTransport) RefreshToken(c echo.Context) error {
	var request RefreshTokenRequest
	if err := c.Bind(&request); err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return err
	}

	token, err := h.app.Queries.RefreshToken.Handle(c.Request().Context(), &query.RefreshTokenRequest{
		RefreshToken: request.RefreshToken,
	})
	if err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return err
	}

	response.SendHTTP(c, &types.Response{
		Body: LoginResponse{
			Data:    TransformToHTTPToken(token),
			Message: "success refresh token",
		},
	})
	return nil
}
