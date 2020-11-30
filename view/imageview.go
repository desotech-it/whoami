package view

import (
	"html/template"
	"io"
)

type ImageView struct {
	Title    string
	Filename string
	AltText  string
}

var templateFiles = []string{
	"template/image.tmpl",
	"template/base.tmpl",
}

func (v *ImageView) Write(w io.Writer) error {
	// TODO: handle error during template parsing
	template := template.Must(template.ParseFiles(templateFiles...))
	return template.Execute(w, v)
}
