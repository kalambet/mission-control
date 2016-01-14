package services

import (
	"time"

	"github.com/kalambet/mission-control/data"
)

// ServiceStatus holds status of the service
type ServiceStatus struct {
	Name         string
	Organization string
	Space        string
	Instances    []data.InstanceState
	UpdateTime   time.Time
}
