package view

import (
	"fmt"
	"io"
	"time"
)

type delayView struct {
	baseView
	Duration time.Duration
	LabInfo  *LabInfo
}

func NewDelayView(title string, duration time.Duration) View {
	return &delayView{
		baseView: baseView{title},
		Duration: duration,
		LabInfo: &LabInfo{
			Description: "Responds after a configurable delay. Useful for testing timeout configurations in Kubernetes probes, Ingress controllers, and service mesh policies. The maximum delay is 60 seconds.",
			Usage: `# Delay for 5 seconds:
curl http://<service>/delay?duration=5s

# Delay for 500 milliseconds:
curl http://<service>/delay?duration=500ms

# Test probe timeout (default 1s):
livenessProbe:
  httpGet:
    path: /delay?duration=2s
    port: 80
  timeoutSeconds: 3

# Test Ingress timeout:
# nginx.ingress.kubernetes.io/proxy-read-timeout: "10"`,
		},
	}
}

func (v *delayView) Write(w io.Writer) error {
	return delayTemplate.Execute(w, v)
}

func (v *delayView) WriteAsText(w io.Writer) {
	fmt.Fprintf(w, "Response delayed by %v\n", v.Duration)
}
