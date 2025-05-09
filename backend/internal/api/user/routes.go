package user

import (
	"github.com/labstack/echo/v4"
	"ocserv/pkg/routing/middlewares"
)

func Routes(e *echo.Group) {
	controller := New()
	e = e.Group("/user", middlewares.AuthMiddleware())

	e.GET("/profile", controller.Profile)
	e.POST("/password", controller.ChangePassword)

	e.GET("/staffs", controller.Staffs, middlewares.AdminPermission())
}
