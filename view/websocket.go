package view

import (
	"fmt"
	"io"
)

type websocketView struct {
	baseView
	LabInfo *LabInfo
}

func NewWebSocketView(title string) View {
	return &websocketView{
		baseView: baseView{title},
		LabInfo: &LabInfo{
			Description: "Provides a WebSocket echo server. When you send a message, the server echoes it back with the hostname prefix. This is useful for testing WebSocket support through Ingress controllers, which require specific annotations to handle WebSocket upgrades.",
			Usage: `# Test with wscat (npm install -g wscat):
wscat -c ws://<service>/ws

# Nginx Ingress annotations for WebSocket:
metadata:
  annotations:
    nginx.ingress.kubernetes.io/proxy-http-version: "1.1"
    nginx.ingress.kubernetes.io/proxy-set-headers: "Upgrade"

# The upgrade headers are typically handled automatically
# by Nginx Ingress, but explicit configuration may be needed
# for custom setups or other Ingress controllers.`,
		},
	}
}

func (v *websocketView) Write(w io.Writer) error {
	return websocketTemplate.Execute(w, v)
}

func (v *websocketView) WriteAsText(w io.Writer) {
	fmt.Fprintln(w, "WebSocket Echo Server")
	fmt.Fprintln(w, "Connect with: wscat -c ws://<host>/ws")
}
