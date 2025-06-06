package ocservUser

import (
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"ocserv/internal/repository"
	"ocserv/pkg/oc"
	"ocserv/pkg/request"
	"slices"
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

// Update		 Ocserv User update
//
// @Summary      Ocserv User update
// @Description  Ocserv User update
// @Tags         Ocserv Users
// @Accept       json
// @Produce      json
// @Param 		 uid path string true "Ocserv User UID"
// @Param        Authorization header string true "Bearer TOKEN"
// @Param        request body  UpdateOcservUserData   true "update ocserv user data"
// @Failure      400 {object} request.ErrorResponse
// @Failure      401 {object} middlewares.Unauthorized
// @Failure      403 {object} middlewares.PermissionDenied
// @Success      200  {object}  oc.OcservUser
// @Router       /oc_users/{uid} [patch]
func (ctrl *Controller) Update(c echo.Context) error {
	var data UpdateOcservUserData
	if err := ctrl.request.DoValidate(c, &data); err != nil {
		return ctrl.request.BadRequest(c, err)
	}

	userUID := c.Param("uid")

	ocUser, err := ctrl.ocservUserRepo.GetUserByID(c.Request().Context(), userUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, nil)
		}
		return ctrl.request.BadRequest(c, err)
	}
	if data.Group != nil {
		ocUser.Group = *data.Group
	}
	if data.Password != nil {
		ocUser.Password = *data.Password
	}
	if data.Description != nil {
		ocUser.Description = *data.Description
	}
	if data.TrafficSize != nil {
		ocUser.TrafficSize = *data.TrafficSize
	}
	if data.TrafficType != nil && slices.Contains([]string{"Free", "MonthlyTransmit", "MonthlyReceive", "TotallyTransmit", "TotallyReceive"}, *data.TrafficType) {
		ocUser.TrafficType = *data.TrafficType
	}

	ocUser, err = ctrl.ocservUserRepo.UpdateUser(c.Request().Context(), ocUser)
	if err != nil {
		return ctrl.request.BadRequest(c, err)
	}
	return c.JSON(http.StatusOK, ocUser)
}

// Delete		 Ocserv User delete
//
// @Summary      Ocserv User delete
// @Description  Ocserv User delete
// @Tags         Ocserv Users
// @Accept       json
// @Produce      json
// @Param 		 uid path string true "Ocserv User UID"
// @Param        Authorization header string true "Bearer TOKEN"
// @Failure      400 {object} request.ErrorResponse
// @Failure      401 {object} middlewares.Unauthorized
// @Failure      403 {object} middlewares.PermissionDenied
// @Success      204  {object} nil
// @Router       /oc_users/{uid} [delete]
func (ctrl *Controller) Delete(c echo.Context) error {
	userUID := c.Param("uid")
	err := ctrl.ocservUserRepo.DeleteUser(c.Request().Context(), userUID)
	if err != nil {
		return ctrl.request.BadRequest(c, err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

// Disconnect	 Ocserv Users disconnect
//
// @Summary      Ocserv Users disconnect
// @Description  Ocserv Users disconnect
// @Tags         Ocserv Users
// @Accept       json
// @Produce      json
// @Param 		 username path string true "Ocserv Username"
// @Param        Authorization header string true "Bearer TOKEN"
// @Failure      400 {object} request.ErrorResponse
// @Failure      401 {object} middlewares.Unauthorized
// @Failure      403 {object} middlewares.PermissionDenied
// @Success      200  {object}  nil
// @Router       /oc_users/{username}/disconnect [post]
func (ctrl *Controller) Disconnect(c echo.Context) error {
	userUID := c.Param("username")
	err := ctrl.ocservUserRepo.DisconnectUser(c.Request().Context(), userUID)
	if err != nil {
		return ctrl.request.BadRequest(c, err)
	}
	return c.JSON(http.StatusOK, nil)
}
