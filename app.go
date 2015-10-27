package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	//for extracting service credentials from VCAP_SERVICES
	//"github.com/cloudfoundry-community/go-cfenv"
)

const (
	DEFAULT_PORT = "8080"
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
func CreateServiceList() (list []Service) {
	list = append(list, Service{Name: "Authorization", HostName: "https://idmx-service-authz.mybluemix.net"})
	list = append(list, Service{Name: "Authentication", HostName: "https://idmx-service-authn.mybluemix.net"})
	list = append(list, Service{Name: "Directory", HostName: "https://idmx-service-directory.mybluemix.net"})
	list = append(list, Service{Name: "Issuance", HostName: "https://idmx-service-issuing.mybluemix.net"})
	list = append(list, Service{Name: "Media", HostName: "https://idmx-service-media.mybluemix.net/docs/idemix.json"})
	list = append(list, Service{Name: "ID Service", HostName: "https://idmx-service-wallet.mybluemix.net"})
	list = append(list, Service{Name: "Service Broker", HostName: "https://idmx-service.mybluemix.net"})
	return
}

var index = template.Must(template.ParseFiles(
	"templates/_base.html",
	"templates/index.html",
))

var table = template.Must(template.ParseFiles(
	"templates/table.html",
))

func helloworld(w http.ResponseWriter, req *http.Request) {
	index.Execute(w, nil)
}

func handleStatusRequest(w http.ResponseWriter, r *http.Request) {
	error := table.Execute(w, CreateServiceList())
	if error != nil {
		fmt.Printf("\nError during template formation: %s\n", error)
	}
}

func main() {
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = DEFAULT_PORT
	}

	http.HandleFunc("/", helloworld)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/status", handleStatusRequest)

	log.Printf("Starting app on port %+v\n", port)
	http.ListenAndServe(":"+port, nil)
}
