package manager

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/kalambet/mission-control/config"
	"github.com/kalambet/mission-control/db"
	"github.com/kalambet/mission-control/services"
)

// Director orchistrates the whole thing
type Director struct {
	browser      *ServiceBrowser
	dbDriver     *db.Driver
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

	d.dbDriver = &db.Driver{
		URI: configuration.DBURI}

	d.dbDriver.InitDatabase()

	d.browser = &ServiceBrowser{
		serviceConfig: configuration,
		token:         ""}

	d.servicesList, err = d.browser.GetServices()
	if err != nil {
		return
	}

	return
}

// ScheduleStatusCollection returns services staus by request
func (d *Director) ScheduleStatusCollection() {
	ticker := time.NewTicker(20 * time.Second)
	quit := make(chan struct{})
	//go func() {
	for {
		select {
		case <-ticker.C:
			for _, service := range d.servicesList {
				d.collectAndSaveServiceState(service)
			}
		case <-quit:
			ticker.Stop()
			return
		}
	}
	//}()

	//	return
}

func (d *Director) collectAndSaveServiceState(service *services.Service) {
	status, err := d.browser.CollectServiceStatus(service)
	if err != nil {
		fmt.Printf("Problem COLLECTING status for servcie %s with the following error: %s\n", service.Name, err)
		return
	}

	err = d.dbDriver.SaveStatus(status)
	if err != nil {
		fmt.Printf("Problem SAVING status for servcie %s with the following error: %s\n", service.Name, err)
		return
	}
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
