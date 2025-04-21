package panel

import "github.com/labstack/echo/v4"

func Routes(e *echo.Group) {
	controller := New()
	e = e.Group("/panel")

	e.GET("", controller.Config)
	e.POST("/setup", controller.Setup)
}
