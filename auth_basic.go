package rfms

import "net/http"

// basicAuthTransport is an HTTP transport that authenticates requests using HTTP basic authentication.
type basicAuthTransport struct {
	username  string
	password  string
	transport http.RoundTripper
}

var _ http.RoundTripper = &basicAuthTransport{}

// RoundTrip implements the [http.RoundTripper] interface.
func (t *basicAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(t.username, t.password)
	return t.transport.RoundTrip(req)
}
