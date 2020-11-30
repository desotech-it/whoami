package main

import (
	"desotech/whoami/app"
	"desotech/whoami/server"
)

func main() {
	config := app.GetConfig()
	server := server.Server{
		config.Port,
	}
	server.Start()
}
