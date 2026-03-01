package view

import (
	"fmt"
	"io"
)

type statusCodeView struct {
	baseView
	Code    int
	Text    string
	LabInfo *LabInfo
}

func NewStatusCodeView(title string, code int, text string) View {
	return &statusCodeView{
		baseView: baseView{title},
		Code:     code,
		Text:     text,
		LabInfo: &LabInfo{
			Description: "Returns a specific HTTP status code. Useful for testing how Kubernetes handles different response codes in health probes, how Ingress controllers handle errors, and how service meshes implement retry policies and circuit breakers.",
			Usage: `# Return specific status codes:
curl -v http://<service>/status?code=200
curl -v http://<service>/status?code=503
curl -v http://<service>/status?code=429

# Test with liveness probe:
livenessProbe:
  httpGet:
    path: /status?code=503
    port: 80

# Test Istio retry policy:
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
spec:
  http:
    - retries:
        attempts: 3
        retryOn: 5xx`,
		},
	}
}

func (v *statusCodeView) Write(w io.Writer) error {
	return statusCodeTemplate.Execute(w, v)
}

func (v *statusCodeView) WriteAsText(w io.Writer) {
	fmt.Fprintf(w, "Status: %d %s\n", v.Code, v.Text)
}
