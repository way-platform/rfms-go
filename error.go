package rfms

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

// Error is an error returned by the rFMS API.
type Error struct {
	// Method is the HTTP method used to make the request.
	Method string
	// URL is the URL of the request.
	URL string
	// Status is the HTTP status.
	Status string
	// StatusCode is the HTTP status code.
	StatusCode int
	// RateLimitReset is the duration until the rate limit is reset.
	RateLimitReset time.Duration
	// Identifier of the rFMS error.
	Identifier string
	// Description of the rFMS error.
	Description string
	// ErrorURI provides more information about the error.
	ErrorURI string
}

// ErrorObject Optional responses for error codes, detailing the error if needed
type ErrorObject struct {
	// Error An identifier for this error
	Error *string `json:"error,omitempty"`

	// ErrorDescription A description of the error
	ErrorDescription *string `json:"error_description,omitempty"`

	// ErrorUri A URI providing more information
	ErrorUri *string `json:"error_uri,omitempty"`
}

func newHTTPError(resp *http.Response) *Error {
	result := &Error{
		Method:     resp.Request.Method,
		URL:        resp.Request.URL.String(),
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
	}
	if data, err := io.ReadAll(resp.Body); err == nil {
		var body struct {
			Error            string `json:"error"`
			ErrorDescription string `json:"error_description"`
			ErrorURI         string `json:"error_uri"`
		}
		if err := json.Unmarshal(data, &body); err == nil {
			result.Identifier = body.Error
			result.Description = body.ErrorDescription
			result.ErrorURI = body.ErrorURI
		}
	}
	if retryAfterHeader := resp.Header.Get("Retry-After"); retryAfterHeader != "" {
		if retryAfter, err := strconv.Atoi(retryAfterHeader); err == nil {
			result.RateLimitReset = time.Duration(retryAfter) * time.Second
		}
	}
	return result
}

// Error implements the error interface.
func (e *Error) Error() string {
	if e.RateLimitReset > 0 {
		return fmt.Sprintf("%s %s - %s: rate limit reset in %s", e.Method, e.URL, e.Status, e.RateLimitReset)
	}
	return fmt.Sprintf("%s %s: %s: %s - %s %s", e.Method, e.URL, e.Status, e.Identifier, e.Description, e.ErrorURI)
}
