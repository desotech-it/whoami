package view

import (
	"github.com/desotech-it/whoami/app"
	"github.com/desotech-it/whoami/view/util"
	"html/template"
	"io"
)

type ImageView struct {
	Title    string
	Filename string
	AltText  string
	Info     app.WhoamiInfo
	Request  string
}

var imageViewTemplateFiles = []string{
	"template/image.tmpl",
	"template/base.tmpl",
	"template/whoami.tmpl",
	"template/request.tmpl",
}

func (v *ImageView) Write(w io.Writer) error {
	// TODO: handle error during template parsing
	template := template.Must(template.ParseFiles(imageViewTemplateFiles...))
	return template.Execute(w, v)
}

func (v *ImageView) WriteAsPlainText(w io.Writer) {
	util.WriteImageAsText(w, v.Filename)
	util.WriteWhoamiInfoAsText(w, v.Info, v.Request)
}
