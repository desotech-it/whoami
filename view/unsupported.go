package view

import (
	"html/template"
	"io"
)

type UnsupportedView struct {
	Title   string
	Message string
	Link    string
}

var unsupportedViewTemplateFiles = []string{
	"template/unsupported.tmpl",
	"template/base.tmpl",
}

func (v *UnsupportedView) Write(w io.Writer) error {
	// TODO: handle error during template parsing
	template := template.Must(template.ParseFiles(unsupportedViewTemplateFiles...))
	return template.Execute(w, v)
}
