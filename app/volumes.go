package app

import (
	"bufio"
	"os"
	"strings"
)

// MountPoint represents a single filesystem mount.
type MountPoint struct {
	Device     string
	Path       string
	Filesystem string
	Options    string
}

// GetMountPoints parses /proc/mounts and returns mount points,
// filtering out noisy system mounts. Returns empty slice outside Linux.
func GetMountPoints() []MountPoint {
	file, err := os.Open("/proc/mounts")
	if err != nil {
		return nil
	}
	defer file.Close()

	var mounts []MountPoint
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 4 {
			continue
		}
		// Skip noisy system filesystems
		fs := fields[2]
		if isSystemFS(fs) {
			continue
		}
		mounts = append(mounts, MountPoint{
			Device:     fields[0],
			Path:       fields[1],
			Filesystem: fs,
			Options:    fields[3],
		})
	}
	return mounts
}

func isSystemFS(fs string) bool {
	switch fs {
	case "proc", "sysfs", "devpts", "devtmpfs", "cgroup", "cgroup2",
		"mqueue", "hugetlbfs", "debugfs", "tracefs", "securityfs",
		"pstore", "bpf", "autofs", "fusectl", "configfs", "ramfs":
		return true
	}
	return false
}
