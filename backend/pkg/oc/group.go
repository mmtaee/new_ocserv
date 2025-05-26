package oc

import (
	"context"
	"errors"
	"fmt"
	"ocserv/pkg/config"
	"os"
	"sort"
	"sync"
)

type OcservGroupService struct {
	ctx context.Context
}

type GroupInfo struct {
	Name   string
	Path   string
	Config OcservDefaultConfigs
}

type OcservGroupServiceInterface interface {
	WithContext(ctx context.Context) *OcservGroupService
	UpdateDefaultGroup(group *OcservDefaultConfigs) error
	GetGroup(name string) (*OcservDefaultConfigs, error)
	GetGroups() (*[]GroupInfo, error)
	GetGroupNames() (*[]string, error)
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

// GetGroups ocserv group with full info(name, path, configs)
func (g *OcservGroupService) GetGroups() (*[]GroupInfo, error) {
	var (
		result []GroupInfo
		err    error
		wg     sync.WaitGroup
	)

	op := func() ([]GroupInfo, error) {
		return getGroupList()
	}

	if g.ctx == nil {
		result, err = op()
	} else {
		result, err = runWithContextResult(g.ctx, func() ([]GroupInfo, error) { return op() })
	}
	if err != nil {
		return nil, err
	}

	for _, v := range result {
		wg.Add(1)

		go func(v GroupInfo) {
			defer wg.Done()
			conf, err := getGroupConfig(v.Path)
			if err != nil {
				return
			}
			v.Config = *conf
		}(v)
	}

	wg.Wait()

	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})
	return &result, nil
}

// GetGroupNames list of groups name that in service exists
func (g *OcservGroupService) GetGroupNames() (*[]string, error) {
	var (
		groups []string
		result []GroupInfo
		err    error
	)

	op := func() ([]GroupInfo, error) {
		return getGroupList()
	}

	if g.ctx == nil {
		result, err = op()
		if err != nil {
			return nil, err
		}
	} else {
		result, err = runWithContextResult(g.ctx, func() ([]GroupInfo, error) { return op() })
		if err != nil {
			return nil, err
		}
	}

	for _, v := range result {
		groups = append(groups, v.Name)
	}
	return &groups, nil
}
