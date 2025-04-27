package rfms

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// HTTPError is an HTTP error returned by the rFMS API.
type HTTPError struct {
	// Method is the HTTP method used to make the request.
	Method string
	// URL is the URL of the request.
	URL string
	// Status is the HTTP status.
	Status string `json:"status"`
	// StatusCode is the HTTP status code.
	StatusCode int `json:"status_code"`
	// RateLimitReset is the duration until the rate limit is reset.
	RateLimitReset time.Duration
}

func newHTTPError(resp *http.Response) *HTTPError {
	result := &HTTPError{
		Method:     resp.Request.Method,
		URL:        resp.Request.URL.String(),
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
	}
	if retryAfterHeader := resp.Header.Get("Retry-After"); retryAfterHeader != "" {
		if retryAfter, err := strconv.Atoi(retryAfterHeader); err == nil {
			result.RateLimitReset = time.Duration(retryAfter) * time.Second
		}
	}
	return result
}

// Error implements the error interface.
func (e *HTTPError) Error() string {
	if e.RateLimitReset > 0 {
		return fmt.Sprintf("%s %s - %s: rate limit reset in %s", e.Method, e.URL, e.Status, e.RateLimitReset)
	}
	return fmt.Sprintf("%s %s - %s", e.Method, e.URL, e.Status)
}
