package setup

import "ocserv/internal/models"

type RequestSetup struct {
	Config struct {
		AdminUsername          string `json:"admin_username" validate:"required,min=2,max=16" example:"john_doe" `
		AdminPassword          string `json:"admin_password" validate:"required,min=2,max=16" example:"doe123456"`
		GoogleCaptchaSecretKey string `json:"google_captcha_secret_key" validate:"omitempty"`
		GoogleCaptchaSiteKey   string `json:"google_captcha_site_key" validate:"omitempty"`
	} `json:"config" validate:"required"`

	DefaultOcservGroup models.OcservUserOrGroupConfigs `json:"default_ocserv_group" validate:"required"`
}

type ResponseSetup struct {
	User  models.User `json:"user" validate:"required"`
	Token string      `json:"token" validate:"required"`
}
