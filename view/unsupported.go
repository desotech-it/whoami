package view

import (
	"fmt"
	"io"
)

type unsupportedView struct {
	baseView
	Message string
	Link    string
}

func NewUnsupportedView(title string, message string, link string) View {
	return &unsupportedView{
		baseView{title},
		message,
		link,
	}
}

func (v *unsupportedView) Write(w io.Writer) error {
	t := unsupportedTemplate
	return t.Execute(w, v)
}

func (v *unsupportedView) WriteAsText(w io.Writer) {
	fmt.Fprintf(w, "%s\nFor futher info check out this link: %s\n", v.Message, v.Link)
}
