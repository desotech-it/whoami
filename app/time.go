package app

import (
	"strconv"
	"time"
)

func ParseDuration(s string) (time.Duration, error) {
	secs, err := strconv.Atoi(s)
	if err == nil {
		return time.Duration(secs) * time.Second, nil
	}
	return time.ParseDuration(s)
}
