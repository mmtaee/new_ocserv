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
	//userGroup := "defaults"
	//if len(group) > 0 {
	//	userGroup = group[0]
	//}
	//userGroup = fmt.Sprintf("-g %s", userGroup)
	//command := fmt.Sprintf("/usr/bin/echo -e \"%s\\n%s\\n\" | %s %s -c %s %s",
	//	password,
	//	password,
	//	ocpasswdCMD,
	//	group,
	//	passwdFile,
	//	username,
	//)
	//return exec.CommandContext(ctx, "sh", "-c", command).Run()
	userGroup := "defaults"
	if len(group) > 0 {
		userGroup = group[0]
	}

	args := []string{
		fmt.Sprintf("-g %s", userGroup),
		"-c",
		passwdFile,
		username,
	}

	cmd := exec.CommandContext(ctx, ocpasswdCMD, args...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return fmt.Errorf("error creating stdin pipe: %w", err)
	}

	go func() {
		defer stdin.Close()
		fmt.Fprintf(stdin, "%s\n%s\n", password, password)
	}()

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("command failed: %w, output: %s", err, string(output))
	}

	return nil
}
