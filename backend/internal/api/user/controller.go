package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"ocserv/internal/models"
	"ocserv/internal/repository"
	"ocserv/pkg/request"
	"strconv"
)

type Controller struct {
	request   request.CustomRequestInterface
	userRepo  repository.UserRepositoryInterface
	tokenRepo repository.TokenRepositoryInterface
}

func New() *Controller {
	return &Controller{
		request:   request.NewCustomRequest(),
		userRepo:  repository.NewUserRepository(),
		tokenRepo: repository.NewTokenRepository(),
	}
}

// Profile		 Get user Profile
//
// @Summary      Get user Profile
// @Description  Get user Profile
// @Tags         User
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.User
// @Router       /user/profile [get]
func (ctrl *Controller) Profile(c echo.Context) error {
	userID := c.Get("userID").(string)

	user, err := ctrl.userRepo.GetUserById(c.Request().Context(), userID)
	if err != nil {
		return ctrl.request.BadRequest(c, err)
	}
	return c.JSON(http.StatusOK, user)
}

// ChangePassword User Change Password
//
// @Summary      User Change Password
// @Description  User Change Password
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        request    body  ChangePasswordData   true "change user password data"
// @Success      200  {object}  nil
// @Router       /user/password [post]
func (ctrl *Controller) ChangePassword(c echo.Context) error {
	var data ChangePasswordData
	if err := ctrl.request.DoValidate(c, &data); err != nil {
		return ctrl.request.BadRequest(c, err)
	}

	userID := c.Get("userID").(string)

	err := ctrl.userRepo.ChangePassword(c.Request().Context(), userID, data.CurrentPassword, data.NewPassword)
	if err != nil {
		return ctrl.request.BadRequest(c, err)
	}
	return c.JSON(http.StatusOK, nil)
}

// Staffs 		 List os Staffs
//
// @Summary      List os Staffs
// @Description  List os Staffs
// @Tags         User
// @Accept       json
// @Produce      json
// @Param page query int false "Page number, starting from 1" minimum(1)
// @Param page_size query int false "Number of items per page" minimum(1) maximum(100)
// @Param order query string false "Field to order by"
// @Param sort query string false "Sort order, either ASC or DESC" Enums(ASC, DESC)
// @Success      200  {object}  StaffsResponse
// @Router       /user/staffs [get]
func (ctrl *Controller) Staffs(c echo.Context) error {
	pagination := ctrl.request.Pagination()

	if err := ctrl.request.DoValidate(c, pagination); err != nil {
		return ctrl.request.BadRequest(c, err)
	}

	staffs, total, err := ctrl.userRepo.GetStaffs(c.Request().Context(), pagination)
	if err != nil {
		return ctrl.request.BadRequest(c, err)
	}

	return c.JSON(http.StatusOK, StaffsResponse{
		Meta: request.Meta{
			Page:         pagination.Page,
			PageSize:     pagination.PageSize,
			TotalRecords: total,
		},
		Result: staffs,
	})
}

// UpdateStaffPermission 		 Update Staff Permission
//
// @Summary      Update Staff Permission
// @Description  Update Staff Permission by given id
// @Tags         User
// @Accept       json
// @Produce      json
// @Param 		 id path int true "User ID"
// @Param        request    body  models.Permission   true "user permission"
// @Success      200  {object}  nil
// @Router       /user/staffs/permissions/${id}/ [put]
func (ctrl *Controller) UpdateStaffPermission(c echo.Context) error {
	var perm models.Permission

	if err := ctrl.request.DoValidate(c, &perm); err != nil {
		return ctrl.request.BadRequest(c, err)
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return ctrl.request.BadRequest(c, err)
	}

	err = ctrl.userRepo.UpdatePermission(c.Request().Context(), uint(id), perm)
	if err != nil {
		return ctrl.request.BadRequest(c, err)
	}
	return c.JSON(http.StatusOK, nil)
}
