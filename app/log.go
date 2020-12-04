package app

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)

	logInterval, err := time.ParseDuration(os.Getenv("WHOAMI_LOG_INTERVAL"))
	if err != nil {
		logInterval = 2 * time.Second
	}

	go logCPUUsage(logInterval)
	go logMemoryUsage(logInterval)
}

func LogRequest(r *http.Request) {
	logrus.WithFields(logrus.Fields{
		"method":         r.Method,
		"remote_address": r.RemoteAddr,
		"resource":       r.RequestURI,
	}).Infof("Handling %s request for %s from %s", r.Method, r.RequestURI, r.RemoteAddr)
}

func logCPUUsage(interval time.Duration) {
	for {
		timer := time.NewTimer(interval)
		<-timer.C
		cpuStats := CPULoad()
		logrus.WithFields(logrus.Fields{
			"cpu_load": cpuStats,
		}).Info("Logging CPU load")
	}
}

func logMemoryUsage(interval time.Duration) {
	for {
		timer := time.NewTimer(interval)
		<-timer.C
		memStats := MemInfo()
		logrus.WithFields(logrus.Fields{
			"used_memory":     memStats.UsedMemory,
			"total_memory":    memStats.TotalMemory,
			"percentage_used": memStats.PercentageUsed,
		}).Info("Logging memory usage")
	}
}
