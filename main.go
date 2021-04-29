package main

import (
	"github.com/desotech-it/whoami/app"
	"github.com/desotech-it/whoami/server"
)

func main() {
	config := app.GetConfig()
	server := server.Server{
		Port: config.Port,
	}
	server.Start()
}
