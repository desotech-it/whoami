package view

import (
	"html/template"
	"sync"
	"time"

	"github.com/desotech-it/whoami/app"
)

var (
	once                sync.Once
	baseTemplate        *template.Template
	indexTemplate       *template.Template
	imageTemplate       *template.Template
	cpuStressTemplate   *template.Template
	unsupportedTemplate *template.Template
	memStressTemplate   *template.Template

	// System Info
	envTemplate       *template.Template
	k8sTemplate       *template.Template
	resourcesTemplate *template.Template
	volumesTemplate   *template.Template
	osinfoTemplate    *template.Template
	dnsTemplate       *template.Template

	// HTTP Tools
	headersTemplate    *template.Template
	statusCodeTemplate *template.Template
	delayTemplate      *template.Template
	counterTemplate    *template.Template

	// Networking
	downstreamTemplate *template.Template
	websocketTemplate  *template.Template

	// Docs
	docsTemplate *template.Template
)

func cloneFromTemplate(src *template.Template, filenames ...string) *template.Template {
	t := template.Must(src.Clone())
	return template.Must(t.ParseFiles(filenames...))
}

func withLabInfo(filenames ...string) []string {
	return append(filenames, "template/labinfo.gohtml")
}

var funcMap = template.FuncMap{
	"year":     func() int { return time.Now().Year() },
	"hostname": func() string { return app.Info.Hostname },
}

func parseAllTemplates() {
	baseTemplate = template.Must(template.New("base.gohtml").Funcs(funcMap).ParseFiles("template/base.gohtml"))

	// Existing pages
	imageTemplate = cloneFromTemplate(baseTemplate, "template/image.gohtml", "template/whoami.gohtml", "template/request.gohtml", "template/clientinfo.gohtml")
	indexTemplate = cloneFromTemplate(baseTemplate, "template/index.gohtml", "template/request.gohtml", "template/clientinfo.gohtml")
	cpuStressTemplate = cloneFromTemplate(baseTemplate, "template/cpustress.gohtml")
	unsupportedTemplate = cloneFromTemplate(baseTemplate, "template/unsupported.gohtml")
	memStressTemplate = cloneFromTemplate(baseTemplate, "template/memstress.gohtml")

	// System Info
	envTemplate = cloneFromTemplate(baseTemplate, withLabInfo("template/environ.gohtml")...)
	k8sTemplate = cloneFromTemplate(baseTemplate, withLabInfo("template/k8s.gohtml")...)
	resourcesTemplate = cloneFromTemplate(baseTemplate, withLabInfo("template/resources.gohtml")...)
	volumesTemplate = cloneFromTemplate(baseTemplate, withLabInfo("template/volumes.gohtml")...)
	osinfoTemplate = cloneFromTemplate(baseTemplate, withLabInfo("template/osinfo.gohtml")...)
	dnsTemplate = cloneFromTemplate(baseTemplate, withLabInfo("template/dns.gohtml")...)

	// HTTP Tools
	headersTemplate = cloneFromTemplate(baseTemplate, withLabInfo("template/headers.gohtml")...)
	statusCodeTemplate = cloneFromTemplate(baseTemplate, withLabInfo("template/statuscode.gohtml")...)
	delayTemplate = cloneFromTemplate(baseTemplate, withLabInfo("template/delay.gohtml")...)
	counterTemplate = cloneFromTemplate(baseTemplate, withLabInfo("template/counter.gohtml")...)

	// Networking
	downstreamTemplate = cloneFromTemplate(baseTemplate, withLabInfo("template/downstream.gohtml")...)
	websocketTemplate = cloneFromTemplate(baseTemplate, withLabInfo("template/websocket.gohtml")...)

	// Docs
	docsTemplate = cloneFromTemplate(baseTemplate, "template/docs.gohtml")
}

func init() {
	once.Do(parseAllTemplates)
}
