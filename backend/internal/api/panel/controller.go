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
	"ocserv/pkg/oc"
	"ocserv/pkg/request"
	"ocserv/pkg/utils/captcha"
	"time"
)

type Controller struct {
	request         request.CustomRequestInterface
	userRepo        repository.UserRepositoryInterface
	tokenRepo       repository.TokenRepositoryInterface
	panelRepo       repository.PanelRepositoryInterface
	ocservGroupRepo oc.OcservGroupServiceInterface
}

func New() *Controller {
	return &Controller{
		request:         request.NewCustomRequest(),
		userRepo:        repository.NewUserRepository(),
		tokenRepo:       repository.NewTokenRepository(),
		panelRepo:       repository.NewPanelRepository(),
		ocservGroupRepo: oc.NewOcGroupService(),
	}
}

// Init		 	Panel Init Config
//
// @Summary      Get panel Init Config
// @Description  Get panel Init Config
// @Tags         Panel
// @Accept       json
// @Produce      json
// @Success      200  {object}  InitResponse
// @Router       /panel/init [get]
func (ctrl *Controller) Init(c echo.Context) error {
	config, err := ctrl.panelRepo.GetConfig(c.Request().Context())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusOK, InitResponse{
				Setup:                false,
				GoogleCaptchaSiteKey: "",
			})
		}
		return ctrl.request.BadRequest(c, err)
	}
	return c.JSON(http.StatusOK, InitResponse{
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
// @Param        request    body  SetupData   true "setup config data"
// @Success      201  {object}  UserResponse
// @Router       /panel/setup/ [post]
func (ctrl *Controller) Setup(c echo.Context) error {
	var data SetupData
	if err := ctrl.request.DoValidate(c, &data); err != nil {
		return ctrl.request.BadRequest(c, err)
	}

	config, err := ctrl.panelRepo.GetConfig(c.Request().Context())
	if err == nil || config != nil {
		return ctrl.request.BadRequest(c, errors.New("config db exists"))
	}

	ctxPanel, cancel1 := context.WithTimeout(context.Background(), 10*time.Second)
	go func() {
		defer cancel1()

		_, err = ctrl.panelRepo.Create(ctxPanel, &models.Panel{
			Setup:                  true,
			GoogleCaptchaSiteKey:   data.Config.GoogleCaptchaSiteKey,
			GoogleCaptchaSecretKey: data.Config.GoogleCaptchaSecretKey,
		})
		if err != nil {
			log.Println("setup err:", err)
		}
	}()

	ctxGroup, cancel2 := context.WithTimeout(context.Background(), 10*time.Second)
	go func() {
		defer cancel2()

		err = ctrl.ocservGroupRepo.WithContext(ctxGroup).UpdateDefaultGroup(data.DefaultOcservGroup)
		if err != nil {
			log.Println("update default group:", err)
		}
	}()

	passwordPKG := crypto.CreatePassword(data.Admin.Password)
	user, err := ctrl.userRepo.CreateAdmin(c.Request().Context(), &models.User{
		Username: data.Admin.Username,
		Password: passwordPKG.Hash,
		Salt:     passwordPKG.Salt,
		IsAdmin:  true,
	})
	if err != nil {
		return ctrl.request.BadRequest(c, err)
	}

	token, err := ctrl.tokenRepo.CreateToken(c.Request().Context(), user.ID, user.UID, true, user.IsAdmin)
	if err != nil {
		return ctrl.request.BadRequest(c, err, "user created")
	}

	return c.JSON(
		http.StatusCreated,
		UserResponse{
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
// @Success      201  {object}  UserResponse
// @Router       /panel/login/ [post]
func (ctrl *Controller) Login(c echo.Context) error {
	var data Login
	if err := ctrl.request.DoValidate(c, &data); err != nil {
		return ctrl.request.BadRequest(c, err)
	}

	config, err := ctrl.panelRepo.GetConfig(c.Request().Context())
	if err != nil {
		return ctrl.request.BadRequest(c, err)
	}

	if secretKey := config.GoogleCaptchaSecretKey; secretKey != "" {
		res := captcha.Verify(secretKey, data.Token)
		if !res {
			return ctrl.request.BadRequest(c, errors.New("captcha challenge failed"))
		}
	}

	user, err := ctrl.userRepo.GetUserByUsername(c.Request().Context(), data.Username)
	if err != nil {
		return ctrl.request.BadRequest(c, errors.New("invalid username or password"))
	}

	token, err := ctrl.tokenRepo.CreateToken(c.Request().Context(), user.ID, user.UID, true, user.IsAdmin)
	if err != nil {
		return ctrl.request.BadRequest(c, err, "user created")
	}

	return c.JSON(http.StatusOK, UserResponse{
		User:  user,
		Token: token,
	})
}

// Config		 Panel Config
//
// @Summary      Get panel Config
// @Description  Get panel Config
// @Tags         Panel
// @Accept       json
// @Produce      json
// @Success      200  {object}  ConfigResponse
// @Router       /panel/config [get]
func (ctrl *Controller) Config(c echo.Context) error {
	config, err := ctrl.panelRepo.GetConfig(c.Request().Context())
	if err != nil {
		return ctrl.request.BadRequest(c, err)
	}
	return c.JSON(http.StatusOK, ConfigResponse{
		GoogleCaptchaSiteKey:   config.GoogleCaptchaSiteKey,
		GoogleCaptchaSecretKey: config.GoogleCaptchaSecretKey,
	})
}

// UpdateConfig		 Panel Config Updating
//
// @Summary      Update Config panel
// @Description  Update Config panel
// @Tags         Panel
// @Accept       json
// @Produce      json
// @Param        request    body  ConfigData   true "setup config data"
// @Success      200  {object}  ConfigResponse
// @Router       /panel/config [patch]
func (ctrl *Controller) UpdateConfig(c echo.Context) error {
	var data ConfigData
	if err := ctrl.request.DoValidate(c, &data); err != nil {
		return ctrl.request.BadRequest(c, err)
	}

	panel := &models.Panel{}

	if data.GoogleCaptchaSecretKey != nil {
		panel.GoogleCaptchaSecretKey = *data.GoogleCaptchaSecretKey
	}
	if data.GoogleCaptchaSiteKey != nil {
		panel.GoogleCaptchaSiteKey = *data.GoogleCaptchaSiteKey
	}

	config, err := ctrl.panelRepo.Update(c.Request().Context(), panel)
	if err != nil {
		return ctrl.request.BadRequest(c, err)
	}
	return c.JSON(http.StatusOK, ConfigResponse{
		GoogleCaptchaSiteKey:   config.GoogleCaptchaSiteKey,
		GoogleCaptchaSecretKey: config.GoogleCaptchaSecretKey,
	})
}
