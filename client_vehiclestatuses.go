package rfms

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func (c *Client) VehicleStatuses(ctx context.Context, request VehicleStatusesRequest) (_ VehicleStatusesResponse, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("rFMS %s vehicle statuses: %w", c.config.apiVersion, err)
		}
	}()
	switch c.config.apiVersion {
	case V2_1:
		return c.vehicleStatusesV2(ctx, request)
	case V4:
		return c.vehicleStatusesV4(ctx, request)
	default:
		return VehicleStatusesResponse{}, fmt.Errorf("unsupported API version")
	}
}

func (c *Client) vehicleStatusesV2(ctx context.Context, request VehicleStatusesRequest) (VehicleStatusesResponse, error) {
	httpRequest, err := c.newRequest(ctx, http.MethodGet, "/vehiclestatuses", nil)
	if err != nil {
		return VehicleStatusesResponse{}, fmt.Errorf("create request: %w", err)
	}
	httpRequest.Header.Set("Accept", "application/vnd.fmsstandard.com.VehicleStatuses.v2.1+json; UTF-8")
	q := httpRequest.URL.Query()
	if request.LastVIN != "" {
		q.Set("lastVin", request.LastVIN)
	}
	if request.DateType != "" {
		q.Set("datetype", request.DateType)
	}
	if !request.StartTime.IsZero() {
		q.Set("starttime", rfmsv4oapi.Time(request.StartTime).String())
	}
	if !request.StopTime.IsZero() {
		q.Set("stoptime", rfmsv4oapi.Time(request.StopTime).String())
	}
	if request.VIN != "" {
		q.Set("vin", request.VIN)
	}
	if len(request.ContentFilter) > 0 {
		q.Set("contentFilter", strings.Join(request.ContentFilter, ","))
	}
	if len(request.TriggerFilter) > 0 {
		q.Set("triggerFilter", strings.Join(request.TriggerFilter, ","))
	}
	if request.LatestOnly {
		q.Set("latestOnly", "true")
	}
	httpRequest.URL.RawQuery = q.Encode()
	httpResponse, err := c.do(httpRequest)
	if err != nil {
		return VehicleStatusesResponse{}, fmt.Errorf("send request: %w", err)
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

func (c *Client) vehicleStatusesV4(ctx context.Context, request VehicleStatusesRequest) (VehicleStatusesResponse, error) {
	httpRequest, err := c.newRequest(ctx, http.MethodGet, "/vehiclestatuses", nil)
	if err != nil {
		return VehicleStatusesResponse{}, fmt.Errorf("create request: %w", err)
	}
	httpRequest.Header.Set("Accept", "application/json; rfms=vehiclestatuses.v4.0")
	q := httpRequest.URL.Query()
	if request.LastVIN != "" {
		q.Set("lastVin", request.LastVIN)
	}
	if request.DateType != "" {
		q.Set("datetype", request.DateType)
	}
	if !request.StartTime.IsZero() {
		q.Set("starttime", rfmsv4oapi.Time(request.StartTime).String())
	}
	if !request.StopTime.IsZero() {
		q.Set("stoptime", rfmsv4oapi.Time(request.StopTime).String())
	}
	if request.VIN != "" {
		q.Set("vin", request.VIN)
	}
	if len(request.ContentFilter) > 0 {
		q.Set("contentFilter", strings.Join(request.ContentFilter, ","))
	}
	if len(request.TriggerFilter) > 0 {
		q.Set("triggerFilter", strings.Join(request.TriggerFilter, ","))
	}
	if request.LatestOnly {
		q.Set("latestOnly", "true")
	}
	httpRequest.URL.RawQuery = q.Encode()
	httpResponse, err := c.do(httpRequest)
	if err != nil {
		return VehicleStatusesResponse{}, fmt.Errorf("send request: %w", err)
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
