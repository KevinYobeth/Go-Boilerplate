package response

import (
	"go-boilerplate/shared/errors"
	"go-boilerplate/shared/types"

	"github.com/labstack/echo/v4"
)

func SendHTTP(c echo.Context, err error) error {
	errObj := errors.GetGenericError(err)

	return c.JSON(errors.ErrorMap[errObj.Type], types.ResponseBody{
		Message: errObj.Message,
	})
}
