package ocservUser

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"ocserv/internal/repository"
	"ocserv/pkg/request"
)

type Controller struct {
	request        request.CustomRequestInterface
	ocservUserRepo repository.OcservUserRepositoryInterface
}

func New() *Controller {
	return &Controller{
		request:        request.NewCustomRequest(),
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
// @Failure      401 {object} middlewares.Unauthorized
// @Failure      403 {object} middlewares.PermissionDenied
// @Success      200  {object}  OcservUsersResponse
// @Router       /oc_users [get]
func (ctrl *Controller) GetUsers(c echo.Context) error {
	pagination := ctrl.request.Pagination()
	ocUsers, count, err := ctrl.ocservUserRepo.GetUsersWithOnlineAttr(c.Request().Context(), pagination)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, OcservUsersResponse{
		Meta: request.Meta{
			Page:         pagination.Page,
			PageSize:     pagination.PageSize,
			TotalRecords: count,
		},
		Result: ocUsers,
	})
}
