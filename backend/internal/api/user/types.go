package user

import (
	"ocserv/internal/models"
	"ocserv/pkg/request"
)

type ChangePasswordData struct {
	CurrentPassword string `json:"current_password" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required"`
}

type StaffsResponse struct {
	Meta   request.Meta   `json:"meta" validate:"required"`
	Result *[]models.User `json:"result" validate:"omitempty"`
}

type ChangeStaffPassword struct {
	Password string `json:"password" validate:"required"`
}

type CreateStaffData struct {
	Username   string            `json:"username" validate:"required"`
	Password   string            `json:"password" validate:"required"`
	Permission models.Permission `json:"permission" validate:"required"`
}
