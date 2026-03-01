package util

import (
	"fmt"
	"io"
	"path/filepath"
	"sort"
	"strings"

	"github.com/desotech-it/whoami/app"
	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/tw"
	"github.com/qeesung/image2ascii/convert"
)

func WriteWhoamiInfoAsText(w io.Writer, info app.WhoamiInfo, request string, clientInfo map[string]string) {
	// Header banner
	fmt.Fprintln(w, "=== WhoAmI ===")
	fmt.Fprintln(w)

	// Hostname section
	fmt.Fprintf(w, "Hostname: %s\n", info.Hostname)
	fmt.Fprintln(w)

	// Network Interfaces section
	{
		table := tablewriter.NewTable(w)
		table.Header("Interface", "Address")

		// Sort interfaces for consistent output
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
	fmt.Fprintln(w)

	// HTTP Request section
	{
		requestNoCarriageReturn := strings.ReplaceAll(request, "\r\n", "\n")
		table := tablewriter.NewTable(w,
			tablewriter.WithConfig(tablewriter.Config{
				Row: tw.CellConfig{
					Formatting: tw.CellFormatting{
						AutoWrap: tw.WrapNone,
					},
				},
			}),
		)
		table.Header("HTTP Request")
		table.Append([]string{requestNoCarriageReturn})
		table.Render()
	}
	fmt.Fprintln(w)

	// Client Info section
	{
		table := tablewriter.NewTable(w)
		table.Header("Key", "Value")

		// Sort keys for consistent output
		keys := make([]string, 0, len(clientInfo))
		for k := range clientInfo {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			table.Append([]string{k, clientInfo[k]})
		}
		table.Render()
	}
}

func WriteImageAsText(w io.Writer, imageFilename string) {
	options := convert.DefaultOptions
	options.Ratio = 0.25
	converter := convert.NewImageConverter()
	imageLocation := filepath.Join("static", "images", imageFilename)
	fmt.Fprint(w, converter.ImageFile2ASCIIString(imageLocation, &options))
}
