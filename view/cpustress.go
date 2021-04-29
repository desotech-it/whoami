package view

import (
	"github.com/desotech-it/whoami/app"
	"html/template"
	"io"
)

type CPUStressView struct {
	Title string
	Stats []app.CPUStats
}

var CPUStressViewTemplateFiles = []string{
	"template/cpustress.tmpl",
	"template/base.tmpl",
}

func (v *CPUStressView) Write(w io.Writer) error {
	// TODO: handle error during template parsing
	template := template.Must(template.ParseFiles(CPUStressViewTemplateFiles...))
	return template.Execute(w, v)
}
