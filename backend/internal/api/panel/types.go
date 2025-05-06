package panel

import (
	"ocserv/internal/models"
	"ocserv/pkg/oc"
)

type SetupData struct {
	Admin struct {
		Username string `json:"username" validate:"required,min=2,max=16" example:"john_doe" `
		Password string `json:"password" validate:"required,min=2,max=16" example:"doe123456"`
	} `json:"admin" validate:"required"`
	Config struct {
		GoogleCaptchaSecretKey string `json:"google_captcha_secret_key" validate:"required"`
		GoogleCaptchaSiteKey   string `json:"google_captcha_site_key" validate:"required"`
	} `json:"config" validate:"omitempty"`
	DefaultOcservGroup *oc.OcservDefaultConfigs `json:"default_ocserv_group" validate:"required"`
}

type UserResponse struct {
	User  *models.User `json:"user" validate:"required"`
	Token string       `json:"token" validate:"required"`
}

type InitResponse struct {
	Setup                bool   `json:"setup" validate:"required"`
	GoogleCaptchaSiteKey string `json:"google_captcha_secret_key" validate:"omitempty"`
}

type Login struct {
	Username   string `json:"username" validate:"required,min=2,max=16" example:"john_doe" `
	Password   string `json:"password" validate:"required,min=2,max=16" example:"doe123456"`
	RememberMe bool   `json:"remember_me" desc:"remember for a month"`
	Token      string `json:"token" desc:"captcha v2 token"`
}

type ConfigResponse struct {
	GoogleCaptchaSiteKey   string `json:"google_captcha_site_key" validate:"omitempty"`
	GoogleCaptchaSecretKey string `json:"google_captcha_secret_key" validate:"omitempty"`
}

type ConfigData struct {
	GoogleCaptchaSiteKey   *string `json:"google_captcha_site_key" validate:"omitempty"`
	GoogleCaptchaSecretKey *string `json:"google_captcha_secret_key" validate:"omitempty"`
}
