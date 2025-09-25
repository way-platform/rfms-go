package rfms

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/way-platform/rfms-go/internal/convert/convertv2"
	"github.com/way-platform/rfms-go/internal/convert/convertv4"
	"github.com/way-platform/rfms-go/internal/openapi/rfmsv2oapi"
	"github.com/way-platform/rfms-go/internal/openapi/rfmsv4oapi"
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/connect/rfms/v5"
)

// VehicleStatusesRequest is the request for the [Client.VehicleStatuses] method.
type VehicleStatusesRequest struct {
	// LastVIN is the last VIN included in the previous response.
	LastVIN string
	// DateType indicates whether the start/stop times are compared to created or received time.
	DateType string
	// StartTime to filter statuses (only statuses after this time).
	StartTime time.Time
	// StopTime to filter statuses (only statuses before this time).
	StopTime time.Time
	// VIN to filter statuses for a specific vehicle.
	VIN string
	// ContentFilter filters statuses by content type (ACCUMULATED, SNAPSHOT, UPTIME).
	ContentFilter []string
	// TriggerFilter filters statuses by trigger type.
	TriggerFilter []string
	// LatestOnly returns only the latest status for each vehicle.
	LatestOnly bool
}

// VehicleStatusesResponse is the response for the [Client.VehicleStatuses] method.
type VehicleStatusesResponse struct {
	// VehicleStatuses in the response.
	VehicleStatuses []*rfmsv5.VehicleStatus `json:"vehicleStatuses"`
	// MoreDataAvailable indicates if there is more data available.
	MoreDataAvailable bool `json:"moreDataAvailable"`
	// RequestServerDateTime is the server time when the request was received.
	RequestServerDateTime time.Time `json:"requestServerDateTime,omitzero"`
}

func (c *Client) VehicleStatuses(ctx context.Context, request VehicleStatusesRequest, opts ...ClientOption) (_ VehicleStatusesResponse, err error) {
	cfg := c.config.with(opts...)
	switch cfg.apiVersion {
	case V2_1:
		return c.vehicleStatusesV2(ctx, request, cfg)
	case V4:
		return c.vehicleStatusesV4(ctx, request, cfg)
	default:
		return VehicleStatusesResponse{}, fmt.Errorf("unsupported API version")
	}
}

func (c *Client) vehicleStatusesV2(ctx context.Context, request VehicleStatusesRequest, cfg ClientConfig) (_ VehicleStatusesResponse, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("rFMS v2 vehicle statuses: %w", err)
		}
	}()
	// Build query parameters
	query := url.Values{}
	if request.LastVIN != "" {
		query.Set("lastVin", request.LastVIN)
	}
	if request.DateType != "" {
		query.Set("datetype", request.DateType)
	}
	if !request.StartTime.IsZero() {
		query.Set("starttime", rfmsv4oapi.Time(request.StartTime).String())
	}
	if !request.StopTime.IsZero() {
		query.Set("stoptime", rfmsv4oapi.Time(request.StopTime).String())
	}
	if request.VIN != "" {
		query.Set("vin", request.VIN)
	}
	if len(request.ContentFilter) > 0 {
		query.Set("contentFilter", strings.Join(request.ContentFilter, ","))
	}
	if len(request.TriggerFilter) > 0 {
		query.Set("triggerFilter", strings.Join(request.TriggerFilter, ","))
	}
	if request.LatestOnly {
		query.Set("latestOnly", "true")
	}
	// Build path with query parameters
	path := "/vehiclestatuses"
	if len(query) > 0 {
		path += "?" + query.Encode()
	}
	// Apply per-request configuration overrides
	fullURL := cfg.baseURL + path
	// Create the request
	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return VehicleStatusesResponse{}, fmt.Errorf("http request: %w", err)
	}
	// Set headers
	httpRequest.Header.Set("User-Agent", getUserAgent())
	httpRequest.Header.Set("Accept", "application/vnd.fmsstandard.com.Vehiclestatuses.v2.1+json")
	// Create HTTP client and make request
	client := c.httpClient(cfg)
	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		return VehicleStatusesResponse{}, fmt.Errorf("http request: %w", err)
	}
	defer httpResponse.Body.Close()
	if httpResponse.StatusCode != http.StatusOK {
		return VehicleStatusesResponse{}, newHTTPError(httpResponse)
	}
	data, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return VehicleStatusesResponse{}, fmt.Errorf("read response body: %w", err)
	}
	var response rfmsv2oapi.VehicleStatuses
	if err := json.Unmarshal(data, &response); err != nil {
		return VehicleStatusesResponse{}, fmt.Errorf("unmarshal v4 response body: %w", err)
	}
	var result VehicleStatusesResponse
	result.MoreDataAvailable = response.MoreDataAvailable != nil && *response.MoreDataAvailable
	for _, vehicleStatus := range response.VehicleStatus {
		result.VehicleStatuses = append(result.VehicleStatuses, convertv2.VehicleStatus(&vehicleStatus))
	}
	if response.RequestServerDateTime != nil {
		result.RequestServerDateTime = *response.RequestServerDateTime
	}
	return result, nil
}

func (c *Client) vehicleStatusesV4(ctx context.Context, request VehicleStatusesRequest, cfg ClientConfig) (_ VehicleStatusesResponse, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("rFMS v4 vehicle statuses: %w", err)
		}
	}()
	// Build query parameters
	query := url.Values{}
	if request.LastVIN != "" {
		query.Set("lastVin", request.LastVIN)
	}
	if request.DateType != "" {
		query.Set("datetype", request.DateType)
	}
	if !request.StartTime.IsZero() {
		query.Set("starttime", rfmsv4oapi.Time(request.StartTime).String())
	}
	if !request.StopTime.IsZero() {
		query.Set("stoptime", rfmsv4oapi.Time(request.StopTime).String())
	}
	if request.VIN != "" {
		query.Set("vin", request.VIN)
	}
	if len(request.ContentFilter) > 0 {
		query.Set("contentFilter", strings.Join(request.ContentFilter, ","))
	}
	if len(request.TriggerFilter) > 0 {
		query.Set("triggerFilter", strings.Join(request.TriggerFilter, ","))
	}
	if request.LatestOnly {
		query.Set("latestOnly", "true")
	}
	// Build path with query parameters
	path := "/vehiclestatuses"
	if len(query) > 0 {
		path += "?" + query.Encode()
	}
	// Apply per-request configuration overrides
	fullURL := cfg.baseURL + path
	// Create the request
	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return VehicleStatusesResponse{}, fmt.Errorf("http request: %w", err)
	}
	// Set headers
	httpRequest.Header.Set("User-Agent", getUserAgent())
	httpRequest.Header.Set("Accept", "application/json; rfms=vehiclestatuses.v4.0")
	// Create HTTP client and make request
	client := c.httpClient(cfg)
	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		return VehicleStatusesResponse{}, fmt.Errorf("http request: %w", err)
	}
	defer httpResponse.Body.Close()
	if httpResponse.StatusCode != http.StatusOK {
		return VehicleStatusesResponse{}, newHTTPError(httpResponse)
	}
	data, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return VehicleStatusesResponse{}, fmt.Errorf("read response body: %w", err)
	}
	var response rfmsv4oapi.VehicleStatusResponseObject
	if err := json.Unmarshal(data, &response); err != nil {
		return VehicleStatusesResponse{}, fmt.Errorf("unmarshal v4 response body: %w", err)
	}
	var result VehicleStatusesResponse
	result.MoreDataAvailable = response.MoreDataAvailable != nil && *response.MoreDataAvailable
	for _, vehicleStatus := range response.VehicleStatusResponse.VehicleStatuses {
		result.VehicleStatuses = append(result.VehicleStatuses, convertv4.VehicleStatus(&vehicleStatus))
	}
	if response.RequestServerDateTime != nil {
		result.RequestServerDateTime = time.Time(*response.RequestServerDateTime)
	}
	return result, nil
}
