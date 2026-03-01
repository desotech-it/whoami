package view

import (
	"fmt"
	"io"
	"net/http"
	"sort"
)

type headersView struct {
	baseView
	Headers http.Header
	LabInfo *LabInfo
}

func NewHeadersView(title string, headers http.Header) View {
	return &headersView{
		baseView: baseView{title},
		Headers:  headers,
		LabInfo: &LabInfo{
			Description: "Echoes all HTTP headers received by the server. In Kubernetes, Ingress controllers and service meshes add headers like X-Forwarded-For, X-Real-IP, X-Request-ID. This endpoint helps verify header injection, routing rules, and proxy configuration.",
			Usage: `# Test through Ingress
curl -H "X-Custom: test" http://<ingress-host>/headers

# Common Ingress headers to look for:
# X-Forwarded-For     - client IP chain
# X-Forwarded-Proto   - original protocol (http/https)
# X-Real-IP           - actual client IP
# X-Request-ID        - unique request identifier

# Nginx Ingress annotation to add headers:
# nginx.ingress.kubernetes.io/configuration-snippet: |
#   proxy_set_header X-Custom-Header "value";`,
		},
	}
}

func (v *headersView) Write(w io.Writer) error {
	return headersTemplate.Execute(w, v)
}

func (v *headersView) WriteAsText(w io.Writer) {
	fmt.Fprintln(w, "HTTP Headers:")
	keys := make([]string, 0, len(v.Headers))
	for k := range v.Headers {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		for _, v := range v.Headers[k] {
			fmt.Fprintf(w, "  %s: %s\n", k, v)
		}
	}
}
