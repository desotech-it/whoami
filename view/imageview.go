package view

import (
	"desotech/whoami/app"
	"html/template"
	"io"
)

type ImageView struct {
	Title    string
	Filename string
	AltText  string
	Info     app.WhoamiInfo
}

var templateFiles = []string{
	"template/image.tmpl",
	"template/base.tmpl",
	"template/whoami.tmpl",
}

func (v *ImageView) Write(w io.Writer) error {
	// TODO: handle error during template parsing
	template := template.Must(template.ParseFiles(templateFiles...))
	return template.Execute(w, v)
}
