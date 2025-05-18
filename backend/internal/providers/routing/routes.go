package routing

import (
	"github.com/labstack/echo/v4"
	ocservUserRoutes "ocserv/internal/api/ocserv_user"
	panelRoutes "ocserv/internal/api/panel"
	userRoutes "ocserv/internal/api/user"
)

func Register(e *echo.Echo) {
	group := e.Group("/api")
	panelRoutes.Routes(group)
	userRoutes.Routes(group)
	ocservUserRoutes.Routes(group)
}
