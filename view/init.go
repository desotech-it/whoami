package view

import (
	"html/template"
)

var (
	baseTemplate        *template.Template
	indexTemplate       *template.Template
	imageTemplate       *template.Template
	cpuStressTemplate   *template.Template
	unsupportedTemplate *template.Template
	memStressTemlate    *template.Template
)

func cloneFromTemplate(src *template.Template, filenames ...string) *template.Template {
	t := template.Must(src.Clone())
	return template.Must(t.ParseFiles(filenames...))
}

func parseAllTemplates() {
	baseTemplate = template.Must(template.ParseFiles("template/base.gohtml"))
	imageTemplate = cloneFromTemplate(baseTemplate, "template/image.gohtml", "template/whoami.gohtml", "template/request.gohtml")
	indexTemplate = cloneFromTemplate(baseTemplate, "template/index.gohtml", "template/whoami.gohtml", "template/request.gohtml")
	cpuStressTemplate = cloneFromTemplate(baseTemplate, "template/cpustress.gohtml")
	unsupportedTemplate = cloneFromTemplate(baseTemplate, "template/unsupported.gohtml")
	memStressTemlate = cloneFromTemplate(baseTemplate, "template/memstress.gohtml")
}

func init() {
	parseAllTemplates()
}
