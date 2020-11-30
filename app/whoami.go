package app

import (
	"net"
	"os"
)

var Info WhoamiInfo

type WhoamiInfo struct {
	Hostname  string
	Addresses []string
}

func getWhoamiInfo() WhoamiInfo {
	hostname, _ := os.Hostname()
	iaddrs, _ := net.InterfaceAddrs()
	var addresses []string
	// handle err
	for _, addr := range iaddrs {
		addresses = append(addresses, addr.String())
	}

	// _, _ = fmt.Fprintln(w, "RemoteAddr:", req.RemoteAddr)
	// if err := req.Write(w); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	return WhoamiInfo{
		hostname,
		addresses,
	}
}

func init() {
	Info = getWhoamiInfo()
}
