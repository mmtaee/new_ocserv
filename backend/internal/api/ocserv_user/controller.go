package ocservUser

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"ocserv/internal/repository"
	"ocserv/pkg/oc"
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
// @Failure      400 {object} request.ErrorResponse
// @Failure      401 {object} middlewares.Unauthorized
// @Failure      403 {object} middlewares.PermissionDenied
// @Success      200  {object}  OcservUsersResponse
// @Router       /oc_users [get]
func (ctrl *Controller) GetUsers(c echo.Context) error {
	pagination := ctrl.request.Pagination()
	ocUsers, count, err := ctrl.ocservUserRepo.GetUsersWithOnlineAttr(c.Request().Context(), pagination)
	if err != nil {
		return ctrl.request.BadRequest(c, err)
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

// CreateUser 	 Create Ocserv Users
//
// @Summary      Create Ocserv Users
// @Description  Create Ocserv Users
// @Tags         Ocserv Users
// @Accept       json
// @Produce      json
// @Param        request    body  createOcservUserData   true "create ocserv user data"
// @Param        Authorization header string true "Bearer TOKEN"
// @Failure      400 {object} request.ErrorResponse
// @Failure      401 {object} middlewares.Unauthorized
// @Failure      403 {object} middlewares.PermissionDenied
// @Success      200  {object}  oc.OcservUser
// @Router       /oc_users [post]
func (ctrl *Controller) CreateUser(c echo.Context) error {
	var data createOcservUserData
	if err := c.Bind(&data); err != nil {
		return ctrl.request.BadRequest(c, err)
	}

	if data.TrafficSize == 0 {
		data.TrafficSize = 10 * 1024 * 1024 * 1024
	}

	ocUser := oc.OcservUser{
		Group:       data.Group,
		Username:    data.Username,
		Password:    data.Password,
		ExpireAt:    data.ExpireAt,
		TrafficType: data.TrafficType,
		TrafficSize: data.TrafficSize,
		Description: data.Description,
	}

	ocUserCreated, err := ctrl.ocservUserRepo.CreateUser(c.Request().Context(), &ocUser)
	if err != nil {
		return ctrl.request.BadRequest(c, err)
	}

	return c.JSON(http.StatusCreated, ocUserCreated)
}
