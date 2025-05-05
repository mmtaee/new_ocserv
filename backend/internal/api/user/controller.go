package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"ocserv/internal/repository"
	"ocserv/pkg/request"
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
