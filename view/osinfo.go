package view

import (
	"fmt"
	"io"

	"github.com/desotech-it/whoami/app"
)

type osinfoView struct {
	baseView
	Info    app.OSInfo
	LabInfo *LabInfo
}

func NewOSInfoView(title string, info app.OSInfo) View {
	return &osinfoView{
		baseView: baseView{title},
		Info:     info,
		LabInfo: &LabInfo{
			Description: "Shows operating system, CPU architecture, Go runtime version, and container uptime. Useful for verifying multi-architecture deployments (amd64 vs arm64), confirming which node architecture a pod is running on, and testing nodeSelector or tolerations.",
			Usage: `# Deploy on a specific architecture:
nodeSelector:
  kubernetes.io/arch: arm64

# Verify the architecture
curl http://<service>/osinfo

# Use node affinity for more control:
affinity:
  nodeAffinity:
    requiredDuringSchedulingIgnoredDuringExecution:
      nodeSelectorTerms:
        - matchExpressions:
            - key: kubernetes.io/arch
              operator: In
              values: ["amd64"]`,
		},
	}
}

func (v *osinfoView) Write(w io.Writer) error {
	return osinfoTemplate.Execute(w, v)
}

func (v *osinfoView) WriteAsText(w io.Writer) {
	fmt.Fprintln(w, "OS & Runtime Info:")
	fmt.Fprintf(w, "  Hostname:   %s\n", v.Info.Hostname)
	fmt.Fprintf(w, "  OS:         %s\n", v.Info.OS)
	fmt.Fprintf(w, "  Arch:       %s\n", v.Info.Arch)
	fmt.Fprintf(w, "  Go Version: %s\n", v.Info.GoVersion)
	fmt.Fprintf(w, "  CPUs:       %d\n", v.Info.NumCPU)
	fmt.Fprintf(w, "  Uptime:     %s\n", v.Info.Uptime)
}
