package view

import (
	"fmt"
	"io"

	"github.com/desotech-it/whoami/app"
)

type memStressView struct {
	baseView
	Stats app.MemStats
}

func NewMemStressView(title string, stats app.MemStats) View {
	return &memStressView{
		baseView{title},
		stats,
	}
}

func (v *memStressView) Write(w io.Writer) error {
	t := memStressTemlate
	return t.Execute(w, v)
}

func (v *memStressView) WriteAsText(w io.Writer) {
	fmt.Fprintf(w, "This feature is not yet implemented for plain-text viewing.")
}
