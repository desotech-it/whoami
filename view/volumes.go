package view

import (
	"fmt"
	"io"

	"github.com/desotech-it/whoami/app"
)

type volumesView struct {
	baseView
	Mounts  []app.MountPoint
	LabInfo *LabInfo
}

func NewVolumesView(title string, mounts []app.MountPoint) View {
	return &volumesView{
		baseView: baseView{title},
		Mounts:   mounts,
		LabInfo: &LabInfo{
			Description: "Displays filesystem mount points visible to the container. In Kubernetes, volumes can be mounted from PersistentVolumes, ConfigMaps, Secrets, emptyDir, and hostPath. This endpoint helps verify that volumes are correctly mounted.",
			Usage: `# Mount a ConfigMap as a volume:
volumes:
  - name: config-vol
    configMap:
      name: my-config
volumeMounts:
  - name: config-vol
    mountPath: /etc/config

# Mount a PersistentVolumeClaim:
volumes:
  - name: data-vol
    persistentVolumeClaim:
      claimName: my-pvc
volumeMounts:
  - name: data-vol
    mountPath: /data

# Verify
curl http://<service>/volumes`,
		},
	}
}

func (v *volumesView) Write(w io.Writer) error {
	return volumesTemplate.Execute(w, v)
}

func (v *volumesView) WriteAsText(w io.Writer) {
	fmt.Fprintln(w, "Mount Points:")
	if len(v.Mounts) == 0 {
		fmt.Fprintln(w, "  (not available)")
		return
	}
	for _, m := range v.Mounts {
		fmt.Fprintf(w, "  %s on %s type %s (%s)\n", m.Device, m.Path, m.Filesystem, m.Options)
	}
}
