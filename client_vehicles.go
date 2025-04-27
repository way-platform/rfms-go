package rfms

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ListVehiclesRequest is the request for the [Client.ListVehicles] method.
type ListVehiclesRequest struct {
	LastVIN string
}

// ListVehiclesResponse is the response for the [Client.ListVehicles] method.
type ListVehiclesResponse struct {
	Raw      json.RawMessage `json:"-"`
	Vehicles []Vehicle       `json:"vehicles,omitempty"`
}

// ListVehicles implements the rFMS API method "GET /vehicles".
func (c *Client) ListVehicles(ctx context.Context, request *ListVehiclesRequest) (_ *ListVehiclesResponse, err error) {
	apiMethod, apiURL := "GET", c.baseURL+"/rfms4/vehicles" // TODO: Don't assume /rfms4
	defer func() {
		if err != nil {
			err = fmt.Errorf("%s %s: %w", apiMethod, apiURL, err)
		}
	}()
	req, err := http.NewRequestWithContext(ctx, apiMethod, apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json; rfms=vehicles.v4.0")
	q := req.URL.Query()
	if request != nil && request.LastVIN != "" {
		q.Set("lastVin", request.LastVIN)
	}
	req.URL.RawQuery = q.Encode()
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("response status code: %d", resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}
	return &ListVehiclesResponse{
		Raw: json.RawMessage(data),
	}, nil
}
