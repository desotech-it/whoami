package view

import (
	"fmt"
	"io"

	"github.com/desotech-it/whoami/app"
)

type resourcesView struct {
	baseView
	Limits  app.ResourceLimits
	LabInfo *LabInfo
}

func NewResourcesView(title string, limits app.ResourceLimits) View {
	return &resourcesView{
		baseView: baseView{title},
		Limits:   limits,
		LabInfo: &LabInfo{
			Description: "Shows container resource limits (CPU and memory) as configured by Kubernetes through cgroups. This helps verify that resource requests and limits defined in the Pod spec are correctly applied to the container.",
			Usage: `# Set resource limits in Pod spec:
resources:
  requests:
    memory: "64Mi"
    cpu: "250m"
  limits:
    memory: "128Mi"
    cpu: "500m"

# Verify limits are applied
curl http://<service>/resources

# Kubernetes QoS Classes:
# - Guaranteed: requests == limits for all containers
# - Burstable: at least one request or limit set
# - BestEffort: no requests or limits set`,
		},
	}
}

func (v *resourcesView) Write(w io.Writer) error {
	return resourcesTemplate.Execute(w, v)
}

func (v *resourcesView) WriteAsText(w io.Writer) {
	fmt.Fprintln(w, "Resource Limits (cgroup):")
	fmt.Fprintf(w, "  CPU Quota:    %s\n", v.Limits.CPUQuota)
	fmt.Fprintf(w, "  CPU Period:   %s\n", v.Limits.CPUPeriod)
	fmt.Fprintf(w, "  Memory Limit: %s\n", v.Limits.MemLimit)
	fmt.Fprintf(w, "  Memory Usage: %s\n", v.Limits.MemUsage)
}
