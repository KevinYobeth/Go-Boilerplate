package http

import (
	"context"

	"github.com/kevinyobeth/go-boilerplate/config"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/constants"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/entity"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/errors"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func AuthenticatedMiddleware(c echo.Context) (*entity.Claims, context.Context, error) {
	authHeader := c.Request().Header.Get("Authorization")

	if authHeader == "" {
		return nil, c.Request().Context(), errors.NewUnauthenticatedError(nil)
	}

	token := authHeader[len("Bearer "):]
	if token == "" {
		return nil, c.Request().Context(), errors.NewUnauthenticatedError(nil)
	}

	jwtConfig := config.LoadJWTConfig()
	claims := &entity.Claims{}
	jwtToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.NewIncorrectInputError(nil, "invalid signing method")
		}

		return []byte(jwtConfig.JWTSecret), nil
	})
	if err != nil {
		return nil, c.Request().Context(), errors.NewUnauthenticatedError(err)
	}

	if !jwtToken.Valid {
		return nil, c.Request().Context(), errors.NewUnauthenticatedError(nil)
	}

	ctx := utils.AddToCtx(c.Request().Context(), constants.ContextKeyClaims, claims)

	return claims, ctx, nil
}

func GlobalAuthenticatedMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, ctx, err := AuthenticatedMiddleware(c)
		if err != nil {
			return err
		}

		context := c.Echo().NewContext(c.Request().WithContext(ctx), c.Response())

		return next(context)
	}
}
