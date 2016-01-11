// Service we are going to monitor

package services

// Service represents one Bluemix service which
// status needs to be collected
type Service struct {
	URL  string
	Name string
	GUID string
}

// GetName returns Service friendly Name
func (s *Service) GetName() string {
	return s.Name
}

// GetStatus returns Service current status
func (s *Service) GetStatus() (status *ServiceStatus) {
	return nil
}
