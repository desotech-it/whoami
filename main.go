package main

import (
	"desotech/whoami/app"
	"desotech/whoami/server"
)

func main() {
	server.Start(app.GetConfig())
}
