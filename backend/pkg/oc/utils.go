package oc

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

// toMap convert data struct to json
func toMap(data interface{}) map[string]interface{} {
	b, _ := json.Marshal(&data)
	var dataStruct map[string]interface{}
	_ = json.Unmarshal(b, &dataStruct)
	return dataStruct
}

// runWithContext run methods commands with context
func runWithContext(c context.Context, operation func() error) error {
	done := make(chan error, 1)

	go func() {
		defer close(done)
		done <- operation()
	}()

	select {
	case <-c.Done():
		return fmt.Errorf("operation canceled or timed out: %w", c.Err())
	case err := <-done:
		return err
	}
}

// runWithContextResult runs an operation with context and returns a result and error
func runWithContextResult[Type any](c context.Context, operation func() (Type, error)) (Type, error) {
	var zero Type // default zero value for T in case of timeout/cancel
	resultChan := make(chan struct {
		result Type
		err    error
	}, 1)

	go func() {
		defer close(resultChan)
		res, err := operation()
		resultChan <- struct {
			result Type
			err    error
		}{res, err}
	}()

	select {
	case <-c.Done():
		return zero, fmt.Errorf("operation canceled or timed out: %w", c.Err())
	case out := <-resultChan:
		return out.result, out.err
	}
}

// Writer a method to write configs in group config file
func groupWriter(file *os.File, group *OcservDefaultConfigs) error {
	for k, v := range toMap(group) {

		if b, ok := v.(bool); ok && !b {
			continue
		}
		if v == nil {
			continue
		}

		if k == "dns" {
			for _, dns := range v.([]interface{}) {
				if _, err := file.WriteString(fmt.Sprintf("dns=%s\n", dns)); err != nil {
					return fmt.Errorf("failed to write to file: %w", err)
				}
			}
			continue
		} else if k == "route" {
			for _, route := range v.([]interface{}) {
				if _, err := file.WriteString(fmt.Sprintf("route=%s\n", route)); err != nil {
					return fmt.Errorf("failed to write to file: %w", err)
				}
			}
			continue
		} else if k == "no-route" {
			for _, route := range v.([]interface{}) {
				if _, err := file.WriteString(fmt.Sprintf("no-route=%s\n", route)); err != nil {
					return fmt.Errorf("failed to write to file: %w", err)
				}
			}
			continue
		} else if k == "split-dns" {
			for _, dns := range v.([]interface{}) {
				if _, err := file.WriteString(fmt.Sprintf("split-dns=%s\n", dns)); err != nil {
					return fmt.Errorf("failed to write to file: %w", err)
				}
			}
			continue
		} else {
			if _, err := file.WriteString(fmt.Sprintf("%s=%v\n", k, v)); err != nil {
				return fmt.Errorf("failed to write to file: %w", err)
			}
		}
	}
	return nil
}

// parseConfigFile parse ocserv group config file in OcGroupConfig data type
func parseConfigFile(filename string) (*OcservDefaultConfigs, error) {
	var config OcservDefaultConfigs
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	var dnsList, routeList, noRouteList, splitDNSList []string

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "rx-data-per-sec":
			vInt, _ := strconv.Atoi(value)
			config.RxDataPerSec = &vInt
		case "tx-data-per-sec":
			vInt, _ := strconv.Atoi(value)
			config.TxDataPerSec = &vInt
		case "max-same-clients":
			if val, err := strconv.Atoi(value); err == nil {
				config.MaxSameClients = &val
			}
		case "ipv4-network":
			config.IPv4Network = &value
		case "dns":
			dnsList = append(dnsList, value)
		case "no-udp":
			if val, err := strconv.ParseBool(value); err == nil {
				config.NoUDP = &val
			}
		case "keepalive":
			if val, err := strconv.Atoi(value); err == nil {
				config.KeepAlive = &val
			}
		case "dpd":
			if val, err := strconv.Atoi(value); err == nil {
				config.DPD = &val
			}
		case "mobile-dpd":
			if val, err := strconv.Atoi(value); err == nil {
				config.MobileDPD = &val
			}
		case "tunnel-all-dns":
			if val, err := strconv.ParseBool(value); err == nil {
				config.TunnelAllDNS = &val
			}
		case "restrict-user-to-routes":
			if val, err := strconv.ParseBool(value); err == nil {
				config.RestrictUserToRoutes = &val
			}
		case "stats-report-time":
			if val, err := strconv.Atoi(value); err == nil {
				config.StatsReportTime = &val
			}
		case "mtu":
			if val, err := strconv.Atoi(value); err == nil {
				config.MTU = &val
			}
		case "idle-timeout":
			if val, err := strconv.Atoi(value); err == nil {
				config.IdleTimeout = &val
			}
		case "mobile-idle-timeout":
			if val, err := strconv.Atoi(value); err == nil {
				config.MobileIdleTimeout = &val
			}
		case "session-timeout":
			if val, err := strconv.Atoi(value); err == nil {
				config.SessionTimeout = &val
			}
		case "route":
			routeList = append(routeList, value)
		case "no-route":
			noRouteList = append(noRouteList, value)
		case "split-dns":
			splitDNSList = append(splitDNSList, value)
		case "nbns":
			config.NBNS = &value
		case "explicit-ipv4":
			config.ExplicitIPv4 = &value
		case "cgroup":
			config.CGroup = &value
		case "iroute":
			config.IRoute = &value
		case "net-priority":
			if val, err := strconv.Atoi(value); err == nil {
				config.NetPriority = &val
			}
		case "deny-roaming":
			if val, err := strconv.ParseBool(value); err == nil {
				config.DenyRoaming = &val
			}
		case "restrict-user-to-ports":
			config.RestrictUserToPorts = &value
		}
	}

	config.DNS = &dnsList
	config.Route = &routeList
	config.NoRoute = &noRouteList
	config.SplitDNS = &splitDNSList

	if err = scanner.Err(); err != nil {
		return &config, err
	}
	return &config, nil
}

// occtlExec CMD execute command runner for ocserv service
func occtlExec(c context.Context, command string) ([]byte, error) {
	// TODO: check sudo in systemd
	cmd := exec.CommandContext(c, "sh", "-c", fmt.Sprintf("/usr/bin/occtl %s", command))
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return output, nil
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

// getGroupList handler to get group name and path
func getGroupList() ([]GroupInfo, error) {
	var groupsInfo []GroupInfo

	err := filepath.Walk(groupDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			groupsInfo = append(groupsInfo, GroupInfo{
				Name: info.Name(),
				Path: path,
			})
		}
		return nil
	})
	return groupsInfo, err
}

// getGroupConfig handler to get group config with path entry
func getGroupConfig(path string) (*OcservDefaultConfigs, error) {
	groupConf, err := parseConfigFile(path)
	if err != nil {
		fmt.Printf("Error parsing file %s: %v\n", path, err)
		return nil, err
	}
	return groupConf, nil
}
