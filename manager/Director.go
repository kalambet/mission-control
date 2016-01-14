package manager

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/kalambet/mission-control/config"
	"github.com/kalambet/mission-control/services"
)

// Director orchistrates the whole thing
type Director struct {
	browser      *ServiceBrowser
	servicesList []*services.Service
}

// Init initialize Director with the users settings
// and searches for the application ids
func (d *Director) Init() (err error) {
	// Reads monitor configuration from the
	configuration, err := getConfig()
	if err != nil || configuration == nil {
		return err
	}

	d.browser = &ServiceBrowser{
		serviceConfig: configuration}
	//token:         "eyJhbGciOiJIUzI1NiJ9.eyJqdGkiOiJkZGQ4NGE4YS05YjNmLTQxMTQtOWViOC0yM2MyYTg2ODUzMGQiLCJzdWIiOiJmNzM1NDVhYy0xYTVkLTRjZGMtOTVkYy04MWE1N2VkODM3NjQiLCJzY29wZSI6WyJjbG91ZF9jb250cm9sbGVyLnJlYWQiLCJwYXNzd29yZC53cml0ZSIsImNsb3VkX2NvbnRyb2xsZXIud3JpdGUiLCJvcGVuaWQiXSwiY2xpZW50X2lkIjoiY2YiLCJjaWQiOiJjZiIsImF6cCI6ImNmIiwiZ3JhbnRfdHlwZSI6InBhc3N3b3JkIiwidXNlcl9pZCI6ImY3MzU0NWFjLTFhNWQtNGNkYy05NWRjLTgxYTU3ZWQ4Mzc2NCIsIm9yaWdpbiI6InVhYSIsInVzZXJfbmFtZSI6InBldGVyLmthbGFtYmV0QHJ1LmlibS5jb20iLCJlbWFpbCI6InBldGVyLmthbGFtYmV0QHJ1LmlibS5jb20iLCJhdXRoX3RpbWUiOjE0NTI2Njc5MjEsInJldl9zaWciOiJjZjkzZDhkZCIsImlhdCI6MTQ1MjY2NzkyMSwiZXhwIjoxNDUyNzExMTIxLCJpc3MiOiJodHRwczovL3VhYS5uZy5ibHVlbWl4Lm5ldC9vYXV0aC90b2tlbiIsInppZCI6InVhYSIsImF1ZCI6WyJjZiIsImNsb3VkX2NvbnRyb2xsZXIiLCJwYXNzd29yZCIsIm9wZW5pZCJdfQ.xyCKDoEJY5G4LiYQop9bbq58boYnf5PQt9JqyS_1Fow"}

	d.servicesList, err = d.browser.GetServices()
	if err != nil {
		return
	}

	return
}

// GetServicesStatus returns services staus by request
func (d *Director) GetServicesStatus() (statusList []*services.ServiceStatus, err error) {
	statusList = make([]*services.ServiceStatus, len(d.servicesList))

	for _, service := range d.servicesList {
		status, err := d.browser.CollectAndSaveServiceState(service)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		statusList = append(statusList, status)
	}

	return statusList, nil
}

// getConfig reads the config from the Environment Variables
func getConfig() (configuration *config.EnvConfig, err error) {
	// Let's read config from Environment Variable
	var configString = os.Getenv("BMC_CONFIG")

	fmt.Println(configString)

	if configString == "" {
		err = errors.New("No config detected, please provide config in $BMC_CONFIG environment varialbe")
		fmt.Println(err)
		return nil, err
	}

	var cfg config.EnvConfig
	err = json.Unmarshal([]byte(configString), &cfg)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &cfg, nil
}
