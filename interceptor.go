package rfms

import "net/http"

type interceptorTransport struct {
	interceptors []func(http.RoundTripper) http.RoundTripper
	next         http.RoundTripper
}

var _ http.RoundTripper = &interceptorTransport{}

func (t *interceptorTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	rt := t.next
	for _, interceptor := range t.interceptors {
		rt = interceptor(rt)
	}
	return rt.RoundTrip(req)
}
