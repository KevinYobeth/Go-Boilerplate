package transport

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterAuthorHTTPRoutes(r echo.Group) {
	r.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "yang ini author")
	})
}
