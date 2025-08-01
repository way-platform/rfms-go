package rfms

import (
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
)

// ClientConfig is the configuration for a [Client].
type ClientConfig struct {
	baseURL    string
	apiVersion Version
	transport  http.RoundTripper
	retryCount int
	logger     Logger
}

// newClientConfig creates a new default [ClientConfig].
func newClientConfig() ClientConfig {
	return ClientConfig{
		baseURL:    ScaniaBaseURL,
		apiVersion: V4,
		transport:  http.DefaultTransport,
	}
}

// ClientOption is an option that configures a [Client].
type ClientOption func(*ClientConfig)

// WithBaseURL sets the API base URL for the [Client].
func WithBaseURL(baseURL string) ClientOption {
	return func(cc *ClientConfig) {
		cc.baseURL = baseURL
	}
}

// WithTransport sets the [http.RoundTripper] HTTP transport for the [Client].
func WithTransport(transport http.RoundTripper) ClientOption {
	return func(cc *ClientConfig) {
		cc.transport = transport
	}
}

// WithVersion sets the rFMS API version for the [Client].
func WithVersion(apiVersion Version) ClientOption {
	return func(cc *ClientConfig) {
		cc.apiVersion = apiVersion
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
				baseURL:      ScaniaAuthBaseURL,
				clientID:     clientID,
				clientSecret: clientSecret,
				httpClient:   retryablehttp.NewClient(),
			},
			transport: cc.transport,
		}
	}
}

// WithReuseTokenAuth authenticates requests by re-using existing [TokenCredentials].
func WithReuseTokenAuth(credentials TokenCredentials) ClientOption {
	return func(cc *ClientConfig) {
		cc.transport = &reuseTokenCredentialsTransport{
			transport:   cc.transport,
			credentials: credentials,
		}
	}
}

// WithScania configures the [Client] to use the Scania rFMS v4 API.
func WithScania(clientID string, clientSecret string) ClientOption {
	return func(cc *ClientConfig) {
		WithBaseURL(ScaniaBaseURL)(cc)
		WithScaniaAuth(clientID, clientSecret)(cc)
		WithVersion(V4)(cc)
	}
}

// WithVolvoTrucks configures the [Client] to use the Volvo Trucks rFMS v2.1 API.
func WithVolvoTrucks(username string, password string) ClientOption {
	return func(cc *ClientConfig) {
		WithBaseURL(VolvoTrucksBaseURL)(cc)
		WithBasicAuth(username, password)(cc)
		WithVersion(V2_1)(cc)
	}
}

// WithRetryCount sets the maximum number of times to retry a request.
func WithRetryCount(retryCount int) ClientOption {
	return func(cc *ClientConfig) {
		cc.retryCount = retryCount
	}
}

// Logger is a leveled logger interface.
type Logger interface {
	Debug(msg string, keysAndValues ...any)
	Info(msg string, keysAndValues ...any)
	Warn(msg string, keysAndValues ...any)
	Error(msg string, keysAndValues ...any)
}

// WithLogger sets the [Logger] for the [Client].
func WithLogger(logger Logger) ClientOption {
	return func(cc *ClientConfig) {
		cc.logger = logger
	}
}
