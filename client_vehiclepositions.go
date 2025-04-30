package rfms

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/way-platform/rfms-go/v4/api/compat/rfmsv2tov4"
	"github.com/way-platform/rfms-go/v4/api/rfmsv2"
	"github.com/way-platform/rfms-go/v4/api/rfmsv4"
)

type VehiclePositionsRequest struct {
	// LastVIN is the last VIN included in the previous response.
	LastVIN string
	// DateType indicates whether the start/stop times are compared to created or received time.
	DateType string
	// StartTime to filter positions (only positions after this time).
	StartTime time.Time
	// StopTime to filter positions (only positions before this time).
	StopTime time.Time
	// VIN to filter positions for a specific vehicle.
	VIN string
	// LatestOnly returns only the latest position for each vehicle.
	LatestOnly bool
	// TriggerFilter filters positions by trigger type.
	TriggerFilter string
}

type VehiclePositionsResponse struct {
	// VehiclePositions in the response.
	VehiclePositions []rfmsv4.VehiclePosition `json:"vehiclePositions"`
	// MoreDataAvailable indicates if there is more data available.
	MoreDataAvailable bool `json:"moreDataAvailable"`
	// RequestServerDateTime is the server time when the request was received.
	RequestServerDateTime time.Time `json:"requestServerDateTime,omitzero"`
}

func (c *Client) VehiclePositions(
	ctx context.Context,
	request *VehiclePositionsRequest,
) (_ *VehiclePositionsResponse, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("rFMS vehicle positions: %w", err)
		}
	}()
	httpRequest, err := c.newRequest(ctx, http.MethodGet, "/vehiclepositions", nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	switch c.config.apiVersion {
	case V4:
		httpRequest.Header.Set("Accept", "application/json; rfms=vehiclepositions.v4.0")
	case V2_1:
		httpRequest.Header.Set("Accept", "application/vnd.fmsstandard.com.vehiclepositions.v2.1+json; UTF-8")
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
		if request.LatestOnly {
			q.Set("latestOnly", "true")
		}
		if request.TriggerFilter != "" {
			q.Set("triggerFilter", request.TriggerFilter)
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
	var v4Response rfmsv4.VehiclePositionsResponse
	switch c.config.apiVersion {
	case V4:
		if err := json.Unmarshal(data, &v4Response); err != nil {
			return nil, fmt.Errorf("unmarshal vresponse body: %w", err)
		}
	case V2_1:
		var v2Response rfmsv2.VehiclePositionsResponse
		if err := json.Unmarshal(data, &v2Response); err != nil {
			return nil, fmt.Errorf("unmarshal v2 response body: %w", err)
		}
		v4Response = *rfmsv2tov4.ConvertVehiclePositionsResponse(&v2Response)
	default:
		return nil, fmt.Errorf("unsupported API version: %s", c.config.apiVersion)
	}
	return &VehiclePositionsResponse{
		VehiclePositions:      v4Response.VehiclePositionResponse.VehiclePositions,
		MoreDataAvailable:     v4Response.MoreDataAvailable,
		RequestServerDateTime: time.Time(v4Response.RequestServerDateTime),
	}, nil
}
