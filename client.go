package rfms

// Client is an rFMS API client.
type Client struct {
	config ClientConfig
}

// NewClient creates a new [Client] with the given base URL and options.
func NewClient(baseURL string, opts ...ClientOption) *Client {
	config := newClientConfig(baseURL)
	for _, opt := range opts {
		opt(&config)
	}
	return &Client{
		config: config,
	}
}
