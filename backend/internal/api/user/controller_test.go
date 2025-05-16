package user

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"ocserv/internal/models"
	"ocserv/mocks"
	"ocserv/pkg/crypto"
	"ocserv/pkg/request"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupEcho(method, path, body string) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	r := httptest.NewRequest(method, path, strings.NewReader(body))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	c := e.NewContext(r, w)
	return c, w
}

func TestProfileSuccess(t *testing.T) {
	mockRequest := new(mocks.CustomRequestInterface)
	mockUserRepo := new(mocks.UserRepositoryInterface)
	mockTokenRepo := new(mocks.TokenRepositoryInterface)

	ctrl := &Controller{
		request:   mockRequest,
		userRepo:  mockUserRepo,
		tokenRepo: mockTokenRepo,
	}

	expectedUser := &models.User{Username: "testuser"}
	c, w := setupEcho(http.MethodGet, "/user/profile", "")
	c.Set("userUID", "1")

	mockUserRepo.On("GetUserById", mock.Anything, "1").Return(expectedUser, nil)

	err := ctrl.Profile(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	mockUserRepo.AssertExpectations(t)
}

func TestProfileError(t *testing.T) {
	mockRequest := new(mocks.CustomRequestInterface)
	mockUserRepo := new(mocks.UserRepositoryInterface)
	mockTokenRepo := new(mocks.TokenRepositoryInterface)

	ctrl := &Controller{
		request:   mockRequest,
		userRepo:  mockUserRepo,
		tokenRepo: mockTokenRepo,
	}

	c, _ := setupEcho(http.MethodGet, "/user/profile", "")
	c.Set("userUID", "1")

	expectedErr := errors.New("user not found")
	mockUserRepo.On("GetUserById", mock.Anything, "1").Return(nil, expectedErr)
	mockRequest.On("BadRequest", mock.Anything, expectedErr).Return(nil)

	err := ctrl.Profile(c)

	assert.NoError(t, err)
	mockUserRepo.AssertExpectations(t)
	mockRequest.AssertExpectations(t)
}

func TestChangePasswordValidationFail(t *testing.T) {
	mockRequest := new(mocks.CustomRequestInterface)
	mockUserRepo := new(mocks.UserRepositoryInterface)
	ctrl := &Controller{request: mockRequest, userRepo: mockUserRepo}

	c, _ := setupEcho(http.MethodPost, "/user/password", "{}")
	err := errors.New("validation failed")
	mockRequest.On("DoValidate", c, mock.Anything).Return(err)
	mockRequest.On("BadRequest", c, err).Return(nil)

	e := ctrl.ChangePassword(c)
	assert.NoError(t, e)
}

func TestChangePasswordSuccess(t *testing.T) {
	mockRequest := new(mocks.CustomRequestInterface)
	mockUserRepo := new(mocks.UserRepositoryInterface)
	ctrl := &Controller{request: mockRequest, userRepo: mockUserRepo}

	c, w := setupEcho(http.MethodPost, "/user/password", `{"current_password":"old","new_password":"new"}`)
	c.Set("userUID", "123")

	// Mock validation to populate the struct with values
	mockRequest.On("DoValidate", mock.Anything, mock.AnythingOfType("*user.ChangePasswordData")).Run(func(args mock.Arguments) {
		arg := args.Get(1).(*ChangePasswordData)
		arg.CurrentPassword = "old"
		arg.NewPassword = "new"
	}).Return(nil)

	mockUserRepo.On("ChangePassword", mock.Anything, "123", "old", "new").Return(nil)

	err := ctrl.ChangePassword(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	mockUserRepo.AssertExpectations(t)
}

func TestStaffsSuccess(t *testing.T) {
	mockRequest := new(mocks.CustomRequestInterface)
	mockUserRepo := new(mocks.UserRepositoryInterface)
	pagination := &request.Pagination{Page: 1, PageSize: 10}

	ctrl := &Controller{request: mockRequest, userRepo: mockUserRepo}
	mockRequest.On("Pagination").Return(pagination)
	mockRequest.On("DoValidate", mock.Anything, pagination).Return(nil)

	mockUserRepo.On("GetStaffs", mock.Anything, pagination).Return(&[]models.User{}, int64(0), nil)

	c, w := setupEcho(http.MethodGet, "/user/staffs", "")
	err := ctrl.Staffs(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateStaffSuccess(t *testing.T) {
	mockRequest := new(mocks.CustomRequestInterface)
	mockUserRepo := new(mocks.UserRepositoryInterface)
	mockHasher := new(mocks.CustomPasswordInterface)

	ctrl := &Controller{request: mockRequest, userRepo: mockUserRepo, cryptoRepo: mockHasher}
	input := `{"username":"staff1","password":"secret","permission":{}}`
	staff := &models.User{Username: "staff1", Password: "password"}

	mockRequest.On("DoValidate", mock.Anything, mock.AnythingOfType("*user.CreateStaffData")).
		Run(func(args mock.Arguments) {
			data := args.Get(1).(*CreateStaffData)
			data.Username = "staff1"
			data.Password = "secret"
		}).Return(nil)

	mockHasher.On("CreatePassword", "secret").Return(crypto.CustomPassword{
		Hash: "hash",
		Salt: "salt",
	})

	mockUserRepo.On("CreateStaff", mock.Anything, mock.MatchedBy(func(u *models.User) bool {
		return u.Username == "staff1" && u.Password == "hash" && u.Salt == "salt"
	})).Return(staff, nil)

	c, w := setupEcho(http.MethodPost, "/user/staffs", input)
	err := ctrl.CreateStaff(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestUpdateStaffPermissionSuccess(t *testing.T) {
	mockRequest := new(mocks.CustomRequestInterface)
	mockUserRepo := new(mocks.UserRepositoryInterface)

	ctrl := &Controller{request: mockRequest, userRepo: mockUserRepo}
	mockRequest.On("DoValidate", mock.Anything, mock.Anything).Return(nil)
	mockUserRepo.On("UpdatePermission", mock.Anything, uint(1), mock.Anything).Return(nil)

	c, w := setupEcho(http.MethodPut, "/user/staffs/permissions/1/", `{"can_manage":true}`)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := ctrl.UpdateStaffPermission(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRemoveStaffSuccess(t *testing.T) {
	mockRequest := new(mocks.CustomRequestInterface)
	mockUserRepo := new(mocks.UserRepositoryInterface)

	ctrl := &Controller{request: mockRequest, userRepo: mockUserRepo}
	mockUserRepo.On("Delete", mock.Anything, uint(1)).Return(nil)

	c, w := setupEcho(http.MethodDelete, "/user/staffs/1/", "")
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := ctrl.RemoveStaff(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestChangeStaffPasswordSuccess(t *testing.T) {
	mockRequest := new(mocks.CustomRequestInterface)
	mockUserRepo := new(mocks.UserRepositoryInterface)
	mockHasher := new(mocks.CustomPasswordInterface)

	ctrl := &Controller{request: mockRequest, userRepo: mockUserRepo, cryptoRepo: mockHasher}
	mockRequest.On("DoValidate", mock.Anything, mock.AnythingOfType("*user.ChangeStaffPassword")).
		Run(func(args mock.Arguments) {
			data := args.Get(1).(*ChangeStaffPassword)
			data.Password = "newpass"
		}).Return(nil)

	mockHasher.On("CreatePassword", "newpass").Return(crypto.CustomPassword{
		Hash: "newpass",
		Salt: "salt",
	})

	mockUserRepo.On("ChangeStaffPassword", mock.Anything, uint(1), "newpass", "salt").Return(nil)

	c, w := setupEcho(http.MethodPost, "/user/staffs/1/password", `{"password":"newpass"}`)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := ctrl.ChangeStaffPassword(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
}
