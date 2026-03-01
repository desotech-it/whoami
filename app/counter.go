package app

import "sync/atomic"

var requestCounter atomic.Int64

// IncrementRequestCounter atomically increments the global request counter
// and returns the new value.
func IncrementRequestCounter() int64 {
	return requestCounter.Add(1)
}

// GetRequestCount returns the current value of the global request counter.
func GetRequestCount() int64 {
	return requestCounter.Load()
}
