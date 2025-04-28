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
	// LastVIN is the last VIN included in the previous response.
	LastVIN string
}

// ListVehiclesResponse is the response for the [Client.ListVehicles] method.
type ListVehiclesResponse struct {
	// Raw response body.
	Raw json.RawMessage `json:"-"`
	// Vehicles in the response.
	Vehicles []json.RawMessage `json:"vehicles,omitempty"`
	// MoreDataAvailable indicates if there is more data available.
	MoreDataAvailable bool `json:"moreDataAvailable,omitempty"`
	// MoreDataAvailableLink is the link to the next page of data.
	MoreDataAvailableLink string `json:"moreDataAvailableLink,omitempty"`
}

// ListVehicles implements the rFMS API method "GET /vehicles".
func (c *Client) ListVehicles(ctx context.Context, request *ListVehiclesRequest) (_ *ListVehiclesResponse, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("list vehicles: %w", err)
		}
	}()
	req, err := c.newRequest(ctx, http.MethodGet, "/vehicles", nil)
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
	if resp.StatusCode != http.StatusOK {
		return nil, newHTTPError(resp)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}
	var responseBody struct {
		VehicleResponse ListVehiclesResponse `json:"vehicleResponse"`
	}
	if err := json.Unmarshal(data, &responseBody); err != nil {
		return nil, fmt.Errorf("unmarshal response body: %w", err)
	}
	var rawVehicles struct {
		VehicleResponse struct {
			Vehicles []json.RawMessage `json:"vehicles"`
		} `json:"vehicleResponse"`
	}
	if err := json.Unmarshal(data, &rawVehicles); err != nil {
		return nil, fmt.Errorf("unmarshal response body: %w", err)
	}
	responseBody.VehicleResponse.Raw = data
	for i, rawVehicle := range rawVehicles.VehicleResponse.Vehicles {
		responseBody.VehicleResponse.Vehicles[i].Raw = rawVehicle
	}
	return &responseBody.VehicleResponse, nil
}
