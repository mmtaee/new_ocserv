package setup

import "github.com/labstack/echo/v4"

func Routes(e *echo.Group) {
	controller := New()
	e = e.Group("/panel")
	e.POST("/setup", controller.Setup)
}
