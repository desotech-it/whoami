package util

import (
	"desotech/whoami/app"
	"io"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func WriteWhoamiInfoAsText(w io.Writer, info app.WhoamiInfo, request string) {
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
		table.SetHeader([]string{"Request"})
		table.Append([]string{requestNoCarriageReturn})
		table.Render()
	}
}
