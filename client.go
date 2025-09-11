package rfms

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
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

func (c *Client) do(request *retryablehttp.Request) (*http.Response, error) {
	if c.config.debug {
		requestData, err := httputil.DumpRequestOut(request.Request, true)
		if err != nil {
			return nil, err
		}
		var output bytes.Buffer
		output.Grow(2 * len(requestData))
		for line := range bytes.Lines(requestData) {
			_, _ = output.WriteString("> ")
			_, _ = output.Write(line)
		}
		_, _ = os.Stderr.Write(output.Bytes())
	}
	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	if c.config.debug {
		responseData, err := httputil.DumpResponse(response, true)
		if err != nil {
			return nil, err
		}
		var output bytes.Buffer
		output.Grow(2 * len(responseData))
		for line := range bytes.Lines(responseData) {
			_, _ = output.WriteString("< ")
			_, _ = output.Write(line)
		}
		_, _ = os.Stderr.Write(output.Bytes())
	}
	return response, nil
}

func getUserAgent() string {
	userAgent := "rfms-go"
	if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "" {
		userAgent = fmt.Sprintf("rfms-go/%s", info.Main.Version)
	}
	return userAgent
}
