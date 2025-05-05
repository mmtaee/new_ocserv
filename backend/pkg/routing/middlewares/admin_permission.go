package middlewares

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func AdminPermission() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Get("isAdmin").(bool) {
				return next(c)
			}
			return echo.NewHTTPError(http.StatusForbidden, "Admin permission required")
		}
	}
}
