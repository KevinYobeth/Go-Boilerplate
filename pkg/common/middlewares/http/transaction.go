package http

import (
	"context"

	"github.com/kevinyobeth/go-boilerplate/pkg/common/database"
	"github.com/labstack/echo/v4"
)

func TransactionMiddleware(ctx context.Context, f func(c context.Context) error) error {
	db := database.InitPostgres()
	manager := database.NewTransactionManager(db)

	return manager.RunInTransaction(ctx, f)
}

func GlobalTransactionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	db := database.InitPostgres()
	manager := database.NewTransactionManager(db)

	return func(echoCtx echo.Context) error {
		return manager.RunInTransaction(echoCtx.Request().Context(), func(c context.Context) error {
			ctx := echoCtx.Echo().NewContext(echoCtx.Request().WithContext(c), echoCtx.Response())
			ctx.SetParamNames(echoCtx.ParamNames()...)
			ctx.SetParamValues(echoCtx.ParamValues()...)

			return next(ctx)
		})
	}
}
