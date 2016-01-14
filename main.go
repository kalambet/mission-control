package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/kalambet/mission-control/manager"
)

const (
	defaultPort = "8080"
)

var indexPage = template.Must(template.ParseFiles(
	"templates/_base.html",
	"templates/index.html",
))

var errorPage = template.Must(template.ParseFiles(
	"templates/_base.html",
	"templates/error.html",
))

var tablePage = template.Must(template.ParseFiles(
	"templates/table.html",
))

var director = manager.Director{}

func root(w http.ResponseWriter, req *http.Request) {
	indexPage.Execute(w, nil)
}

func handleStatusRequest(w http.ResponseWriter, r *http.Request) {
	// For now we'll do it per request
	// statistic collection capabilities will be added later
	servicesStatus, err := director.GetServicesStatus()
	if err != nil {
		errorPage.Execute(w, nil)
		return
	}
	err = tablePage.Execute(w, servicesStatus)
	if err != nil {
		fmt.Printf("\nError during template formation: %s\n", err)
	}
}

func main() {
	director.Init()

	http.HandleFunc("/", root)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/status", handleStatusRequest)

	log.Printf("Starting app on port %+v\n", defaultPort)
	http.ListenAndServe(":"+defaultPort, nil)
}
