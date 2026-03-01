package app

import (
	"bufio"
	"os"
	"strings"
)

// DNSConfig holds parsed DNS configuration.
type DNSConfig struct {
	RawContent    string
	Nameservers   []string
	SearchDomains []string
	Options       string
}

// GetDNSConfig reads and parses /etc/resolv.conf.
// Returns empty struct if the file is not readable.
func GetDNSConfig() DNSConfig {
	data, err := os.ReadFile("/etc/resolv.conf")
	if err != nil {
		return DNSConfig{RawContent: "(not available)"}
	}

	config := DNSConfig{
		RawContent: string(data),
	}

	scanner := bufio.NewScanner(strings.NewReader(config.RawContent))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}
		switch fields[0] {
		case "nameserver":
			config.Nameservers = append(config.Nameservers, fields[1])
		case "search":
			config.SearchDomains = fields[1:]
		case "options":
			config.Options = strings.Join(fields[1:], " ")
		}
	}
	return config
}
