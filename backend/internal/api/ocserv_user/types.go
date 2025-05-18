package ocservUser

import (
	"ocserv/pkg/oc"
	"ocserv/pkg/request"
)

type OcservUsersResponse struct {
	Meta   request.Meta     `json:"meta" validate:"required"`
	Result *[]oc.OcservUser `json:"result" validate:"omitempty"`
}
