package view

import (
	"io"

	"github.com/desotech-it/whoami/app"
	"github.com/desotech-it/whoami/view/util"
)

type indexView struct {
	baseView
	Info    app.WhoamiInfo
	Request string
}

func NewIndexView(title string, info app.WhoamiInfo, request string) View {
	return &indexView{
		baseView{title},
		info,
		request,
	}
}

func (v *indexView) Write(w io.Writer) error {
	t := indexTemplate
	return t.Execute(w, v)
}

func (v *indexView) WriteAsText(w io.Writer) {
	util.WriteWhoamiInfoAsText(w, v.Info, v.Request)
}
