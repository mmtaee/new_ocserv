package panel

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"ocserv/internal/models"
	"ocserv/mocks"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func setupEcho(method, path string, body string) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return c, rec
}

func TestControllerInitConfigExists(t *testing.T) {
	panelRepo := new(mocks.PanelRepositoryInterface)
	mockRequest := new(mocks.CustomRequestInterface)

	ctrl := &Controller{
		request:   mockRequest,
		panelRepo: panelRepo,
	}

	panelRepo.On("GetConfig", mock.Anything).Return(&models.Panel{
		GoogleCaptchaSiteKey: "test-key",
	}, nil)

	c, rec := setupEcho(http.MethodGet, "/panel/init", "")
	err := ctrl.Init(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), `"setup":true`)
	assert.Contains(t, rec.Body.String(), `"google_captcha_site_key":"test-key"`)
}

func TestControllerInitConfigNotFound(t *testing.T) {
	panelRepo := new(mocks.PanelRepositoryInterface)
	mockRequest := new(mocks.CustomRequestInterface)

	ctrl := &Controller{
		request:   mockRequest,
		panelRepo: panelRepo,
	}

	panelRepo.On("GetConfig", mock.Anything).Return(nil, gorm.ErrRecordNotFound)

	c, rec := setupEcho(http.MethodGet, "/panel/init", "")
	err := ctrl.Init(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), `"setup":false`)
	assert.Contains(t, rec.Body.String(), `"google_captcha_site_key":""`)
}

func TestControllerSetupConfigAlreadyExists(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/panel/setup", strings.NewReader(`{}`))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	panelRepo := new(mocks.PanelRepositoryInterface)
	mockRequest := new(mocks.CustomRequestInterface)

	ctrl := &Controller{
		request:   mockRequest,
		panelRepo: panelRepo,
	}

	expectedErr := errors.New("config db exists")

	mockRequest.On("DoValidate", mock.Anything, mock.Anything).Return(nil)
	panelRepo.On("GetConfig", mock.Anything).Return(&models.Panel{}, nil)
	mockRequest.On("BadRequest", mock.Anything, expectedErr).Return(c.JSON(http.StatusBadRequest, map[string]string{"error": expectedErr.Error()}))

	err := ctrl.Setup(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "config db exists")

	panelRepo.AssertCalled(t, "GetConfig", mock.Anything)
	mockRequest.AssertCalled(t, "BadRequest", mock.Anything, expectedErr)
}

func TestControllerConfigSuccess(t *testing.T) {
	panelRepo := new(mocks.PanelRepositoryInterface)
	mockRequest := new(mocks.CustomRequestInterface)

	ctrl := &Controller{
		request:   mockRequest,
		panelRepo: panelRepo,
	}

	panelRepo.On("GetConfig", mock.Anything).Return(&models.Panel{
		GoogleCaptchaSiteKey:   "site",
		GoogleCaptchaSecretKey: "secret",
	}, nil)

	c, rec := setupEcho(http.MethodGet, "/panel/config", "")
	err := ctrl.Config(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), `"google_captcha_site_key":"site"`)
	assert.Contains(t, rec.Body.String(), `"google_captcha_secret_key":"secret"`)
}

func TestControllerUpdateConfigSuccess(t *testing.T) {
	panelRepo := new(mocks.PanelRepositoryInterface)
	mockRequest := new(mocks.CustomRequestInterface)

	ctrl := &Controller{
		request:   mockRequest,
		panelRepo: panelRepo,
	}

	mockRequest.On("DoValidate", mock.Anything, mock.Anything).Return(nil)

	panelRepo.On("Update", mock.Anything, mock.Anything).Return(&models.Panel{
		GoogleCaptchaSiteKey:   "updated-site",
		GoogleCaptchaSecretKey: "updated-secret",
	}, nil)

	json := `{"google_captcha_site_key":"updated-site","google_captcha_secret_key":"updated-secret"}`
	c, rec := setupEcho(http.MethodPatch, "/panel/config", json)

	err := ctrl.UpdateConfig(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), `"google_captcha_site_key":"updated-site"`)
	assert.Contains(t, rec.Body.String(), `"google_captcha_secret_key":"updated-secret"`)
}
