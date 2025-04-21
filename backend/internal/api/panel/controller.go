package panel

import (
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"log"
	"net/http"
	"ocserv/internal/models"
	"ocserv/internal/repository"
	"ocserv/pkg/crypto"
	"ocserv/pkg/database"
	"ocserv/pkg/oc"
	"ocserv/pkg/request"
)

type Controller struct {
	DB              *gorm.DB
	request         request.CustomRequestInterface
	userRepo        repository.UserRepositoryInterface
	tokenRepo       repository.TokenRepositoryInterface
	panelRepo       repository.PanelRepositoryInterface
	ocservGroupRepo oc.OcservGroupServiceInterface
}

func New() *Controller {
	return &Controller{
		DB:              database.Get(),
		request:         request.NewCustomRequest(),
		userRepo:        repository.NewUserRepository(),
		tokenRepo:       repository.NewTokenRepository(),
		panelRepo:       repository.NewPanelRepository(),
		ocservGroupRepo: oc.NewOcGroupService(),
	}
}

// Config		 Panel Config
//
// @Summary      Get panel Config
// @Description  Get panel Config
// @Tags         Panel
// @Accept       json
// @Produce      json
// @Success      200  {object}  ConfigResponse
// @Router       /panel [get]
func (ctrl *Controller) Config(c echo.Context) error {
	config, err := ctrl.panelRepo.GetConfig(c.Request().Context())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusOK, ConfigResponse{
				Setup:                false,
				GoogleCaptchaSiteKey: "",
			})
		}
		return ctrl.request.BadRequest(c, err)
	}
	return c.JSON(http.StatusOK, config)
}

// Setup		 Initializing and setup panel
//
// @Summary      Setup panel with admin user
// @Description  Setup panel with admin user, captcha and ocserv default group configs
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

	go func() {
		_, err := ctrl.panelRepo.Create(c.Request().Context(), &models.Panel{
			Setup:                  true,
			GoogleCaptchaSiteKey:   data.Config.GoogleCaptchaSiteKey,
			GoogleCaptchaSecretKey: data.Config.GoogleCaptchaSecretKey,
		})
		log.Println("create user panel:", err)
	}()

	go func() {
		err := ctrl.ocservGroupRepo.WithContext(c.Request().Context()).UpdateDefaultGroup(data.DefaultOcservGroup)
		if err != nil {
			log.Println("update default group:", err)
		}
	}()

	passwordPKG := crypto.CreatePassword(data.Config.AdminPassword)
	user, err := ctrl.userRepo.CreateAdmin(c.Request().Context(), &models.User{
		Username: data.Config.AdminUsername,
		Password: passwordPKG.Hash,
		Salt:     passwordPKG.Salt,
		IsAdmin:  true,
	})
	if err != nil {
		return ctrl.request.BadRequest(c, err)
	}

	token, err := ctrl.tokenRepo.CreateToken(c.Request().Context(), user.ID, user.UID, true)
	if err != nil {
		return ctrl.request.BadRequest(c, err, "user created")
	}

	return c.JSON(
		http.StatusCreated,
		ResponseSetup{
			User:  user,
			Token: token,
		},
	)
}
