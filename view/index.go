package view

import (
	"desotech/whoami/app"
	"html/template"
	"io"
)

type IndexView struct {
	Title   string
	Info    app.WhoamiInfo
	Request string
}

var indexViewTemplateFiles = []string{
	"template/index.tmpl",
	"template/base.tmpl",
	"template/whoami.tmpl",
	"template/request.tmpl",
}

func (v *IndexView) Write(w io.Writer) error {
	// TODO: handle error during template parsing
	template := template.Must(template.ParseFiles(indexViewTemplateFiles...))
	return template.Execute(w, v)
}
