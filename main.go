package main

import (
	"html/template"
	"net/http"

	"github.com/kalambet/mission-control/manager"

	//for extracting service credentials from VCAP_SERVICES
	//"github.com/cloudfoundry-community/go-cfenv"
)

const (
	DEFAULT_PORT = "8080"
)

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

	//error := table.Execute(w, CreateServiceList())
	//if error != nil {
	//	fmt.Printf("\nError during template formation: %s\n", error)
	//}
}

func main() {
	manager.Direct()

	/*
		var port string
		if port = os.Getenv("PORT"); len(port) == 0 {
			port = DEFAULT_PORT
		}
	*/
	//http.HandleFunc("/", helloworld)
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	//http.HandleFunc("/status", handleStatusRequest)

	//log.Printf("Starting app on port %+v\n", port)
	//http.ListenAndServe(":"+port, nil)
}
