package ocservUser

import (
	"github.com/labstack/echo/v4"
	"ocserv/internal/repository"
)

type Controller struct {
	ocservUserRepo repository.OcservUserRepositoryInterface
}

func New() *Controller {
	return &Controller{
		ocservUserRepo: repository.NewOcservUserRepository(),
	}
}

// GetUsers 	 List of Ocserv Users
//
// @Summary      List of Ocserv Users
// @Description  List of Ocserv Users
// @Tags         Ocserv Users
// @Accept       json
// @Produce      json
// @Param page query int false "Page number, starting from 1" minimum(1)
// @Param page_size query int false "Number of items per page" minimum(1) maximum(100)
// @Param order query string false "Field to order by"
// @Param sort query string false "Sort order, either ASC or DESC" Enums(ASC, DESC)
// @Param        Authorization header string true "Bearer TOKEN"
// @Success      200  {object}  OcservUsersResponse
// @Failure      401 {object} middlewares.Unauthorized
// @Router       /oc_users [get]
func (ctrl *Controller) GetUsers(c echo.Context) error {
	return nil
}
