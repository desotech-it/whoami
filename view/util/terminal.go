package util

import (
	"bytes"
	"fmt"
	"io"
	"path/filepath"
	"sort"
	"strings"
	"unicode/utf8"

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
			for _, addr := range info.Addresses[iface] {
				table.Append([]string{iface, addr})
			}
		}
		table.Render()
	}

	fmt.Fprintln(w)

	// HTTP Request and Client Info side by side
	var leftBuf, rightBuf bytes.Buffer

	// Left: HTTP Request
	{
		table := tablewriter.NewTable(&leftBuf)
		table.Header("HTTP Request")
		requestClean := strings.TrimSpace(strings.ReplaceAll(request, "\r\n", "\n"))
		for _, line := range strings.Split(requestClean, "\n") {
			if line != "" {
				table.Append([]string{line})
			}
		}
		table.Render()
	}

	// Right: Client Info
	{
		table := tablewriter.NewTable(&rightBuf)
		table.Header("Key", "Value")
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

	printSideBySide(w, leftBuf.String(), rightBuf.String(), 2)
}

// printSideBySide renders two multi-line strings side by side with a gap between them.
func printSideBySide(w io.Writer, left, right string, gap int) {
	leftLines := strings.Split(strings.TrimRight(left, "\n"), "\n")
	rightLines := strings.Split(strings.TrimRight(right, "\n"), "\n")

	maxWidth := 0
	for _, line := range leftLines {
		if width := utf8.RuneCountInString(line); width > maxWidth {
			maxWidth = width
		}
	}

	maxRows := len(leftLines)
	if len(rightLines) > maxRows {
		maxRows = len(rightLines)
	}

	for i := 0; i < maxRows; i++ {
		var l, r string
		if i < len(leftLines) {
			l = leftLines[i]
		}
		if i < len(rightLines) {
			r = rightLines[i]
		}
		padding := maxWidth - utf8.RuneCountInString(l) + gap
		fmt.Fprintf(w, "%s%s%s\n", l, strings.Repeat(" ", padding), r)
	}
}

func WriteImageAsText(w io.Writer, imageFilename string) {
	options := convert.DefaultOptions
	options.Ratio = 0.25
	converter := convert.NewImageConverter()
	imageLocation := filepath.Join("static", "images", imageFilename)
	fmt.Fprint(w, converter.ImageFile2ASCIIString(imageLocation, &options))
}
