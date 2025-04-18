package routing

import (
	"github.com/labstack/echo/v4"
	setupRoutes "ocserv/internal/api/setup"
)

func Register(e *echo.Echo) {
	group := e.Group("/api")
	setupRoutes.Routes(group)
}
