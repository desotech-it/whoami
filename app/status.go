package app

import (
	"encoding/json"
	"os"
)

const (
	STARTING = iota
	UP       = iota
	STOPPING = iota
	DOWN     = iota
	ERRORED  = iota
)

type Health int
type Liveness int

func GetHealth() Health {
	healthValue := os.Getenv("WHOAMI_HEALTH_STATUS")

	if healthValue == "STARTING" {
		return STARTING
	}

	if healthValue == "STOPPING" {
		return STOPPING
	}

	if healthValue == "DOWN" {
		return DOWN
	}

	if healthValue == "ERRORED" {
		return ERRORED
	}

	return UP
}

func GetLiveness() Liveness {
	livenessValue := os.Getenv("WHOAMI_LIVENESS_STATUS")

	if livenessValue == "STARTING" {
		return STARTING
	}

	if livenessValue == "STOPPING" {
		return STOPPING
	}

	if livenessValue == "DOWN" {
		return DOWN
	}

	if livenessValue == "ERRORED" {
		return ERRORED
	}

	return UP
}

func (h Health) String() string {
	if h == STARTING {
		return "STARTING"
	} else if h == UP {
		return "UP"
	} else if h == STOPPING {
		return "STOPPING"
	} else if h == DOWN {
		return "DOWN"
	} else if h == ERRORED {
		return "ERRORED"
	}

	return "UNKNOWN"
}

func (l Liveness) String() string {
	if l == STARTING {
		return "STARTING"
	} else if l == UP {
		return "UP"
	} else if l == STOPPING {
		return "STOPPING"
	} else if l == DOWN {
		return "DOWN"
	} else if l == ERRORED {
		return "ERRORED"
	}

	return "UNKNOWN"
}

func makeJsonResponse(value string) ([]byte, error) {
	type response struct {
		Status string `json:"status"`
	}

	bytes, err := json.Marshal(response{value})

	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (h Health) GetJsonResponse() ([]byte, error) {
	return makeJsonResponse(h.String())
}

func (l Liveness) GetJsonResponse() ([]byte, error) {
	return makeJsonResponse(l.String())
}
