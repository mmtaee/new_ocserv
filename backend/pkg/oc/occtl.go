package oc

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type OcctlService struct {
	ctx context.Context
}

type OcctlServiceInterface interface {
	WithContext(ctx context.Context) *OcctlService
	GetOnlineUsers() ([]string, error)
	DisconnectUser(username string) (string, error)
}

func NewOcctlService() *OcctlService {
	return &OcctlService{}
}

func doExecWithContext(c context.Context, command string) ([]byte, error) {
	cmd := exec.CommandContext(c, "sh", "-c", fmt.Sprintf("/usr/bin/occtl %s", command))
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return output, nil
}

func doExec(command string) ([]byte, error) {
	cmd := exec.Command("sh", "-c", fmt.Sprintf("/usr/bin/occtl %s", command))
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (o *OcctlService) WithContext(ctx context.Context) *OcctlService {
	return &OcctlService{ctx: ctx}
}

// GetOnlineUsers list of online users
func (o *OcctlService) GetOnlineUsers() ([]string, error) {
	var (
		//ocUsers []OcservUser
		users  []string
		result []byte
		err    error
	)

	command := "-j show users | jq -r '.[].Username'"

	if o.ctx == nil {
		result, err = doExec(command)
	} else {
		result, err = doExecWithContext(o.ctx, command)
	}

	if err != nil {
		return nil, err
	}

	for _, line := range strings.Split(string(result), "\n") {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			users = append(users, trimmed)
		}
	}

	return users, nil
}

// DisconnectUser disconnect user
func (o *OcctlService) DisconnectUser(username string) (string, error) {
	var (
		result []byte
		err    error
	)

	command := fmt.Sprintf("disconnect user %s", username)
	log.Println("disconnect user command:", command)
	if o.ctx == nil {
		result, err = doExec(command)
	} else {
		result, err = doExecWithContext(o.ctx, command)
	}
	return string(result), err
}
