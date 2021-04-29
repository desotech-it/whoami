package view

import (
	"github.com/desotech-it/whoami/app"
	"html/template"
	"io"
)

type MemStressView struct {
	Title string
	Stats app.MemStats
}

var memStressViewTemplateFiles = []string{
	"template/memstress.tmpl",
	"template/base.tmpl",
}

func (v *MemStressView) Write(w io.Writer) error {
	// TODO: handle error during template parsing
	template := template.Must(template.ParseFiles(memStressViewTemplateFiles...))
	return template.Execute(w, v)
}
