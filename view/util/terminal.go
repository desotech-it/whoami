package util

import (
	"fmt"
	"io"
	"path/filepath"
	"sort"
	"strings"

	"github.com/desotech-it/whoami/app"
	"github.com/olekukonko/tablewriter"
	"github.com/qeesung/image2ascii/convert"
)

func WriteWhoamiInfoAsText(w io.Writer, info app.WhoamiInfo, request string, clientInfo map[string]string) {
	// Hostname
	fmt.Fprintf(w, "Hostname: %s\n\n", info.Hostname)

	// Network Interfaces
	{
		table := tablewriter.NewTable(w)
		table.Header("Interface", "Address")

		ifaces := make([]string, 0, len(info.Addresses))
		for iface := range info.Addresses {
			ifaces = append(ifaces, iface)
		}
		sort.Strings(ifaces)

		for _, iface := range ifaces {
			addrs := info.Addresses[iface]
			for _, addr := range addrs {
				table.Append([]string{iface, addr})
			}
		}
		table.Render()
	}

	// HTTP Request (plain text, no table)
	fmt.Fprintln(w)
	requestClean := strings.ReplaceAll(request, "\r\n", "\n")
	for _, line := range strings.Split(requestClean, "\n") {
		fmt.Fprintf(w, "  %s\n", line)
	}

	// Client Info (single line, compact)
	fmt.Fprintln(w)
	keys := make([]string, 0, len(clientInfo))
	for k := range clientInfo {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%s=%s", k, clientInfo[k]))
	}
	fmt.Fprintf(w, "Client: %s\n", strings.Join(parts, "  "))
}

func WriteImageAsText(w io.Writer, imageFilename string) {
	options := convert.DefaultOptions
	options.Ratio = 0.25
	converter := convert.NewImageConverter()
	imageLocation := filepath.Join("static", "images", imageFilename)
	fmt.Fprint(w, converter.ImageFile2ASCIIString(imageLocation, &options))
}
