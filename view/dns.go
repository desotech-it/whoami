package view

import (
	"fmt"
	"io"
	"strings"

	"github.com/desotech-it/whoami/app"
)

type dnsView struct {
	baseView
	DNS     app.DNSConfig
	LabInfo *LabInfo
}

func NewDNSView(title string, dns app.DNSConfig) View {
	return &dnsView{
		baseView: baseView{title},
		DNS:      dns,
		LabInfo: &LabInfo{
			Description: "Displays the DNS configuration (/etc/resolv.conf) of the container. In Kubernetes, CoreDNS provides cluster DNS resolution. This endpoint helps debug DNS issues, verify search domains, and understand the ndots:5 setting that affects how service names are resolved.",
			Usage: `# Default K8s resolv.conf:
# nameserver 10.96.0.10         (CoreDNS ClusterIP)
# search default.svc.cluster.local svc.cluster.local cluster.local
# options ndots:5

# Change DNS policy:
dnsPolicy: ClusterFirst        # default
dnsPolicy: Default             # use node's DNS
dnsPolicy: None                # custom DNS config

# Custom DNS config:
dnsConfig:
  nameservers:
    - 8.8.8.8
  searches:
    - custom.local

# Verify
curl http://<service>/dns`,
		},
	}
}

func (v *dnsView) Write(w io.Writer) error {
	return dnsTemplate.Execute(w, v)
}

func (v *dnsView) WriteAsText(w io.Writer) {
	fmt.Fprintln(w, "DNS Configuration:")
	fmt.Fprintf(w, "  Nameservers:   %s\n", strings.Join(v.DNS.Nameservers, ", "))
	fmt.Fprintf(w, "  Search:        %s\n", strings.Join(v.DNS.SearchDomains, ", "))
	fmt.Fprintf(w, "  Options:       %s\n", v.DNS.Options)
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Raw /etc/resolv.conf:")
	for _, line := range strings.Split(v.DNS.RawContent, "\n") {
		fmt.Fprintf(w, "  %s\n", line)
	}
}
