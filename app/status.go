package app

import (
	"os"
	"time"
)

var isReady bool = false
var isHealthy bool = false

func init() {
	readinessDelay := 15 * time.Second
	if newValue, err := time.ParseDuration(os.Getenv("READINESS_DELAY")); err == nil {
		readinessDelay = newValue
	}

	go func() {
		readinessTimer := time.NewTimer(readinessDelay)
		<-readinessTimer.C
		isReady = true
	}()

	healthDelay, err := time.ParseDuration(os.Getenv("HEALTH_DELAY"))
	if err == nil {
		go func() {
			healthTimer := time.NewTimer(healthDelay)
			<-healthTimer.C
			isHealthy = true
		}()
	} else {
		isHealthy = true
	}

	go func() {
		readinessTimer := time.NewTimer(readinessDelay)
		<-readinessTimer.C
		isReady = true
	}()
}

func GetHealth() int {
	if isHealthy {
		return 200
	} else {
		return 503
	}
}

func GetReadiness() int {
	if isReady {
		return 200
	} else {
		return 503
	}
}
