package app

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var isReady bool = false
var isHealthy bool = false

var (
	LogWarn = &logrus.Logger{
		Out: os.Stderr,
		Formatter: new(logrus.TextFormatter),
		Hooks: make(logrus.LevelHooks),
		Level: logrus.WarnLevel,
	}

)

func init() {
	readinessDelay := 15 * time.Second
	if maybeNewReadinessDelay, ok := os.LookupEnv("READINESS_DELAY"); ok {
		if newReadinessDelay, err := ParseDuration(maybeNewReadinessDelay); err != nil {
			if newReadinessDelay < 0 {
				LogWarn.Warnf("READINESS_DELAY cannot be a negative value. Defaulting to %v. You passed: %v.", readinessDelay, maybeNewReadinessDelay)
			} else {
				readinessDelay = newReadinessDelay
			}
		} else {
			LogWarn.Warnf("failed to parse value for READINESS_DELAY: %v. Defaulting to %v.", err, readinessDelay)
		}
	}

	go func() {
		readinessTimer := time.NewTimer(readinessDelay)
		<-readinessTimer.C
		isReady = true
	}()

	healthDelay := time.Duration(0)
	if maybeNewHealthDelay, ok := os.LookupEnv("HEALTH_DELAY"); ok {
		if newHealthDelay, err := ParseDuration(maybeNewHealthDelay); err != nil {
			if newHealthDelay < 0 {
				LogWarn.Warnf("HEALTH_DELAY cannot be a negative value. Defaulting to %v. You passed: %v.", healthDelay, maybeNewHealthDelay)
			} else {
				healthDelay = newHealthDelay
			}
		} else {
			LogWarn.Warnf("failed to parse value for HEALTH_DELAY: %v. Defaulting to %v.", err, healthDelay)
		}
	}

	go func() {
		healthTimer := time.NewTimer(healthDelay)
		<-healthTimer.C
		isHealthy = true
	}()

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
