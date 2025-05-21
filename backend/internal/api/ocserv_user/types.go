package ocservUser

import (
	"ocserv/pkg/oc"
	"ocserv/pkg/request"
	"time"
)

type OcservUsersResponse struct {
	Meta   request.Meta     `json:"meta" validate:"required"`
	Result *[]oc.OcservUser `json:"result" validate:"omitempty"`
}

const (
	Free            = "Free"
	MonthlyTransmit = "MonthlyTransmit"
	MonthlyReceive  = "MonthlyReceive"
	TotallyTransmit = "TotallyTransmit"
	TotallyReceive  = "TotallyReceive"
)

type createOcservUserData struct {
	Group       string     `json:"group" validate:"required" example:"default"`
	Username    string     `json:"username" validate:"required,min=3,max=32" example:"john_doe"`
	Password    string     `json:"password" validate:"required,min=6" example:"strongpassword123"`
	ExpireAt    *time.Time `json:"expire_at" validate:"required" example:"2025-12-31T23:59:59Z"`
	TrafficType string     `json:"traffic_type" validate:"required,oneof=Free MonthlyTransmit MonthlyReceive TotallyTransmit TotallyReceive" example:"MonthlyTransmit"`
	TrafficSize int        `json:"traffic_size" validate:"required,gt=0" example:"10737418240"` // 10 GiB
	Description string     `json:"description,omitempty" validate:"omitempty,max=1024" example:"User for testing VPN access"`
}
