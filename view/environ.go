package view

import (
	"fmt"
	"io"
	"strings"
)

// EnvVar represents a single environment variable.
type EnvVar struct {
	Key   string
	Value string
}

type envView struct {
	baseView
	Vars    []EnvVar
	LabInfo *LabInfo
}

func NewEnvView(title string, env []string) View {
	vars := make([]EnvVar, 0, len(env))
	for _, e := range env {
		k, v, _ := strings.Cut(e, "=")
		vars = append(vars, EnvVar{Key: k, Value: v})
	}
	return &envView{
		baseView: baseView{title},
		Vars:     vars,
		LabInfo: &LabInfo{
			Description: "Shows all environment variables visible to the container. In Kubernetes, environment variables can be injected from ConfigMaps, Secrets, and the Downward API. This endpoint lets you verify that your configuration is correctly propagated to the pod.",
			Usage: `# Inject from ConfigMap
kubectl create configmap app-config --from-literal=APP_COLOR=blue
---
env:
  - name: APP_COLOR
    valueFrom:
      configMapKeyRef:
        name: app-config
        key: APP_COLOR

# Inject from Secret
env:
  - name: DB_PASSWORD
    valueFrom:
      secretKeyRef:
        name: db-secret
        key: password

# Verify with curl
curl http://<service>/env`,
		},
	}
}

func (v *envView) Write(w io.Writer) error {
	return envTemplate.Execute(w, v)
}

func (v *envView) WriteAsText(w io.Writer) {
	fmt.Fprintln(w, "Environment Variables:")
	fmt.Fprintln(w)
	for _, e := range v.Vars {
		fmt.Fprintf(w, "  %s=%s\n", e.Key, e.Value)
	}
}
