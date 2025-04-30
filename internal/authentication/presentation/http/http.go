package http

import (
	"go-boilerplate/internal/authentication/services"
	"go-boilerplate/internal/authentication/services/command"
	"go-boilerplate/internal/authentication/services/query"
	"go-boilerplate/shared/middlewares/http"
	"go-boilerplate/shared/response"
	"go-boilerplate/shared/types"

	"github.com/google/uuid"
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

	token, err := h.app.Queries.Login.Handle(c.Request().Context(), query.LoginRequest{
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

	err := h.app.Commands.Register.Handle(c.Request().Context(), command.RegisterRequest{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Password:  request.Password,
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

	token, err := h.app.Queries.RefreshToken.Handle(c.Request().Context(), query.RefreshTokenRequest{
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

// GET /user
func (h HTTPTransport) GetUser(c echo.Context) error {
	claims, ctx, err := http.AuthenticatedMiddleware(c)
	if err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return err
	}

	user, err := h.app.Queries.GetUser.Handle(ctx, query.GetUserRequest{
		ID: uuid.MustParse(claims.Subject),
	})
	if err != nil {
		response.SendHTTP(c, &types.Response{
			Error: err,
		})
		return err
	}

	response.SendHTTP(c, &types.Response{
		Body: UserResponse{
			Data:    TransformToHTTPUser(user),
			Message: "success get user",
		},
	})
	return nil
}
