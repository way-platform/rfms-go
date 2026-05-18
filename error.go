package rfms

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"connectrpc.com/connect"
	"github.com/way-platform/rfms-go/internal/openapi/rfmsv4oapi"
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

func newHTTPError(resp *http.Response) error {
	result := &Error{
		Method:     resp.Request.Method,
		URL:        resp.Request.URL.String(),
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
	}
	if data, err := io.ReadAll(resp.Body); err == nil {
		var body rfmsv4oapi.ErrorObject
		if err := json.Unmarshal(data, &body); err == nil {
			if body.Error != nil {
				result.Identifier = *body.Error
			}
			if body.ErrorDescription != nil {
				result.Description = *body.ErrorDescription
			}
			if body.ErrorURI != nil {
				result.ErrorURI = *body.ErrorURI
			}
		} else {
			result.Description = string(data)
		}
	}
	if retryAfterHeader := resp.Header.Get("Retry-After"); retryAfterHeader != "" {
		if retryAfter, err := strconv.Atoi(retryAfterHeader); err == nil {
			result.RateLimitReset = time.Duration(retryAfter) * time.Second
		}
	}
	return result.connectError()
}

// Error implements the error interface.
func (e *Error) Error() string {
	if e.RateLimitReset > 0 {
		return fmt.Sprintf(
			"%s %s - %s: rate limit reset in %s",
			e.Method,
			e.URL,
			e.Status,
			e.RateLimitReset,
		)
	}
	return fmt.Sprintf(
		"%s %s: %s: %s - %s %s",
		e.Method,
		e.URL,
		e.Status,
		e.Identifier,
		e.Description,
		e.ErrorURI,
	)
}

func (e *Error) connectError() error {
	return connect.NewError(httpStatusToConnectCode(e.StatusCode), e)
}

func httpStatusToConnectCode(statusCode int) connect.Code {
	switch statusCode {
	case http.StatusBadRequest:
		return connect.CodeInvalidArgument
	case http.StatusUnauthorized:
		return connect.CodeUnauthenticated
	case http.StatusForbidden:
		return connect.CodePermissionDenied
	case http.StatusNotFound:
		return connect.CodeNotFound
	case http.StatusConflict:
		return connect.CodeAlreadyExists
	case http.StatusTooManyRequests:
		return connect.CodeResourceExhausted
	case http.StatusNotImplemented:
		return connect.CodeUnimplemented
	case http.StatusServiceUnavailable:
		return connect.CodeUnavailable
	case http.StatusGatewayTimeout:
		return connect.CodeDeadlineExceeded
	case http.StatusInternalServerError:
		return connect.CodeInternal
	default:
		return connect.CodeUnknown
	}
}
