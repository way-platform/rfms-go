package rfms

import "net/http"

// ClientConfig is the configuration for a [Client].
type ClientConfig struct {
	baseURL   string
	transport http.RoundTripper
}

// newClientConfig creates a new default [ClientConfig].
func newClientConfig(baseURL string) ClientConfig {
	return ClientConfig{
		baseURL:   baseURL,
		transport: http.DefaultTransport,
	}
}

// ClientOption is an option that configures a [Client].
type ClientOption func(*ClientConfig)

// WithTransport sets the [http.RoundTripper] HTTP transport for the [Client].
func WithTransport(transport http.RoundTripper) ClientOption {
	return func(cc *ClientConfig) {
		cc.transport = transport
	}
}

// WithBasicAuth authenticates requests using HTTP basic authentication.
func WithBasicAuth(username string, password string) ClientOption {
	return func(cc *ClientConfig) {
		cc.transport = &basicAuthTransport{
			username:  username,
			password:  password,
			transport: cc.transport,
		}
	}
}

// WithScaniaAuth authenticates requests using Scania's HMAC-SHA256 challenge-response mechanism.
func WithScaniaAuth(clientID string, clientSecret string) ClientOption {
	return func(cc *ClientConfig) {
		cc.transport = &tokenAuthenticatorTransport{
			tokenAuthenticator: &scaniaTokenAuthenticator{
				clientID:     clientID,
				clientSecret: clientSecret,
				httpClient:   http.DefaultClient,
			},
			transport: cc.transport,
		}
	}
}
