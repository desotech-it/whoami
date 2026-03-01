package view

import (
	"fmt"
	"io"
)

type counterView struct {
	baseView
	Count   int64
	LabInfo *LabInfo
}

func NewCounterView(title string, count int64) View {
	return &counterView{
		baseView: baseView{title},
		Count:    count,
		LabInfo: &LabInfo{
			Description: "Shows the total number of HTTP requests served by this instance since startup. When running multiple replicas behind a Kubernetes Service, each pod has its own counter. This helps demonstrate load balancing behavior, sticky sessions, and scaling effects.",
			Usage: `# Watch request distribution across replicas:
kubectl scale deployment whoami --replicas=3
for i in $(seq 1 100); do curl -s http://<service>/counter; done

# With sticky sessions (session affinity):
apiVersion: v1
kind: Service
spec:
  sessionAffinity: ClientIP
  sessionAffinityConfig:
    clientIP:
      timeoutSeconds: 10800`,
		},
	}
}

func (v *counterView) Write(w io.Writer) error {
	return counterTemplate.Execute(w, v)
}

func (v *counterView) WriteAsText(w io.Writer) {
	fmt.Fprintf(w, "Request count: %d\n", v.Count)
}
