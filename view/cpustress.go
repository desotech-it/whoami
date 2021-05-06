package view

import (
	"fmt"
	"io"

	"github.com/desotech-it/whoami/app"
)

type cpuStressView struct {
	baseView
	Stats []app.CPUStats
}

func NewCPUStressView(title string, stats []app.CPUStats) View {
	return &cpuStressView{
		baseView{title},
		stats,
	}
}

func (v *cpuStressView) Write(w io.Writer) error {
	t := cpuStressTemplate
	return t.Execute(w, v)
}

func (v *cpuStressView) WriteAsText(w io.Writer) {
	fmt.Fprintln(w, "This feature is not yet implemented for plain-text viewing.")
}
