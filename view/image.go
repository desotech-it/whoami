package view

import (
	"io"

	"github.com/desotech-it/whoami/app"
	"github.com/desotech-it/whoami/view/util"
)

type imageView struct {
	baseView
	Filename   string
	AltText    string
	Info       app.WhoamiInfo
	Request    string
	ClientInfo map[string]string
}

func NewImageView(title string, filename string, altText string, info app.WhoamiInfo, request string, clientInfo map[string]string) View {
	return &imageView{
		baseView{title},
		filename,
		altText,
		info,
		request,
		clientInfo,
	}
}

func (v *imageView) Write(w io.Writer) error {
	t := imageTemplate
	return t.Execute(w, v)
}

func (v *imageView) WriteAsText(w io.Writer) {
	util.WriteImageAsText(w, v.Filename)
	util.WriteWhoamiInfoAsText(w, v.Info, v.Request, v.ClientInfo)
}
