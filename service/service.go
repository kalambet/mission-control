package service

// Service is type describes serves
import (
	"fmt"
	"net/http"

	"github.com/kalambet/mission-control/service"
)

// Service we are going to monitor
type Service struct {
	HostName string
	Name     string
}

// GetName returns Service friendly Name
func (s Service) GetName() string {
	return s.Name
}

// GetStatus returns Service current status
func (s Service) GetStatus() bool {
	resp, err := http.Get(s.HostName)
	if err != nil {
		fmt.Printf("\nError getting %s service: %s\n", s.HostName, err)
		return false
	}

	if code := resp.StatusCode; code != 401 && code != 200 {
		fmt.Printf("\nError getting %s service, response status is %d\n", s.HostName, resp.StatusCode)
		return false
	}

	return true
}

// CreateServiceList creates default service list
func CreateServiceList() (list []service.Service) {
	list = append(list, service.Service{Name: "Authorization", HostName: "https://idmx-service-authz.mybluemix.net"})
	list = append(list, service.Service{Name: "Authentication", HostName: "https://idmx-service-authn.mybluemix.net"})
	list = append(list, service.Service{Name: "Directory", HostName: "https://idmx-service-directory.mybluemix.net"})
	list = append(list, service.Service{Name: "Issuance", HostName: "https://idmx-service-issuing.mybluemix.net"})
	list = append(list, service.Service{Name: "Media", HostName: "https://idmx-service-media.mybluemix.net/docs/idemix.json"})
	list = append(list, service.Service{Name: "ID Service", HostName: "https://idmx-service-wallet.mybluemix.net"})
	list = append(list, service.Service{Name: "Service Broker", HostName: "https://idmx-service.mybluemix.net"})
	return
}
