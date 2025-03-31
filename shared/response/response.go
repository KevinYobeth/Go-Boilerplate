package response

import (
	"encoding/json"
	"go-boilerplate/shared/errors"
	"go-boilerplate/shared/telemetry"
	"go-boilerplate/shared/types"

	"github.com/labstack/echo/v4"
)

func SendHTTP(c echo.Context, response *types.Response) error {
	body := types.ResponseBody{}
	jsonBytes, err := json.Marshal(response.Body)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Internal Server Error"})
	}
	json.Unmarshal(jsonBytes, &body)

	body.TraceID = telemetry.GetTraceID(c.Request().Context())

	if response.Error != nil {
		errObj := errors.GetGenericError(response.Error)
		body.Message = errObj.Message

		return c.JSON(errors.ErrorMap[errObj.Type], body)
	}

	return c.JSON(200, body)
}
