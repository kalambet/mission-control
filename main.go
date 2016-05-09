package main

import (
	"github.com/kalambet/mission-control/manager"
	"github.com/kalambet/mission-control/net"
)

var director = manager.Director{}

func main() {
	director.Init()
	var restServer = net.RestHandler{Director: director}
	director.ScheduleStatusCollection()
	// Still need to check but for now it sould be the last command
	restServer.StartRestServer()
}
