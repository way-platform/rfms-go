package rfms

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/way-platform/rfms-go/v4/api/compat/rfmsv2tov4"
	"github.com/way-platform/rfms-go/v4/api/rfmsv2"
	"github.com/way-platform/rfms-go/v4/api/rfmsv4"
)

// VehiclesRequest is the request for the [Client.Vehicles] method.
type VehiclesRequest struct {
	// LastVIN is the last VIN included in the previous response.
	LastVIN string
}

// VehiclesResponse is the response for the [Client.Vehicles] method.
type VehiclesResponse struct {
	// Vehicles in the response.
	Vehicles []rfmsv4.Vehicle `json:"vehicles"`
	// MoreDataAvailable indicates if there is more data available.
	MoreDataAvailable bool `json:"moreDataAvailable"`
}

// Vehicles implements the rFMS API method "GET /vehicles".
func (c *Client) Vehicles(ctx context.Context, request *VehiclesRequest) (_ *VehiclesResponse, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("vehicles: %w", err)
		}
	}()
	httpRequest, err := c.newRequest(ctx, http.MethodGet, "/vehicles", nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	switch c.config.apiVersion {
	case V4:
		httpRequest.Header.Set("Accept", "application/json; rfms=vehicles.v4.0")
	case V2_1:
		httpRequest.Header.Set("Accept", "application/vnd.fmsstandard.com.Vehicles.v2.1+json; UTF-8")
	default:
		return nil, fmt.Errorf("unsupported API version: %s", c.config.apiVersion)
	}
	q := httpRequest.URL.Query()
	if request != nil && request.LastVIN != "" {
		q.Set("lastVin", request.LastVIN)
	}
	httpRequest.URL.RawQuery = q.Encode()
	httpResponse, err := c.httpClient.Do(httpRequest)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}
	defer httpResponse.Body.Close()
	if httpResponse.StatusCode != http.StatusOK {
		return nil, newHTTPError(httpResponse)
	}
	data, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}
	var v4Response rfmsv4.VehiclesResponse
	switch c.config.apiVersion {
	case V4:
		if err := json.Unmarshal(data, &v4Response); err != nil {
			return nil, fmt.Errorf("unmarshal v4 response body: %w", err)
		}
	case V2_1:
		var v2Response rfmsv2.VehiclesResponse
		if err := json.Unmarshal(data, &v2Response); err != nil {
			return nil, fmt.Errorf("unmarshal v2 response body: %w", err)
		}
		v4Response = *rfmsv2tov4.ConvertVehiclesResponse(&v2Response)
	}
	return &VehiclesResponse{
		Vehicles:          v4Response.VehicleResponse.Vehicles,
		MoreDataAvailable: v4Response.MoreDataAvailable,
	}, nil
}
