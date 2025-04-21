package setup

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"ocserv/pkg/database"
	"ocserv/pkg/request"
)

type Controller struct {
	DB      *gorm.DB
	request request.CustomRequestInterface
}

func New() *Controller {
	return &Controller{
		DB:      database.Get(),
		request: request.NewCustomRequest(),
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
	var data RequestSetup
	if err := ctrl.request.DoValidate(c, &data); err != nil {
		return ctrl.request.BadRequest(c, err)
	}
	return c.JSON(http.StatusCreated, nil)
}
