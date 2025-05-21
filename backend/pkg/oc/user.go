package oc

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"os/exec"
)

type OcservUserService struct {
}

type OcservUserServiceInterface interface {
	CreateUser(ctx context.Context, username, password string, group ...string) error
}

func NewOcservUserService(db *gorm.DB) *OcservUserService {
	return &OcservUserService{}
}

var (
	ocpasswdCMD = "/usr/bin/ocpasswd"    // ocpasswd os command path
	passwdFile  = "/etc/ocserv/ocpasswd" // ocpasswd file path
)

// CreateUser  ocserv user creation with password and group
func (o *OcservUserService) CreateUser(ctx context.Context, username, password string, group ...string) error {
	userGroup := "defaults"
	if len(group) > 0 {
		userGroup = group[0]
	}
	userGroup = fmt.Sprintf("-g %s", userGroup)
	command := fmt.Sprintf("/usr/bin/echo -e \"%s\\n%s\\n\" | %s %s -c %s %s",
		password,
		password,
		ocpasswdCMD,
		group,
		passwdFile,
		username,
	)
	return exec.CommandContext(ctx, "sh", "-c", command).Run()
}
