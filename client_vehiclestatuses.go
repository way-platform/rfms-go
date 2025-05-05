package rfms

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/way-platform/rfms-go/api/compat/rfmsv2tov4"
	"github.com/way-platform/rfms-go/api/rfmsv2"
	"github.com/way-platform/rfms-go/api/rfmsv4"
)

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

type VehicleStatusesResponse struct {
	// VehicleStatuses in the response.
	VehicleStatuses []rfmsv4.VehicleStatus `json:"vehicleStatuses"`
	// MoreDataAvailable indicates if there is more data available.
	MoreDataAvailable bool `json:"moreDataAvailable"`
	// RequestServerDateTime is the server time when the request was received.
	RequestServerDateTime time.Time `json:"requestServerDateTime,omitzero"`
}

func (c *Client) VehicleStatuses(
	ctx context.Context,
	request *VehicleStatusesRequest,
) (_ *VehicleStatusesResponse, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("rFMS vehicle statuses: %w", err)
		}
	}()
	httpRequest, err := c.newRequest(ctx, http.MethodGet, "/vehiclestatuses", nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	switch c.config.apiVersion {
	case V4:
		httpRequest.Header.Set("Accept", "application/json; rfms=vehiclestatuses.v4.0")
	case V2_1:
		httpRequest.Header.Set("Accept", "application/vnd.fmsstandard.com.VehicleStatuses.v2.1+json; UTF-8")
	default:
		return nil, fmt.Errorf("unsupported API version: %s", c.config.apiVersion)
	}
	q := httpRequest.URL.Query()
	if request != nil {
		if request.LastVIN != "" {
			q.Set("lastVin", request.LastVIN)
		}
		if request.DateType != "" {
			q.Set("datetype", request.DateType)
		}
		if !request.StartTime.IsZero() {
			q.Set("starttime", rfmsv4.Time(request.StartTime).String())
		}
		if !request.StopTime.IsZero() {
			q.Set("stoptime", rfmsv4.Time(request.StopTime).String())
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
	var v4Response rfmsv4.VehicleStatusesResponse
	if err := json.Unmarshal(data, &v4Response); err != nil {
		return nil, fmt.Errorf("unmarshal v4 response body: %w", err)
	}
	switch c.config.apiVersion {
	case V4:
		if err := json.Unmarshal(data, &v4Response); err != nil {
			return nil, fmt.Errorf("unmarshal vresponse body: %w", err)
		}
	case V2_1:
		var v2Response rfmsv2.VehicleStatusesResponse
		if err := json.Unmarshal(data, &v2Response); err != nil {
			return nil, fmt.Errorf("unmarshal v2 response body: %w", err)
		}
		v4Response = *rfmsv2tov4.ConvertVehicleStatusesResponse(&v2Response)
	default:
		return nil, fmt.Errorf("unsupported API version: %s", c.config.apiVersion)
	}
	return &VehicleStatusesResponse{
		VehicleStatuses:       v4Response.VehicleStatusResponse.VehicleStatuses,
		MoreDataAvailable:     v4Response.MoreDataAvailable,
		RequestServerDateTime: time.Time(v4Response.RequestServerDateTime),
	}, nil
}
