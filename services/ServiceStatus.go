package services

import "github.com/kalambet/mission-control/types"

// ServiceStatus holds status of the service
type ServiceStatus struct {
	Name         string
	Organization string
	Space        string
	UpdateTime   string
	Instances    []types.InstanceState
}
