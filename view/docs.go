package view

import (
	"fmt"
	"io"
)

// DocEndpoint describes a single API endpoint for the docs page.
type DocEndpoint struct {
	Path        string
	Method      string
	Description string
	Params      string
}

// DocCategory groups endpoints by category.
type DocCategory struct {
	Name      string
	Color     string
	Endpoints []DocEndpoint
}

type docsView struct {
	baseView
	Categories []DocCategory
}

func NewDocsView(title string) View {
	return &docsView{
		baseView: baseView{title},
		Categories: []DocCategory{
			{
				Name:  "System Info",
				Color: "primary",
				Endpoints: []DocEndpoint{
					{Path: "/", Method: "GET", Description: "Hostname, network interfaces, HTTP request, and client info"},
					{Path: "/env", Method: "GET", Description: "All environment variables (ConfigMap, Secret, Downward API)"},
					{Path: "/k8s", Method: "GET", Description: "Kubernetes pod metadata via Downward API"},
					{Path: "/resources", Method: "GET", Description: "Container resource limits from cgroup (CPU/memory)"},
					{Path: "/volumes", Method: "GET", Description: "Filesystem mount points (PV, ConfigMap, Secret volumes)"},
					{Path: "/osinfo", Method: "GET", Description: "OS, architecture, Go version, uptime"},
					{Path: "/dns", Method: "GET", Description: "DNS configuration (/etc/resolv.conf, CoreDNS)"},
				},
			},
			{
				Name:  "HTTP Tools",
				Color: "info",
				Endpoints: []DocEndpoint{
					{Path: "/headers", Method: "GET", Description: "Echo all received HTTP headers"},
					{Path: "/status", Method: "GET", Description: "Return a specific HTTP status code", Params: "?code=503"},
					{Path: "/delay", Method: "GET", Description: "Respond after a configurable delay", Params: "?duration=5s"},
					{Path: "/counter", Method: "GET", Description: "Request counter for load balancing verification"},
				},
			},
			{
				Name:  "Resilience",
				Color: "warning",
				Endpoints: []DocEndpoint{
					{Path: "/healthz", Method: "GET", Description: "Liveness probe endpoint (env: HEALTH_DELAY)"},
					{Path: "/readiness", Method: "GET", Description: "Readiness probe endpoint (env: READINESS_DELAY)"},
				},
			},
			{
				Name:  "Networking",
				Color: "success",
				Endpoints: []DocEndpoint{
					{Path: "/downstream", Method: "GET", Description: "Call another service and display response", Params: "?url=http://..."},
					{Path: "/ws", Method: "GET", Description: "WebSocket echo server for Ingress testing"},
				},
			},
			{
				Name:  "Stress Testing",
				Color: "danger",
				Endpoints: []DocEndpoint{
					{Path: "/cpustress", Method: "GET/POST", Description: "CPU load generation and monitoring"},
					{Path: "/memstress", Method: "GET/POST", Description: "Memory usage and stress testing"},
				},
			},
			{
				Name:  "Monitoring",
				Color: "secondary",
				Endpoints: []DocEndpoint{
					{Path: "/metrics", Method: "GET", Description: "Prometheus metrics endpoint"},
				},
			},
		},
	}
}

func (v *docsView) Write(w io.Writer) error {
	return docsTemplate.Execute(w, v)
}

func (v *docsView) WriteAsText(w io.Writer) {
	fmt.Fprintln(w, "WhoAmI - API Reference")
	fmt.Fprintln(w, "======================")
	fmt.Fprintln(w)
	for _, cat := range v.Categories {
		fmt.Fprintf(w, "[%s]\n", cat.Name)
		for _, ep := range cat.Endpoints {
			params := ""
			if ep.Params != "" {
				params = " " + ep.Params
			}
			fmt.Fprintf(w, "  %-8s %-20s %s\n", ep.Method, ep.Path+params, ep.Description)
		}
		fmt.Fprintln(w)
	}
}
