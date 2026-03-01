package view

import (
	"fmt"
	"io"

	"github.com/desotech-it/whoami/app"
)

type k8sView struct {
	baseView
	Meta    app.K8sMetadata
	LabInfo *LabInfo
}

func NewK8sView(title string, meta app.K8sMetadata) View {
	return &k8sView{
		baseView: baseView{title},
		Meta:     meta,
		LabInfo: &LabInfo{
			Description: "Displays Kubernetes pod metadata injected via the Downward API. This lets you verify that pod identity information (name, namespace, node, service account, IP) is correctly exposed to the application.",
			Usage: `# Add to your Pod spec:
env:
  - name: POD_NAME
    valueFrom:
      fieldRef:
        fieldPath: metadata.name
  - name: POD_NAMESPACE
    valueFrom:
      fieldRef:
        fieldPath: metadata.namespace
  - name: NODE_NAME
    valueFrom:
      fieldRef:
        fieldPath: spec.nodeName
  - name: POD_IP
    valueFrom:
      fieldRef:
        fieldPath: status.podIP
  - name: SERVICE_ACCOUNT
    valueFrom:
      fieldRef:
        fieldPath: spec.serviceAccountName

# Verify
curl http://<service>/k8s`,
		},
	}
}

func (v *k8sView) Write(w io.Writer) error {
	return k8sTemplate.Execute(w, v)
}

func (v *k8sView) WriteAsText(w io.Writer) {
	fmt.Fprintln(w, "Kubernetes Metadata:")
	fmt.Fprintf(w, "  Pod Name:        %s\n", v.Meta.PodName)
	fmt.Fprintf(w, "  Namespace:       %s\n", v.Meta.Namespace)
	fmt.Fprintf(w, "  Node Name:       %s\n", v.Meta.NodeName)
	fmt.Fprintf(w, "  Service Account: %s\n", v.Meta.ServiceAccount)
	fmt.Fprintf(w, "  Pod IP:          %s\n", v.Meta.PodIP)
}
