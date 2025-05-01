package response

import (
	"encoding/json"
	"net/http"

	"github.com/kevinyobeth/go-boilerplate/shared/errors"
	"github.com/kevinyobeth/go-boilerplate/shared/telemetry"
	"github.com/kevinyobeth/go-boilerplate/shared/types"

	"github.com/labstack/echo/v4"
)

func SendHTTP(c echo.Context, response *types.Response) error {
	body := types.ResponseBody{}
	jsonBytes, err := json.Marshal(response.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.ResponseBody{
			Message: "Internal Server Error",
		})
	}
	json.Unmarshal(jsonBytes, &body)

	body.TraceID = telemetry.GetTraceID(c.Request().Context())

	var statusCode = 200
	if response.StatusCode != 0 {
		statusCode = response.StatusCode
	}

	if response.Error != nil {
		errObj := errors.GetGenericError(response.Error)
		body.Message = errObj.Message

		if errObj.Unwrap() != nil {
			body.Error = errObj.Unwrap().Error()
		}

		if response.StatusCode == 0 {
			statusCode = errors.ErrorMap[errObj.Type]
		}

		return c.JSON(statusCode, body)
	}

	return c.JSON(statusCode, body)
}
