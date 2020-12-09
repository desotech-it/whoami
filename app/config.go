package app

import (
	"flag"
)

type Config struct {
	Port       uint64
}

var port = flag.Uint64("port", 8080, "")

func init() {
	// short version
	flag.Uint64Var(port, "p", 8080, "")
}

func GetConfig() Config {
	flag.Parse()
	return Config{
		*port,
	}
}
