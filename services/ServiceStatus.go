package services

import (
	"time"

	"github.com/kalambet/mission-control/types"
)

// ServiceStatus holds status of the service
type ServiceStatus struct {
	Name         string
	Organization string
	Space        string
	Instances    []types.InstanceState
	UpdateTime   time.Time
}
