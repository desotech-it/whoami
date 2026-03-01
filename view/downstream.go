package view

import (
	"fmt"
	"io"

	"github.com/desotech-it/whoami/app"
)

type downstreamView struct {
	baseView
	Result  *app.DownstreamResult
	LabInfo *LabInfo
}

func NewDownstreamView(title string, result *app.DownstreamResult) View {
	return &downstreamView{
		baseView: baseView{title},
		Result:   result,
		LabInfo: &LabInfo{
			Description: "Makes an HTTP call to another service and displays the response. This demonstrates service-to-service communication in Kubernetes using ClusterIP services and DNS-based service discovery. Useful for testing network policies, service mesh routing, and distributed tracing.",
			Usage: `# Call another whoami instance:
curl "http://<service>/downstream?url=http://whoami-backend.default.svc/osinfo"

# K8s DNS service discovery formats:
# <service>                          (same namespace)
# <service>.<namespace>              (cross namespace)
# <service>.<namespace>.svc          (explicit)
# <service>.<namespace>.svc.cluster.local (FQDN)

# Test network policy:
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
spec:
  podSelector:
    matchLabels:
      app: whoami
  ingress:
    - from:
        - podSelector:
            matchLabels:
              app: frontend`,
		},
	}
}

func (v *downstreamView) Write(w io.Writer) error {
	return downstreamTemplate.Execute(w, v)
}

func (v *downstreamView) WriteAsText(w io.Writer) {
	if v.Result == nil {
		fmt.Fprintln(w, "Downstream Call: no URL provided")
		fmt.Fprintln(w, "Usage: /downstream?url=http://other-service/")
		return
	}
	fmt.Fprintf(w, "Downstream Call: %s\n", v.Result.URL)
	fmt.Fprintf(w, "  Duration: %v\n", v.Result.Duration)
	if v.Result.Error != "" {
		fmt.Fprintf(w, "  Error: %s\n", v.Result.Error)
	} else {
		fmt.Fprintf(w, "  Status: %d\n", v.Result.StatusCode)
		fmt.Fprintf(w, "  Body:\n%s\n", v.Result.Body)
	}
}
