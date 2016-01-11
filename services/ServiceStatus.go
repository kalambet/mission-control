package services

const (
	// ServiceStatusTypeBasic is a status constant that represents the lightweight
	// service with not much to monitor. For example like node.js, Ruby, Python
	ServiceStatusTypeBasic = iota

	// ServiceStatusTypeExtented is a status constant that represents the Java
	// service with much more parmeters to monitor
	ServiceStatusTypeExtented = iota
)

// ServiceStatus holds status of the service
type ServiceStatus struct {
	_type int
}
