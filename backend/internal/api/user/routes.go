package user

import (
	"github.com/labstack/echo/v4"
	"ocserv/pkg/routing/middlewares"
)

func Routes(e *echo.Group) {
	controller := New()

	groupBase := e.Group("/user", middlewares.AuthMiddleware())
	groupBase.GET("/profile", controller.Profile)
	groupBase.POST("/password", controller.ChangePassword)

	groupWithPerm := e.Group("/user", middlewares.AuthMiddleware(), middlewares.AdminPermission())
	groupWithPerm.GET("/staffs", controller.Staffs)
	groupWithPerm.PUT("/staffs/permissions/:id", controller.UpdateStaffPermission)
	groupWithPerm.DELETE("/staffs/:id", controller.RemoveStaff)
	groupWithPerm.POST("/staffs/:id/password", controller.ChangeStaffPassword)
	groupWithPerm.POST("/staffs", controller.CreateStaff)
}
