package panel

import (
	"ocserv/internal/models"
	"ocserv/pkg/oc"
)

type RequestSetup struct {
	Config struct {
		AdminUsername          string `json:"admin_username" validate:"required,min=2,max=16" example:"john_doe" `
		AdminPassword          string `json:"admin_password" validate:"required,min=2,max=16" example:"doe123456"`
		GoogleCaptchaSecretKey string `json:"google_captcha_secret_key" validate:"omitempty"`
		GoogleCaptchaSiteKey   string `json:"google_captcha_site_key" validate:"omitempty"`
	} `json:"config" validate:"required"`
	DefaultOcservGroup *oc.OcservDefaultConfigs `json:"default_ocserv_group" validate:"omitempty"`
}

type ResponseSetup struct {
	User  *models.User `json:"user" validate:"required"`
	Token string       `json:"token" validate:"required"`
}

type ConfigResponse struct {
	Setup                bool   `json:"setup" validate:"required"`
	GoogleCaptchaSiteKey string `json:"google_captcha_secret_key" validate:"omitempty"`
}

type Login struct {
	Username   string `json:"username" validate:"required,min=2,max=16" example:"john_doe" `
	Password   string `json:"password" validate:"required,min=2,max=16" example:"doe123456"`
	RememberMe bool   `json:"remember_me" desc:"remember for a month"`
	Token      string `json:"token" desc:"captcha v2 token"`
}

type LoginResponse struct {
	User  *models.User `json:"user" validate:"required"`
	Token string       `json:"token" validate:"required"`
}
