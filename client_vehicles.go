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

// Vehicles implements the rFMS API method "GET /vehicles".
func (c *Client) Vehicles(
	ctx context.Context,
	request *rfmsv5.VehiclesRequest,
) (_ *rfmsv5.VehiclesResponse, err error) {
	switch c.config.apiVersion {
	case V2_1:
		return c.vehiclesV2(ctx, request)
	case V4:
		return c.vehiclesV4(ctx, request)
	default:
		return nil, fmt.Errorf("unsupported API version")
	}
}

func (c *Client) vehiclesV2(
	ctx context.Context,
	request *rfmsv5.VehiclesRequest,
) (_ *rfmsv5.VehiclesResponse, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("rFMS v2 vehicles: %w", err)
		}
	}()
	query := url.Values{}
	if request.GetLastVin() != "" {
		query.Set("lastVin", request.GetLastVin())
	}
	path := "/vehicles"
	if len(query) > 0 {
		path += "?" + query.Encode()
	}
	fullURL := c.config.baseURL + path
	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("http request: %w", err)
	}
	httpRequest.Header.Set("User-Agent", getUserAgent())
	httpRequest.Header.Set("Accept", "application/vnd.fmsstandard.com.Vehicles.v2.1+json")
	httpResponse, err := c.httpClient(c.config).Do(httpRequest)
	if err != nil {
		return nil, fmt.Errorf("http request: %w", err)
	}
	defer func() {
		_ = httpResponse.Body.Close()
	}()
	if httpResponse.StatusCode != http.StatusOK {
		return nil, newHTTPError(httpResponse)
	}
	data, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}
	var response rfmsv2oapi.Vehicles
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, fmt.Errorf("unmarshal v2 response body: %w", err)
	}
	resp := &rfmsv5.VehiclesResponse{}
	vehicles := make([]*rfmsv5.Vehicle, 0, len(response.Vehicle))
	for _, vehicle := range response.Vehicle {
		vehicles = append(vehicles, convertv2.Vehicle(&vehicle))
	}
	resp.SetVehicles(vehicles)
	if response.MoreDataAvailable != nil {
		resp.SetMoreDataAvailable(*response.MoreDataAvailable)
	}
	return resp, nil
}

func (c *Client) vehiclesV4(
	ctx context.Context,
	request *rfmsv5.VehiclesRequest,
) (_ *rfmsv5.VehiclesResponse, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("rFMS v4 vehicles: %w", err)
		}
	}()
	query := url.Values{}
	if request.GetLastVin() != "" {
		query.Set("lastVin", request.GetLastVin())
	}
	path := "/vehicles"
	if len(query) > 0 {
		path += "?" + query.Encode()
	}
	fullURL := c.config.baseURL + path
	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("http request: %w", err)
	}
	httpRequest.Header.Set("User-Agent", getUserAgent())
	httpRequest.Header.Set("Accept", "application/json; rfms=vehicles.v4.0")
	httpResponse, err := c.httpClient(c.config).Do(httpRequest)
	if err != nil {
		return nil, fmt.Errorf("http request: %w", err)
	}
	defer func() {
		_ = httpResponse.Body.Close()
	}()
	if httpResponse.StatusCode != http.StatusOK {
		return nil, newHTTPError(httpResponse)
	}
	data, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}
	slog.Debug("vehicles v4 response", "body", json.RawMessage(data))
	var response rfmsv4oapi.VehicleResponseObject
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, fmt.Errorf("unmarshal v4 response body: %w", err)
	}
	resp := &rfmsv5.VehiclesResponse{}
	vehicles := make([]*rfmsv5.Vehicle, 0, len(response.VehicleResponse.Vehicles))
	for _, vehicle := range response.VehicleResponse.Vehicles {
		vehicles = append(vehicles, convertv4.Vehicle(&vehicle))
	}
	resp.SetVehicles(vehicles)
	if response.MoreDataAvailable != nil {
		resp.SetMoreDataAvailable(*response.MoreDataAvailable)
	}
	return resp, nil
}
