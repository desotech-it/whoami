package app

import (
	"net"
	"os"
)

var Info WhoamiInfo

type WhoamiInfo struct {
	Hostname  string
	Addresses map[string][]string
}

func getWhoamiInfo() WhoamiInfo {
	hostname, _ := os.Hostname()
	ifaces, _ := net.Interfaces()

	addresses := make(map[string][]string)

	for _, iface := range ifaces {
		ifaceAddrs, _ := iface.Addrs()
		lenIfaceAddrs := len(ifaceAddrs)
		_, present := addresses[iface.Name]
		if !present {
			addresses[iface.Name] = make([]string, lenIfaceAddrs)
		}
		addrsList := addresses[iface.Name]
		i := 0
		for _, addr := range ifaceAddrs {
			addrsList[i] = addr.String()
			i++
		}
	}

	return WhoamiInfo{
		hostname,
		addresses,
	}
}

func init() {
	Info = getWhoamiInfo()
}
