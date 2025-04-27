package rfms

import "net/http"

// Client is an rFMS API client.
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// NewClient creates a new [Client] with the given base URL and options.
func NewClient(baseURL string, opts ...ClientOption) *Client {
	config := newClientConfig(baseURL)
	for _, opt := range opts {
		opt(&config)
	}
	httpClient := &http.Client{Transport: config.transport}
	if httpClient.Transport == nil {
		httpClient.Transport = http.DefaultTransport
	}
	return &Client{
		baseURL:    baseURL,
		httpClient: httpClient,
	}
}
