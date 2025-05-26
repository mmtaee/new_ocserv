package ocservUser

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"ocserv/internal/repository"
	"ocserv/pkg/oc"
	"ocserv/pkg/request"
	"time"
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
// @Param 		 page query int false "Page number, starting from 1" minimum(1)
// @Param 		 page_size query int false "Number of items per page" minimum(1) maximum(100)
// @Param 		 order query string false "Field to order by"
// @Param 		 sort query string false "Sort order, either ASC or DESC" Enums(ASC, DESC)
// @Param        Authorization header string true "Bearer TOKEN"
// @Failure      400 {object} request.ErrorResponse
// @Failure      401 {object} middlewares.Unauthorized
// @Failure      403 {object} middlewares.PermissionDenied
// @Success      200  {object}  OcservUsersResponse
// @Router       /oc_users [get]
func (ctrl *Controller) GetUsers(c echo.Context) error {
	pagination := ctrl.request.Pagination(c)

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

	// TODO: check existence of  group

	if data.TrafficSize == 0 {
		data.TrafficSize = 10 * 1024 * 1024 * 1024
	}
	parsed, err := time.Parse("2006-01-02", data.ExpireAt)
	if err != nil {
		parsed, _ = time.Parse(
			"2006-01-02",
			time.Now().AddDate(0, 0, 30).Format("2006-01-02"),
		)
	}

	ocUser := oc.OcservUser{
		Group:       data.Group,
		Username:    data.Username,
		Password:    data.Password,
		ExpireAt:    &parsed,
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

// Lock			 Unlock Or Lock Ocserv Users
//
// @Summary      Lock Or Unlock Ocserv Users
// @Description  Lock Or Unlock Ocserv Users
// @Tags         Ocserv Users
// @Accept       json
// @Produce      json
// @Param 		 uid path string true "Ocserv User UID"
// @Param        Authorization header string true "Bearer TOKEN"
// @Param        request body  LockOcservUserData   true "lock or unlock ocserv user data"
// @Failure      400 {object} request.ErrorResponse
// @Failure      401 {object} middlewares.Unauthorized
// @Failure      403 {object} middlewares.PermissionDenied
// @Success      200  {object}  nil
// @Router       /oc_users/{uid}/lock [post]
func (ctrl *Controller) Lock(c echo.Context) error {
	var data LockOcservUserData
	if err := ctrl.request.DoValidate(c, &data); err != nil {
		return ctrl.request.BadRequest(c, err)
	}

	userUID := c.Param("uid")
	err := ctrl.ocservUserRepo.LockUser(c.Request().Context(), userUID, *data.Lock)
	if err != nil {
		return ctrl.request.BadRequest(c, err)
	}

	return c.JSON(http.StatusNoContent, nil)
}

func (ctrl *Controller) Update(c echo.Context) error {
	var data UpdateOcservUserData
	if err := ctrl.request.DoValidate(c, &data); err != nil {
		return ctrl.request.BadRequest(c, err)
	}

	//userUID := c.Param("uid")
	return nil
}

func (ctrl *Controller) Delete(c echo.Context) error {
	//userUID := c.Param("uid")
	return nil
}

func (ctrl *Controller) Disconnect(c echo.Context) error {
	// 	userUID := c.Param("username")
	return nil
}
