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
	Update(c context.Context, username, password, group string) error
	LockUser(ctx context.Context, username string, lock bool) error
	DeleteUser(ctx context.Context, username string) error
}

func NewOcservUserService(db *gorm.DB) *OcservUserService {
	return &OcservUserService{}
}

var (
	ocpasswdCMD = "/usr/bin/ocpasswd"    // ocpasswd os command path
	passwdFile  = "/etc/ocserv/ocpasswd" // ocpasswd file path
)

func (o *OcservUserService) CreateUser(ctx context.Context, username, password string, group ...string) error {
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

func (o *OcservUserService) Update(c context.Context, username, password, group string) error {
	return o.CreateUser(c, username, password, group)
}

func (o *OcservUserService) LockUser(ctx context.Context, username string, lock bool) error {
	action := "-l"
	if !lock {
		action = "-u"
	}

	args := []string{
		action,
		"-c",
		passwdFile,
		username,
	}

	return exec.CommandContext(ctx, ocpasswdCMD, args...).Run()
}

func (o *OcservUserService) DeleteUser(ctx context.Context, username string) error {
	args := []string{
		"-c",
		passwdFile,
		"-d",
		username,
	}
	return exec.CommandContext(ctx, ocpasswdCMD, args...).Run()
}
