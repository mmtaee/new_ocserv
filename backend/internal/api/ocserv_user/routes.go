package ocservUser

import (
	"github.com/labstack/echo/v4"
	"ocserv/pkg/routing/middlewares"
)

func Routes(e *echo.Group) {
	controller := New()

	group := e.Group("/oc_users", middlewares.AuthMiddleware(), middlewares.RoutePermission())

	group.GET("", controller.GetUsers)
	group.POST("", controller.CreateUser)

}
