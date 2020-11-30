package app

import (
	"flag"
	"os"
)

type Config struct {
	Port       uint64
	WhoamiName string
}

var port = flag.Uint64("port", 8080, "")
var whoamiName = flag.String("name", os.Getenv("WHOAMI_NAME"), "")

func init() {
	// short version
	flag.Uint64Var(port, "p", 8080, "")
	flag.StringVar(whoamiName, "n", os.Getenv("WHOAMI_NAME"), "")
}

func GetConfig() Config {
	flag.Parse()
	return Config{
		*port,
		*whoamiName,
	}
}
