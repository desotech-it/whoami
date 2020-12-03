package app

import (
	"errors"

	"github.com/shirou/gopsutil/v3/cpu"
)

type CPUStats struct {
	Name string
	Load float64
}

func CPULoad() []float64 {
	data, _ := cpu.Percent(0, true)
	return data
}

func CPUInfo() []CPUStats {
	info, _ := cpu.Info()
	load, _ := cpu.Percent(0, true)

	stats := make([]CPUStats, len(info))

	if len(info) != len(load) {
		panic(errors.New("mismatch in CPU count between two methods"))
	}

	for i, cpuInfo := range info {
		stats[i].Name = cpuInfo.ModelName
		stats[i].Load = load[i]
	}

	return stats
}
