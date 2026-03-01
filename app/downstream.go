package app

import (
	"io"
	"net/http"
	"time"
)

// DownstreamResult holds the result of a downstream HTTP call.
type DownstreamResult struct {
	URL        string
	StatusCode int
	Headers    http.Header
	Body       string
	Duration   time.Duration
	Error      string
}

// CallDownstream makes an HTTP GET to the given URL and returns the result.
// Timeout is 5 seconds and response body is limited to 10KB.
func CallDownstream(targetURL string) DownstreamResult {
	client := &http.Client{Timeout: 5 * time.Second}
	start := time.Now()

	resp, err := client.Get(targetURL)
	duration := time.Since(start)

	if err != nil {
		return DownstreamResult{
			URL:      targetURL,
			Error:    err.Error(),
			Duration: duration,
		}
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(io.LimitReader(resp.Body, 10240))

	return DownstreamResult{
		URL:        targetURL,
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       string(bodyBytes),
		Duration:   duration,
	}
}
