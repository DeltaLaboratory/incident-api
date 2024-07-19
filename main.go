package main

import "github.com/DeltaLaboratory/incident-api/server"

// @title			Incident API
// @version		0.1.0
// @description	This is the Incident API server.
// @schemes		https
// @host			incident.deltalab.dev
// @BasePath		/v0
func main() {
	s := &server.Server{}
	s.Run()
}
