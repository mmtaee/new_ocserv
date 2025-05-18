package oc

import (
	"context"
	"errors"
	"fmt"
	"log"
	"ocserv/pkg/config"
	"os"
)

type OcservGroupService struct {
	ctx context.Context
}

type OcservGroupServiceInterface interface {
	WithContext(ctx context.Context) *OcservGroupService
	UpdateDefaultGroup(group *OcservDefaultConfigs) error
	GetGroup(name string) (*OcservDefaultConfigs, error)
	CreateOrUpdateGroup(name string, group *OcservDefaultConfigs) error
	Delete(name string) error
}

var groupDir = "/etc/ocserv/groups"

func NewOcservGroupService() *OcservGroupService {
	return &OcservGroupService{}
}

func (g *OcservGroupService) WithContext(ctx context.Context) *OcservGroupService {
	return &OcservGroupService{ctx: ctx}
}

// UpdateDefaultGroup update or create defaults/group.conf and update configs
func (g *OcservGroupService) UpdateDefaultGroup(group *OcservDefaultConfigs) error {
	groupPath := "/etc/ocserv/defaults/group.conf"
	// TODO: remove for prod
	if config.Get().Debug {
		groupPath = "/tmp/group.conf"
	}
	op := func() error {
		return createOrUpdate(group, groupPath)
	}
	if g.ctx == nil {
		return op()
	}
	return runWithContext(g.ctx, func() error { return op() })
}

// GetGroup retrieve config with group name
func (g *OcservGroupService) GetGroup(name string) (*OcservDefaultConfigs, error) {
	groupPath := fmt.Sprintf("%s/%s", groupDir, name)

	op := func(path string) (*OcservDefaultConfigs, error) {
		return parseConfigFile(path)
	}
	if g.ctx == nil {
		return op(groupPath)
	}

	return runWithContextResult(g.ctx, func() (*OcservDefaultConfigs, error) {
		return op(groupPath)
	})
}

// CreateOrUpdateGroup create or update groups with config
func (g *OcservGroupService) CreateOrUpdateGroup(name string, group *OcservDefaultConfigs) error {
	if name == "defaults" {
		return errors.New("default group cannot be created or updated with this method")
	}
	groupPath := fmt.Sprintf("%s/%s", groupDir, name)
	if config.Get().Debug {
		groupPath = "/tmp/group.conf"
	}
	op := func() error {
		return createOrUpdate(group, groupPath)
	}
	if g.ctx == nil {
		return op()
	}
	return runWithContext(g.ctx, func() error { return op() })
}

// Delete ocserv group delete
func (g *OcservGroupService) Delete(name string) error {
	if name == "defaults" {
		return errors.New("default group cannot be deleted")
	}
	op := func() error {
		err := os.Remove(fmt.Sprintf("%s/%s", groupDir, name))
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				return fmt.Errorf("ocgroup %s does not exist", name)
			}
			return fmt.Errorf("failed to delete ocgroup %s: %w", name, err)
		}
		return nil
	}
	if g.ctx == nil {
		return op()
	}
	return runWithContext(g.ctx, func() error { return op() })
}

// createOrUpdate handler for group files
func createOrUpdate(group *OcservDefaultConfigs, path string) error {
	// os.O_WRONLY: Open the file write-only.
	// os.O_CREATE: Create the file if it doesn't exist.
	// os.O_TRUNC: Truncate the file (i.e., make it empty) if it exists.
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			log.Println(fmt.Sprintf("failed to close file: %v", closeErr))
		}
	}()
	return groupWriter(file, group)
}
