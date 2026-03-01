package app

import (
	"os"
	"time"
)

var (
	cpuStressDuration time.Duration
	memStressDuration time.Duration
)

func init() {
	if v, ok := os.LookupEnv("CPU_STRESS"); ok {
		if d, err := ParseDuration(v); err == nil && d > 0 {
			cpuStressDuration = d
		}
	}
	if v, ok := os.LookupEnv("MEM_STRESS"); ok {
		if d, err := ParseDuration(v); err == nil && d > 0 {
			memStressDuration = d
		}
	}
}

// GetCPUStressDuration returns the configured CPU stress duration.
func GetCPUStressDuration() time.Duration {
	return cpuStressDuration
}

// GetMemStressDuration returns the configured memory stress duration.
func GetMemStressDuration() time.Duration {
	return memStressDuration
}
