package app

import (
	"errors"

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

func MemInfo() MemStats {
	stats, _ := mem.VirtualMemory()

	return MemStats{
		UsedMemory:     stats.Used,
		TotalMemory:    stats.Total,
		PercentageUsed: stats.UsedPercent,
	}
}
