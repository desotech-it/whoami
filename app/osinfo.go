package app

import (
	"fmt"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v3/host"
)

// OSInfo holds operating system and runtime information.
type OSInfo struct {
	OS        string
	Arch      string
	GoVersion string
	NumCPU    int
	Uptime    string
	Hostname  string
}

// GetOSInfo returns system and Go runtime information.
func GetOSInfo() OSInfo {
	uptime := "N/A"
	if u, err := host.Uptime(); err == nil {
		d := time.Duration(u) * time.Second
		hours := int(d.Hours())
		minutes := int(d.Minutes()) % 60
		uptime = fmt.Sprintf("%dd %dh %dm", hours/24, hours%24, minutes)
	}

	return OSInfo{
		OS:        runtime.GOOS,
		Arch:      runtime.GOARCH,
		GoVersion: runtime.Version(),
		NumCPU:    runtime.NumCPU(),
		Uptime:    uptime,
		Hostname:  Info.Hostname,
	}
}
