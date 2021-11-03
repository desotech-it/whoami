package util

import (
	"github.com/desotech-it/whoami/app"
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/qeesung/image2ascii/convert"
)

func WriteWhoamiInfoAsText(w io.Writer, info app.WhoamiInfo, request string, clientInfo map[string]string) {
	{
		table := tablewriter.NewWriter(w)
		table.SetHeader([]string{"Hostname"})
		table.Append([]string{info.Hostname})
		table.Render()
	}
	{
		table := tablewriter.NewWriter(w)
		table.SetHeader([]string{"IP", "Interface"})
		for iface, addrs := range info.Addresses {
			for _, addr := range addrs {
				table.Append([]string{addr, iface})
			}
		}
		table.Render()
	}
	{
		requestNoCarriageReturn := strings.ReplaceAll(request, "\r\n", "\n")
		table := tablewriter.NewWriter(w)
		table.SetAutoWrapText(false)
		table.SetHeader([]string{"Request"})
		table.Append([]string{requestNoCarriageReturn})
		table.Render()
	}
	{
		table := tablewriter.NewWriter(w)
		table.SetHeader([]string{"Client Info"})
		for k, v := range(clientInfo) {
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
