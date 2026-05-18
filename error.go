package rfms

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"connectrpc.com/connect"
	"github.com/way-platform/rfms-go/internal/openapi/rfmsv4oapi"
)

func newHTTPError(resp *http.Response) error {
	var msg string
	if data, err := io.ReadAll(resp.Body); err == nil {
		var body rfmsv4oapi.ErrorObject
		if err := json.Unmarshal(data, &body); err == nil {
			parts := fmt.Sprintf("http %d", resp.StatusCode)
			if body.Error != nil {
				parts += ": " + *body.Error
			}
			if body.ErrorDescription != nil {
				parts += " - " + *body.ErrorDescription
			}
			msg = parts
		} else if len(data) > 0 {
			msg = fmt.Sprintf("http %d: %s", resp.StatusCode, data)
		} else {
			msg = fmt.Sprintf("http %d", resp.StatusCode)
		}
	} else {
		msg = fmt.Sprintf("http %d", resp.StatusCode)
	}
	return connect.NewError(httpStatusToConnectCode(resp.StatusCode), errors.New(msg))
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
