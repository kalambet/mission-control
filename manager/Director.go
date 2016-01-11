package manager

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/kalambet/mission-control/config"
	"github.com/kalambet/mission-control/services"
)

// Director is the main orchestator of the status cheking process
type Director struct {
	_token string
}

// Direct is the main orchestration method
func (d *Director) Direct() {
	servicesList, err := d.Init()
	if err != nil {
		return
	}

	fmt.Println(servicesList)

	//Goroutine for statistic gethering
	return
}

// Init intialize all the environment and prepares monitor for status gathering
func (d *Director) Init() (serviceList []*services.Service, err error) {
	configuration, err := d.getConfig()
	if err != nil || configuration == nil {
		return
	}

	serviceList, err = d.getServices(configuration)
	if err != nil {
		return
	}

	return
}

func (d *Director) getConfig() (configuration *config.EnvConfig, err error) {
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

func (d *Director) getServices(configuration *config.EnvConfig) (serviceList []*services.Service, err error) {
	browser := ServiceBrowser{
		serviceConfig: configuration,
		token:         "eyJhbGciOiJIUzI1NiJ9.eyJqdGkiOiJhMTNjM2IxZi02NmQxLTQyNTYtYmI3Zi0yM2M3OWU4ODQyNmIiLCJzdWIiOiJmNzM1NDVhYy0xYTVkLTRjZGMtOTVkYy04MWE1N2VkODM3NjQiLCJzY29wZSI6WyJjbG91ZF9jb250cm9sbGVyLnJlYWQiLCJwYXNzd29yZC53cml0ZSIsImNsb3VkX2NvbnRyb2xsZXIud3JpdGUiLCJvcGVuaWQiXSwiY2xpZW50X2lkIjoiY2YiLCJjaWQiOiJjZiIsImF6cCI6ImNmIiwiZ3JhbnRfdHlwZSI6InBhc3N3b3JkIiwidXNlcl9pZCI6ImY3MzU0NWFjLTFhNWQtNGNkYy05NWRjLTgxYTU3ZWQ4Mzc2NCIsIm9yaWdpbiI6InVhYSIsInVzZXJfbmFtZSI6InBldGVyLmthbGFtYmV0QHJ1LmlibS5jb20iLCJlbWFpbCI6InBldGVyLmthbGFtYmV0QHJ1LmlibS5jb20iLCJhdXRoX3RpbWUiOjE0NTI1MDY2ODUsInJldl9zaWciOiJjZjkzZDhkZCIsImlhdCI6MTQ1MjUwNjY4NSwiZXhwIjoxNDUyNTQ5ODg1LCJpc3MiOiJodHRwczovL3VhYS5uZy5ibHVlbWl4Lm5ldC9vYXV0aC90b2tlbiIsInppZCI6InVhYSIsImF1ZCI6WyJjZiIsImNsb3VkX2NvbnRyb2xsZXIiLCJwYXNzd29yZCIsIm9wZW5pZCJdfQ.1BBitoYuzgaP7eMxRnHHOqCQAS9EfhJhd8K3L2r88dw"}

	serviceList, err = browser.GetServices()
	if err != nil {
		return nil, err
	}

	return nil, errors.New("Not yet implemented")
}

//ReadServiceList returns list of services providede to the MC
func ReadServiceList() (serviceList []*services.Service, err error) {
	return nil, errors.New("Not yet implemented")
}

func getServicesStatus() (list []*services.ServiceStatus) {
	return
}
