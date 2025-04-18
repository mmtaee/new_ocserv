package setup

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"ocserv/pkg/database"
)

type Controller struct {
	DB *gorm.DB
}

func New() *Controller {
	return &Controller{
		DB: database.Get(),
	}
}

// Setup		 Initializing and setup panel
//
// @Summary      Setup panel with admin user
// @Description  Setup panel with admin user and
// @Tags         Panel
// @Accept       json
// @Produce      json
// @Param        request    body  RequestSetup   true "setup config data"
// @Success      201  {object}  ResponseSetup
// @Router       /panel/setup/ [post]
func (ctrl *Controller) Setup(c echo.Context) error {
	return c.JSON(http.StatusCreated, nil)
}
