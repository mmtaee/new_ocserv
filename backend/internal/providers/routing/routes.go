package routing

import (
	"github.com/labstack/echo/v4"
	panelRoutes "ocserv/internal/api/panel"
	userRoutes "ocserv/internal/api/user"
)

func Register(e *echo.Echo) {
	group := e.Group("/api")
	panelRoutes.Routes(group)
	userRoutes.Routes(group)
}
