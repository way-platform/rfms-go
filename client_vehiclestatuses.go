package rfms

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type ListVehicleStatusesRequest struct {
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

type ListVehicleStatusesResponse struct {
	// Raw response body.
	Raw json.RawMessage `json:"-"`
	// VehicleStatuses in the response.
	VehicleStatuses []*VehicleStatus `json:"vehicleStatuses,omitempty"`
	// MoreDataAvailable indicates if there is more data available.
	MoreDataAvailable bool `json:"moreDataAvailable,omitempty"`
	// MoreDataAvailableLink is the link to the next page of data.
	MoreDataAvailableLink string `json:"moreDataAvailableLink,omitempty"`
	// RequestServerDateTime is the server time when the request was received.
	RequestServerDateTime Time `json:"requestServerDateTime,omitempty"`
}

func (c *Client) ListVehicleStatuses(
	ctx context.Context,
	request *ListVehicleStatusesRequest,
) (_ *ListVehicleStatusesResponse, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("list vehicle statuses: %w", err)
		}
	}()
	req, err := c.newRequest(ctx, http.MethodGet, "/vehiclestatuses", nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	switch c.apiVersion {
	case Version4:
		req.Header.Set("Accept", "application/json; rfms=vehiclestatuses.v4.0")
	case Version21:
		req.Header.Set("Accept", "application/vnd.fmsstandard.com.VehicleStatuses.v2.1+json; UTF-8")
	default:
		return nil, fmt.Errorf("unsupported API version: %s", c.apiVersion)
	}
	q := req.URL.Query()
	if request != nil {
		if request.LastVIN != "" {
			q.Set("lastVin", request.LastVIN)
		}
		if request.DateType != "" {
			q.Set("datetype", request.DateType)
		}
		if !request.StartTime.IsZero() {
			q.Set("starttime", request.StartTime.UTC().Format(time.RFC3339))
		}
		if !request.StopTime.IsZero() {
			q.Set("stoptime", request.StopTime.UTC().Format(time.RFC3339))
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
		VehicleStatusResponse ListVehicleStatusesResponse `json:"vehicleStatusResponse"`
		MoreDataAvailable     bool                        `json:"moreDataAvailable"`
		MoreDataAvailableLink string                      `json:"moreDataAvailableLink,omitempty"`
		RequestServerDateTime Time                        `json:"requestServerDateTime"`
	}
	if err := json.Unmarshal(data, &responseBody); err != nil {
		return nil, fmt.Errorf("unmarshal response body: %w", err)
	}
	var rawStatuses struct {
		VehicleStatusResponse struct {
			VehicleStatuses []json.RawMessage `json:"vehicleStatuses"`
		} `json:"vehicleStatusResponse"`
	}
	if err := json.Unmarshal(data, &rawStatuses); err != nil {
		return nil, fmt.Errorf("unmarshal raw vehicle statuses: %w", err)
	}
	for i, rawStatus := range rawStatuses.VehicleStatusResponse.VehicleStatuses {
		responseBody.VehicleStatusResponse.VehicleStatuses[i].Raw = rawStatus
	}
	responseBody.VehicleStatusResponse.Raw = data
	responseBody.VehicleStatusResponse.MoreDataAvailable = responseBody.MoreDataAvailable
	responseBody.VehicleStatusResponse.MoreDataAvailableLink = responseBody.MoreDataAvailableLink
	responseBody.VehicleStatusResponse.RequestServerDateTime = responseBody.RequestServerDateTime
	return &responseBody.VehicleStatusResponse, nil
}
