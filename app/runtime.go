package app

import (
	"errors"
	"runtime"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

type CPUStats struct {
	Name string
	Load float64
}

type MemStats struct {
	UsedMemory     uint64  `json:"used_memory"`
	TotalMemory    uint64  `json:"total_memory"`
	PercentageUsed float64 `json:"percentage_used"`
}

func CPULoad() []float64 {
	data, _ := cpu.Percent(0, true)
	return data
}

func CPUInfo() []CPUStats {
	info, err := cpu.Info()
	if err != nil {
		panic(errors.New("unable to retrieve CPU info"))
	}

	if len(info) == 0 {
		panic(errors.New("CPU info is an empty list"))
	}

	load, err := cpu.Percent(0, true)
	if err != nil {
		panic(errors.New("unable to retrieve CPU load"))
	}

	stats := make([]CPUStats, len(load))
	cpuModelName := info[0].ModelName

	for i := range load {
		stats[i].Name = cpuModelName
		stats[i].Load = load[i]
	}

	return stats
}

func MemInfo() MemStats {
	stats, _ := mem.VirtualMemory()

	return MemStats{
		UsedMemory:     stats.Used,
		TotalMemory:    stats.Total,
		PercentageUsed: stats.UsedPercent,
	}
}

func IsOSDarwin() bool {
	return runtime.GOOS == "darwin"
}
