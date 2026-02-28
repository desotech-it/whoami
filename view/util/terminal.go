package util

import (
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/desotech-it/whoami/app"
	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/tw"
	"github.com/qeesung/image2ascii/convert"
)

func WriteWhoamiInfoAsText(w io.Writer, info app.WhoamiInfo, request string, clientInfo map[string]string) {
	{
		table := tablewriter.NewTable(w)
		table.Header("Hostname")
		table.Append([]string{info.Hostname})
		table.Render()
	}
	{
		table := tablewriter.NewTable(w)
		table.Header("IP", "Interface")
		for iface, addrs := range info.Addresses {
			for _, addr := range addrs {
				table.Append([]string{addr, iface})
			}
		}
		table.Render()
	}
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
		table.Header("Request")
		table.Append([]string{requestNoCarriageReturn})
		table.Render()
	}
	{
		table := tablewriter.NewTable(w)
		table.Header("Client Info")
		for k, v := range clientInfo {
			table.Append([]string{fmt.Sprintf("%s=%s", k, v)})
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
