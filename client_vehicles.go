package rfms

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/way-platform/rfms-go/v4/rfmsv4"
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
	Vehicles []rfmsv4.Vehicle `json:"vehicles,omitzero"`
	// MoreDataAvailable indicates if there is more data available.
	MoreDataAvailable bool `json:"moreDataAvailable,omitempty"`
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
	switch c.apiVersion {
	case Version4:
		req.Header.Set("Accept", "application/json; rfms=vehicles.v4.0")
	case Version21:
		req.Header.Set("Accept", "application/vnd.fmsstandard.com.Vehicles.v2.1+json; UTF-8")
	default:
		return nil, fmt.Errorf("unsupported API version: %s", c.apiVersion)
	}
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
	var responseBody rfmsv4.VehiclesResponse
	log.Println(string(data))
	if err := json.Unmarshal(data, &responseBody); err != nil {
		return nil, fmt.Errorf("unmarshal response body: %w", err)
	}
	// var rawVehicles struct {
	// 	VehicleResponse struct {
	// 		Vehicles []json.RawMessage `json:"vehicles"`
	// 	} `json:"vehicleResponse"`
	// }
	// if err := json.Unmarshal(data, &rawVehicles); err != nil {
	// 	return nil, fmt.Errorf("unmarshal response body: %w", err)
	// }
	// responseBody.VehicleResponse.Raw = data
	// for i, rawVehicle := range rawVehicles.VehicleResponse.Vehicles {
	// 	responseBody.VehicleResponse.Vehicles[i].Raw = rawVehicle
	// }
	return &ListVehiclesResponse{
		Vehicles:          *responseBody.VehicleResponse.Vehicles,
		MoreDataAvailable: responseBody.MoreDataAvailable,
	}, nil
}
