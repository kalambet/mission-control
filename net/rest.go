package net

import (
	"errors"
	"fmt"
	"net/http"
	"text/template"

	"github.com/kalambet/mission-control/manager"
)

const (
	defaultPort = "8080"
)

/*
var indexPage = template.Must(template.ParseFiles(
	"templates/_base.html",
	"templates/index.html",
))
*/
var errorPage = template.Must(template.ParseFiles(
	"templates/_base.html",
	"templates/error.html",
))

var tablePage = template.Must(template.ParseFiles(
	"templates/table.html"))

// RestHandler handler for all REST endpoints
type RestHandler struct {
	Director *manager.Director
}

// StartRestServer fires up all the endpoiunts
func (handler *RestHandler) StartRestServer() error {
	if handler.Director == nil {
		return errors.New("The Director is not initialized")
	}

	http.HandleFunc("/", handler.handleRootRequest)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/status", handler.handleStatusRequest)
	http.ListenAndServe(":"+defaultPort, nil)

	return nil
}

func (handler *RestHandler) handleStatusRequest(w http.ResponseWriter, r *http.Request) {
	servicesStatus, err := handler.director.GetServicesStatus()
	if err != nil {
		errorPage.Execute(w, nil)
		return
	}
	err = tablePage.Execute(w, servicesStatus)
	if err != nil {
		fmt.Printf("\nError during template formation: %s\n", err)
	}

}

func (handler *RestHandler) handleRootRequest(w http.ResponseWriter, r *http.Request) {
	return
}
