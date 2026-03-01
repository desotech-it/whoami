package app

import "os"

// K8sMetadata holds Kubernetes pod metadata typically injected
// via the Downward API.
type K8sMetadata struct {
	PodName        string
	Namespace      string
	NodeName       string
	ServiceAccount string
	PodIP          string
}

// GetK8sMetadata reads Kubernetes metadata from well-known environment
// variables set via the Downward API. Falls back to hostname for PodName.
func GetK8sMetadata() K8sMetadata {
	podName := os.Getenv("POD_NAME")
	if podName == "" {
		podName = Info.Hostname
	}
	return K8sMetadata{
		PodName:        podName,
		Namespace:      envOrDefault("POD_NAMESPACE", "(not set)"),
		NodeName:       envOrDefault("NODE_NAME", "(not set)"),
		ServiceAccount: envOrDefault("SERVICE_ACCOUNT", "(not set)"),
		PodIP:          envOrDefault("POD_IP", "(not set)"),
	}
}

func envOrDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
