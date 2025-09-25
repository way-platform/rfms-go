package rfms

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/way-platform/rfms-go/internal/convert/convertv2"
	"github.com/way-platform/rfms-go/internal/convert/convertv4"
	"github.com/way-platform/rfms-go/internal/openapi/rfmsv2oapi"
	"github.com/way-platform/rfms-go/internal/openapi/rfmsv4oapi"
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/connect/rfms/v5"
)

// VehiclesRequest is the request for the [Client.Vehicles] method.
type VehiclesRequest struct {
	// LastVIN is the last VIN included in the previous response.
	LastVIN string `json:"lastVin"`
}

// VehiclesResponse is the response for the [Client.Vehicles] method.
type VehiclesResponse struct {
	// Vehicles in the response.
	Vehicles []*rfmsv5.Vehicle `json:"vehicles"`
	// MoreDataAvailable indicates if there is more data available.
	MoreDataAvailable bool `json:"moreDataAvailable"`
}

// Vehicles implements the rFMS API method "GET /vehicles".
func (c *Client) Vehicles(ctx context.Context, request VehiclesRequest, opts ...ClientOption) (_ VehiclesResponse, err error) {
	cfg := c.config.with(opts...)
	switch cfg.apiVersion {
	case V2_1:
		return c.vehiclesV2(ctx, request, cfg)
	case V4:
		return c.vehiclesV4(ctx, request, cfg)
	default:
		return VehiclesResponse{}, fmt.Errorf("unsupported API version")
	}
}

func (c *Client) vehiclesV2(ctx context.Context, request VehiclesRequest, cfg ClientConfig) (_ VehiclesResponse, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("rFMS v2 vehicles: %w", err)
		}
	}()
	// Build query parameters
	query := url.Values{}
	if request.LastVIN != "" {
		query.Set("lastVin", request.LastVIN)
	}

	// Build path with query parameters
	path := "/vehicles"
	if len(query) > 0 {
		path += "?" + query.Encode()
	}
	// Apply per-request configuration overrides
	fullURL := cfg.baseURL + path
	// Create the request
	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return VehiclesResponse{}, fmt.Errorf("http request: %w", err)
	}
	// Set headers
	httpRequest.Header.Set("User-Agent", getUserAgent())
	httpRequest.Header.Set("Accept", "application/vnd.fmsstandard.com.Vehicles.v2.1+json")
	// Create HTTP client and make request
	client := c.httpClient(cfg)
	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		return VehiclesResponse{}, fmt.Errorf("http request: %w", err)
	}
	defer httpResponse.Body.Close()
	if httpResponse.StatusCode != http.StatusOK {
		return VehiclesResponse{}, newHTTPError(httpResponse)
	}
	data, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return VehiclesResponse{}, fmt.Errorf("read response body: %w", err)
	}
	var response rfmsv2oapi.Vehicles
	if err := json.Unmarshal(data, &response); err != nil {
		return VehiclesResponse{}, fmt.Errorf("unmarshal v2 response body: %w", err)
	}
	result := VehiclesResponse{
		MoreDataAvailable: response.MoreDataAvailable != nil && *response.MoreDataAvailable,
		Vehicles:          make([]*rfmsv5.Vehicle, 0, len(response.Vehicle)),
	}
	for _, vehicle := range response.Vehicle {
		result.Vehicles = append(result.Vehicles, convertv2.Vehicle(&vehicle))
	}
	return result, nil
}

func (c *Client) vehiclesV4(ctx context.Context, request VehiclesRequest, cfg ClientConfig) (_ VehiclesResponse, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("rFMS v4 vehicles: %w", err)
		}
	}()
	// Build query parameters
	query := url.Values{}
	if request.LastVIN != "" {
		query.Set("lastVin", request.LastVIN)
	}
	// Build path with query parameters
	path := "/vehicles"
	if len(query) > 0 {
		path += "?" + query.Encode()
	}
	// Apply per-request configuration overrides
	fullURL := cfg.baseURL + path
	// Create the request
	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return VehiclesResponse{}, fmt.Errorf("http request: %w", err)
	}
	// Set headers
	httpRequest.Header.Set("User-Agent", getUserAgent())
	httpRequest.Header.Set("Accept", "application/json; rfms=vehicles.v4.0")
	// Create HTTP client and make request
	client := c.httpClient(cfg)
	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		return VehiclesResponse{}, fmt.Errorf("http request: %w", err)
	}
	defer httpResponse.Body.Close()
	if httpResponse.StatusCode != http.StatusOK {
		return VehiclesResponse{}, newHTTPError(httpResponse)
	}
	data, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return VehiclesResponse{}, fmt.Errorf("read response body: %w", err)
	}
	slog.Debug("vehicles v4 response", "body", json.RawMessage(data))
	var response rfmsv4oapi.VehicleResponseObject
	if err := json.Unmarshal(data, &response); err != nil {
		return VehiclesResponse{}, fmt.Errorf("unmarshal v4 response body: %w", err)
	}
	result := VehiclesResponse{
		MoreDataAvailable: response.MoreDataAvailable != nil && *response.MoreDataAvailable,
		Vehicles:          make([]*rfmsv5.Vehicle, 0, len(response.VehicleResponse.Vehicles)),
	}
	for _, vehicle := range response.VehicleResponse.Vehicles {
		result.Vehicles = append(result.Vehicles, convertv4.Vehicle(&vehicle))
	}
	return result, nil
}
