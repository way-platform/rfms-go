package rfms

import (
	"context"
	"fmt"
	"io"
	"runtime/debug"

	"github.com/hashicorp/go-retryablehttp"
)

// Client is an rFMS API client.
type Client struct {
	config     ClientConfig
	httpClient *retryablehttp.Client
}

// NewClient creates a new [Client] with the given base URL and options.
func NewClient(opts ...ClientOption) *Client {
	config := newClientConfig()
	for _, opt := range opts {
		opt(&config)
	}
	httpClient := retryablehttp.NewClient()
	if config.transport != nil {
		httpClient.HTTPClient.Transport = config.transport
	}
	httpClient.RetryMax = config.retryCount
	// TODO: Add custom retry logic for non-standard rFMS retry headers.
	if config.logger != nil {
		httpClient.Logger = config.logger
	}
	return &Client{
		config:     config,
		httpClient: httpClient,
	}
}

func (c *Client) newRequest(ctx context.Context, method, path string, body io.Reader) (_ *retryablehttp.Request, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("new request: %w", err)
		}
	}()
	request, err := retryablehttp.NewRequestWithContext(ctx, method, c.config.baseURL+path, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("User-Agent", getUserAgent())
	return request, nil
}

func getUserAgent() string {
	userAgent := "rfms-go"
	if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "" {
		userAgent = fmt.Sprintf("rfms-go/%s", info.Main.Version)
	}
	return userAgent
}
