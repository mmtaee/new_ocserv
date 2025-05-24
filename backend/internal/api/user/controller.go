package user

import (
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"log"
	"net/http"
	"ocserv/internal/models"
	"ocserv/internal/repository"
	"ocserv/pkg/crypto"
	"ocserv/pkg/request"
	"strconv"
)

type Controller struct {
	request    request.CustomRequestInterface
	userRepo   repository.UserRepositoryInterface
	tokenRepo  repository.TokenRepositoryInterface
	cryptoRepo crypto.CustomPasswordInterface
}

func New() *Controller {
	return &Controller{
		request:    request.NewCustomRequest(),
		userRepo:   repository.NewUserRepository(),
		tokenRepo:  repository.NewTokenRepository(),
		cryptoRepo: crypto.NewCustomPassword(),
	}
}

// Profile		 Get user Profile
//
// @Summary      Get user Profile
// @Description  Get user Profile
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Bearer TOKEN"
// @Failure      400 {object} request.ErrorResponse
// @Failure      401 {object} middlewares.Unauthorized
// @Success      200  {object}  models.User
// @Router       /user/profile [get]
func (ctrl *Controller) Profile(c echo.Context) error {
	userUID := c.Get("userUID").(string)

	user, err := ctrl.userRepo.GetUserById(c.Request().Context(), userUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusUnauthorized, nil)
		}
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
// @Param        Authorization header string true "Bearer TOKEN"
// @Param        request    body  ChangePasswordData   true "change user password data"
// @Failure      400 {object} request.ErrorResponse
// @Failure      401 {object} middlewares.Unauthorized
// @Success      200  {object}  nil
// @Router       /user/password [post]
func (ctrl *Controller) ChangePassword(c echo.Context) error {
	var data ChangePasswordData
	if err := ctrl.request.DoValidate(c, &data); err != nil {
		return ctrl.request.BadRequest(c, err)
	}

	userUID := c.Get("userUID").(string)

	err := ctrl.userRepo.ChangePassword(c.Request().Context(), userUID, data.CurrentPassword, data.NewPassword)
	if err != nil {
		return ctrl.request.BadRequest(c, err)
	}
	return c.JSON(http.StatusOK, nil)
}

// Staffs 		 List of Staffs
//
// @Summary      List of Staffs
// @Description  List of Staffs
// @Tags         User
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
// @Success      200  {object}  StaffsResponse
// @Router       /user/staffs [get]
func (ctrl *Controller) Staffs(c echo.Context) error {
	pagination := ctrl.request.Pagination(c)

	log.Println("pagination: ", pagination)

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
// @Param        Authorization header string true "Bearer TOKEN"
// @Param        request    body  models.Permission   true "user permission"
// @Failure      400 {object} request.ErrorResponse
// @Failure      401 {object} middlewares.Unauthorized
// @Failure      403 {object} middlewares.PermissionDenied
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

// RemoveStaff 	 Remove Staff
//
// @Summary      Remove Staff
// @Description  Remove Staff by given id
// @Tags         User
// @Accept       json
// @Produce      json
// @Param 		 id path int true "User ID"
// @Param        Authorization header string true "Bearer TOKEN"
// @Failure      400 {object} request.ErrorResponse
// @Failure      401 {object} middlewares.Unauthorized
// @Failure      403 {object} middlewares.PermissionDenied
// @Success      204  {object}  nil
// @Router       /user/staffs/${id}/ [delete]
func (ctrl *Controller) RemoveStaff(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return ctrl.request.BadRequest(c, err)
	}

	err = ctrl.userRepo.Delete(c.Request().Context(), uint(id))
	if err != nil {
		return ctrl.request.BadRequest(c, err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

// ChangeStaffPassword 	 Change Staff Password
//
// @Summary      Change Staff Password
// @Description  Change Staff Password by given id
// @Tags         User
// @Accept       json
// @Produce      json
// @Param 		 id path int true "User ID"
// @Param        request    body  ChangeStaffPassword  true "user new password"
// @Param        Authorization header string true "Bearer TOKEN"
// @Failure      400 {object} request.ErrorResponse
// @Failure      401 {object} middlewares.Unauthorized
// @Failure      403 {object} middlewares.PermissionDenied
// @Success      200  {object}  nil
// @Router       /user/staffs/{id}/password [post]
func (ctrl *Controller) ChangeStaffPassword(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return ctrl.request.BadRequest(c, err)
	}

	var data ChangeStaffPassword
	if err := ctrl.request.DoValidate(c, &data); err != nil {
		return ctrl.request.BadRequest(c, err)
	}

	passwd := ctrl.cryptoRepo.CreatePassword(data.Password)

	err = ctrl.userRepo.ChangeStaffPassword(c.Request().Context(), uint(id), passwd.Hash, passwd.Salt)
	if err != nil {
		return ctrl.request.BadRequest(c, err)
	}
	return c.JSON(http.StatusOK, nil)
}

// CreateStaff 	 Create Admin or Staff
//
// @Summary      Create Admin or Staff
// @Description  Create Admin or Staff
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        request    body  CreateStaffData  true "user create data"
// @Param        Authorization header string true "Bearer TOKEN"
// @Failure      400 {object} request.ErrorResponse
// @Failure      401 {object} middlewares.Unauthorized
// @Failure      403 {object} middlewares.PermissionDenied
// @Success      201  {object}  models.User
// @Router       /user/staffs [post]
func (ctrl *Controller) CreateStaff(c echo.Context) error {
	var data CreateStaffData
	if err := ctrl.request.DoValidate(c, &data); err != nil {
		return ctrl.request.BadRequest(c, err)
	}
	passwd := ctrl.cryptoRepo.CreatePassword(data.Password)

	user := &models.User{
		Username:    data.Username,
		Password:    passwd.Hash,
		Salt:        passwd.Salt,
		IsAdmin:     false,
		Permissions: &data.Permission,
	}

	staff, err := ctrl.userRepo.CreateStaff(c.Request().Context(), user)
	if err != nil {
		return ctrl.request.BadRequest(c, err)
	}
	return c.JSON(http.StatusCreated, staff)
}
