package app

import (
	"encoding/json"
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
	return UP
}

func GetLiveness() Liveness {
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
