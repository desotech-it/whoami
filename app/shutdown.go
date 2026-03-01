package app

import (
	"os"
	"time"
)

var (
	shutdownDelay time.Duration
	startupDelay  time.Duration
)

func init() {
	if v, ok := os.LookupEnv("SHUTDOWN_DELAY"); ok {
		if d, err := ParseDuration(v); err == nil && d >= 0 {
			shutdownDelay = d
		}
	}
	if v, ok := os.LookupEnv("STARTUP_DELAY"); ok {
		if d, err := ParseDuration(v); err == nil && d >= 0 {
			startupDelay = d
		}
	}
}

// GetShutdownDelay returns the configured shutdown delay.
func GetShutdownDelay() time.Duration {
	return shutdownDelay
}

// GetStartupDelay returns the configured startup delay.
func GetStartupDelay() time.Duration {
	return startupDelay
}
