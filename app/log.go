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

	logInterval := 2 * time.Second
	if maybeNewLogInterval, ok := os.LookupEnv("LOG_INTERVAL"); ok {
		if newLogInterval, err := ParseDuration(maybeNewLogInterval); err == nil {
			logInterval = newLogInterval
		} else {
			LogWarn.Warnf("failed to parse value for LOG_INTERVAL: %v. Defaulting to %v.", err, logInterval)
		}
	}

	if logInterval > 0 {
		go logCPUUsage(logInterval)
		go logMemoryUsage(logInterval)
	}
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
