package ocservUser

import (
	"ocserv/pkg/oc"
	"ocserv/pkg/request"
)

type OcservUsersResponse struct {
	Meta   request.Meta     `json:"meta" validate:"required"`
	Result *[]oc.OcservUser `json:"result" validate:"omitempty"`
}

type createOcservUserData struct {
	Group       string `json:"group" validate:"required" example:"default"`
	Username    string `json:"username" validate:"required,min=3,max=32" example:"john_doe"`
	Password    string `json:"password" validate:"required,min=6" example:"strongpassword123"`
	ExpireAt    string `json:"expire_at" validate:"omitempty" example:"2025-12-31"`
	TrafficType string `json:"traffic_type" validate:"required,oneof=Free MonthlyTransmit MonthlyReceive TotallyTransmit TotallyReceive" example:"MonthlyTransmit"`
	TrafficSize int    `json:"traffic_size" validate:"required,gt=0" example:"10737418240"` // 10 GiB
	Description string `json:"description,omitempty" validate:"omitempty,max=1024" example:"User for testing VPN access"`
}

type LockOcservUserData struct {
	Lock *bool `json:"lock" validate:"required" example:"false"`
}

type UpdateOcservUserData struct {
	Group       *string `json:"group" example:"default"`
	Password    *string `json:"password" validate:"min=6" example:"strongpassword123"`
	TrafficType *string `json:"traffic_type" validate:"oneof=Free MonthlyTransmit MonthlyReceive TotallyTransmit TotallyReceive" example:"MonthlyTransmit"`
	TrafficSize *int    `json:"traffic_size" validate:"gt=0" example:"10737418240"` // 10 GiB
	Description *string `json:"description,omitempty" validate:"omitempty,max=1024" example:"User for testing VPN access"`
}
