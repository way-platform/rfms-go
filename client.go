package rfms

import (
	"context"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"golang.org/x/oauth2"
)

// Client is an rFMS API client.
type Client struct {
	config ClientConfig
}

// ClientConfig is the configuration for a [Client].
type ClientConfig struct {
	baseURL      string
	apiVersion   Version
	retryCount   int
	debug        bool
	timeout      time.Duration
	tokenSource  oauth2.TokenSource
	username     string
	password     string
	interceptors []func(http.RoundTripper) http.RoundTripper
}

// newClientConfig creates a new default [ClientConfig].
func newClientConfig() ClientConfig {
	return ClientConfig{
		timeout:    30 * time.Second,
		retryCount: 3,
	}
}

// with returns a new ClientConfig with the given options applied.
// This enables per-request configuration overrides.
func (cc ClientConfig) with(opts ...ClientOption) ClientConfig {
	for _, opt := range opts {
		opt(&cc)
	}
	return cc
}

// ClientOption is an option that configures a [Client].
type ClientOption func(*ClientConfig)

// NewClient creates a new [Client] with the given base URL and options.
func NewClient(opts ...ClientOption) (*Client, error) {
	config := newClientConfig()
	for _, opt := range opts {
		opt(&config)
	}
	if config.apiVersion == "" {
		return nil, fmt.Errorf("api version is required")
	}
	if config.baseURL == "" {
		return nil, fmt.Errorf("base URL is required")
	}
	return &Client{
		config: config,
	}, nil
}

// httpClient creates a new HTTP client with the given configuration.
func (c *Client) httpClient(cfg ClientConfig) *http.Client {
	transport := http.DefaultTransport
	// Add debug transport if debug is enabled.
	if cfg.debug {
		transport = &debugTransport{
			next: transport,
		}
	}
	// Add basic auth transport if username/password are configured
	if cfg.username != "" && cfg.password != "" {
		transport = &basicAuthTransport{
			username:  cfg.username,
			password:  cfg.password,
			transport: transport,
		}
	}
	// Add OAuth2 transport if token source is configured.
	if cfg.tokenSource != nil {
		transport = &oauth2.Transport{
			Source: cfg.tokenSource,
			Base:   transport,
		}
	}
	// Add middleware transport if middlewares are configured.
	if len(cfg.interceptors) > 0 {
		transport = &interceptorTransport{
			interceptors: cfg.interceptors,
			next:         transport,
		}
	}
	// Add retry transport if retry count > 0.
	if cfg.retryCount > 0 {
		transport = &retryTransport{
			maxRetries: cfg.retryCount,
			next:       transport,
		}
	}
	return &http.Client{
		Timeout:   cfg.timeout,
		Transport: transport,
	}
}

// WithDebug sets the debug flag for the [Client].
func WithDebug(debug bool) ClientOption {
	return func(cc *ClientConfig) {
		cc.debug = debug
	}
}

// WithBaseURL sets the API base URL for the [Client].
func WithBaseURL(baseURL string) ClientOption {
	return func(cc *ClientConfig) {
		cc.baseURL = baseURL
	}
}

// WithVersion sets the rFMS API version for the [Client].
func WithVersion(apiVersion Version) ClientOption {
	return func(cc *ClientConfig) {
		cc.apiVersion = apiVersion
	}
}

// WithRetryCount sets the maximum number of times to retry a request.
func WithRetryCount(retryCount int) ClientOption {
	return func(cc *ClientConfig) {
		cc.retryCount = retryCount
	}
}

// WithTimeout sets the timeout for HTTP requests.
func WithTimeout(timeout time.Duration) ClientOption {
	return func(cc *ClientConfig) {
		cc.timeout = timeout
	}
}

// WithTokenSource sets the OAuth2 token source for authentication.
func WithTokenSource(tokenSource oauth2.TokenSource) ClientOption {
	return func(cc *ClientConfig) {
		cc.tokenSource = tokenSource
	}
}

// WithBasicAuth sets the username and password for HTTP basic authentication.
func WithBasicAuth(username, password string) ClientOption {
	return func(cc *ClientConfig) {
		cc.username = username
		cc.password = password
	}
}

// Convenience Configuration Functions

// WithScania configures the [Client] to use the Scania rFMS v4 API.
func WithScania(clientID string, clientSecret string) ClientOption {
	return func(cc *ClientConfig) {
		WithBaseURL(ScaniaBaseURL)(cc)
		WithTokenSource(ScaniaAuthConfig{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Debug:        cc.debug,
			MaxRetries:   cc.retryCount,
			Timeout:      cc.timeout,
		}.TokenSource(context.Background()))
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

// WithInterceptor adds a request interceptor for the [Client].
func WithInterceptor(interceptor func(http.RoundTripper) http.RoundTripper) ClientOption {
	return func(cc *ClientConfig) {
		cc.interceptors = append(cc.interceptors, interceptor)
	}
}

func getUserAgent() string {
	const userAgent = "WayPlatformRFMSGo"
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, dep := range info.Deps {
			if dep.Path == "github.com/way-platform/rfms-go" {
				if dep.Version != "" && dep.Version != "v0.0.0-00010101000000-000000000000" {
					return userAgent + "/" + dep.Version
				}
				break
			}
		}
	}
	return userAgent
}
