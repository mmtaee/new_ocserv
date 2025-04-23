package panel

import (
	"context"
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
		log.Println("cache db: ", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusOK, ConfigResponse{
				Setup:                false,
				GoogleCaptchaSiteKey: "",
			})
		}
		return ctrl.request.BadRequest(c, err)
	}
	log.Println("config:", config)
	return c.JSON(http.StatusOK, ConfigResponse{
		Setup:                true,
		GoogleCaptchaSiteKey: config.GoogleCaptchaSiteKey,
	})
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
	// TODO: check config ok or not

	var data RequestSetup
	if err := ctrl.request.DoValidate(c, &data); err != nil {
		return ctrl.request.BadRequest(c, err)
	}

	ctx := context.Background()

	go func() {
		_, err := ctrl.panelRepo.Create(ctx, &models.Panel{
			Setup:                  true,
			GoogleCaptchaSiteKey:   data.Config.GoogleCaptchaSiteKey,
			GoogleCaptchaSecretKey: data.Config.GoogleCaptchaSecretKey,
		})
		log.Println("create user panel:", err)
	}()

	go func() {
		err := ctrl.ocservGroupRepo.WithContext(ctx).UpdateDefaultGroup(data.DefaultOcservGroup)
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

// Login		 Admin and Staff users login
//
// @Summary      Admin and Staff users login
// @Description  Admin and Staff users login with Google captcha(captcha site key required in get config api)
// @Tags         Panel
// @Accept       json
// @Produce      json
// @Param        request    body  Login   true "setup config data"
// @Success      201  {object}  LoginResponse
// @Router       /panel/login/ [post]
func (ctrl *Controller) Login(c echo.Context) error {
	return c.JSON(http.StatusOK, LoginResponse{
		User:  nil,
		Token: "",
	})
}
