package app

import (
	"os"
	"strings"
)

// ResourceLimits holds container resource limits from cgroup.
type ResourceLimits struct {
	CPUQuota  string
	CPUPeriod string
	MemLimit  string
	MemUsage  string
}

// GetResourceLimits reads cgroup files to determine CPU and memory limits.
// Supports both cgroup v1 and v2. Gracefully returns "N/A" when not in a container.
func GetResourceLimits() ResourceLimits {
	rl := ResourceLimits{
		CPUQuota:  "N/A",
		CPUPeriod: "N/A",
		MemLimit:  "N/A",
		MemUsage:  "N/A",
	}

	// Try cgroup v2 first
	if cpuMax := readFileContent("/sys/fs/cgroup/cpu.max"); cpuMax != "" {
		parts := strings.Fields(cpuMax)
		if len(parts) >= 2 {
			rl.CPUQuota = parts[0]
			rl.CPUPeriod = parts[1]
		}
	} else {
		// Fallback to cgroup v1
		rl.CPUQuota = readFileContentOrDefault("/sys/fs/cgroup/cpu/cpu.cfs_quota_us", "N/A")
		rl.CPUPeriod = readFileContentOrDefault("/sys/fs/cgroup/cpu/cpu.cfs_period_us", "N/A")
	}

	// Memory limits
	if memMax := readFileContent("/sys/fs/cgroup/memory.max"); memMax != "" {
		rl.MemLimit = memMax
		rl.MemUsage = readFileContentOrDefault("/sys/fs/cgroup/memory.current", "N/A")
	} else {
		rl.MemLimit = readFileContentOrDefault("/sys/fs/cgroup/memory/memory.limit_in_bytes", "N/A")
		rl.MemUsage = readFileContentOrDefault("/sys/fs/cgroup/memory/memory.usage_in_bytes", "N/A")
	}

	return rl
}

func readFileContent(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(data))
}

func readFileContentOrDefault(path, fallback string) string {
	if content := readFileContent(path); content != "" {
		return content
	}
	return fallback
}
