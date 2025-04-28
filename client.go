package rfms

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"runtime/debug"
)

// Client is an rFMS API client.
type Client struct {
	baseURL    string
	apiVersion Version
	httpClient *http.Client
}

// NewClient creates a new [Client] with the given base URL and options.
func NewClient(baseURL string, opts ...ClientOption) *Client {
	config := newClientConfig()
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

func (c *Client) newRequest(ctx context.Context, method, path string, body io.Reader) (_ *http.Request, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("new request: %w", err)
		}
	}()
	request, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, body)
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
