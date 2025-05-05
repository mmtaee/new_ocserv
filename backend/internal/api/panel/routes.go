package panel

import (
	"github.com/labstack/echo/v4"
	"ocserv/pkg/routing/middlewares"
)

func Routes(e *echo.Group) {
	controller := New()
	e = e.Group("/panel")

	e.GET("/init", controller.Init)
	e.POST("/setup", controller.Setup)
	e.POST("/login", controller.Login)
	e.GET("/config", controller.Config, middlewares.AuthMiddleware(), middlewares.AdminPermission())
}
