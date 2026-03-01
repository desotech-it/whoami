package app

import (
	"flag"
	"fmt"
	"os"
)

var (
	version string
	commit string
	date string
)

type Config struct {
	Port    uint64
	TLSPort uint64
}

var port = flag.Uint64("port", 8080, "")
var tlsPort = flag.Uint64("tls-port", 0, "")
var versionFlag = flag.Bool("version", false, "Print app version")

func init() {
	// short versions
	flag.Uint64Var(port, "p", 8080, "")
	flag.Uint64Var(tlsPort, "tp", 0, "")
	flag.Parse()

	if *versionFlag {
		fmt.Printf("whoami by Desotech %s\nGit Commit: %s\nBuilt On: %s\n", version, commit, date)
		os.Exit(0)
	}
}

func GetConfig() Config {
	return Config{
		Port:    *port,
		TLSPort: *tlsPort,
	}
}
