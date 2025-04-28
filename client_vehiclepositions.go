package rfms

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/way-platform/rfms-go/v4/rfmsv4"
)

type ListVehiclePositionsRequest struct {
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

type ListVehiclePositionsResponse struct {
	// Raw response body.
	Raw json.RawMessage `json:"-"`
	// VehiclePositions in the response.
	VehiclePositions []rfmsv4.VehiclePosition `json:"vehiclePositions,omitempty"`
	// MoreDataAvailable indicates if there is more data available.
	MoreDataAvailable bool `json:"moreDataAvailable,omitempty"`
	// MoreDataAvailableLink is the link to the next page of data.
	MoreDataAvailableLink string `json:"moreDataAvailableLink,omitempty"`
	// RequestServerDateTime is the server time when the request was received.
	RequestServerDateTime time.Time `json:"requestServerDateTime,omitzero"`
}

func (c *Client) ListVehiclePositions(
	ctx context.Context,
	request *ListVehiclePositionsRequest,
) (_ *ListVehiclePositionsResponse, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("list vehicle positions: %w", err)
		}
	}()
	req, err := c.newRequest(ctx, http.MethodGet, "/vehiclepositions", nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	switch c.config.apiVersion {
	case Version4:
		req.Header.Set("Accept", "application/json; rfms=vehiclepositions.v4.0")
	case Version21:
		req.Header.Set("Accept", "application/vnd.fmsstandard.com.vehiclepositions.v2.1+json; UTF-8")
	default:
		return nil, fmt.Errorf("unsupported API version: %s", c.config.apiVersion)
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
		if request.LatestOnly {
			q.Set("latestOnly", "true")
		}
		if request.TriggerFilter != "" {
			q.Set("triggerFilter", request.TriggerFilter)
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
	var responseBody rfmsv4.VehiclePositionsResponse
	if err := json.Unmarshal(data, &responseBody); err != nil {
		return nil, fmt.Errorf("unmarshal response body: %w", err)
	}
	// var rawPositions struct {
	// 	VehiclePositionResponse struct {
	// 		VehiclePositions []json.RawMessage `json:"vehiclePositions"`
	// 	} `json:"vehiclePositionResponse"`
	// }
	// if err := json.Unmarshal(data, &rawPositions); err != nil {
	// 	return nil, fmt.Errorf("unmarshal raw vehicle positions: %w", err)
	// }
	// Set raw data and propagate MoreDataAvailable/Link and RequestServerDateTime
	// responseBody.VehiclePositionResponse.Raw = data
	// responseBody.VehiclePositionResponse.MoreDataAvailable = responseBody.MoreDataAvailable
	// responseBody.VehiclePositionResponse.MoreDataAvailableLink = responseBody.MoreDataAvailableLink
	// responseBody.VehiclePositionResponse.RequestServerDateTime = responseBody.RequestServerDateTime
	// Set raw data for individual vehicle positions
	// for i, rawPosition := range rawPositions.VehiclePositionResponse.VehiclePositions {
	// 	if i < len(responseBody.VehiclePositionResponse.VehiclePositions) {
	// 		responseBody.VehiclePositionResponse.VehiclePositions[i].Raw = rawPosition
	// 	}
	// }
	return &ListVehiclePositionsResponse{
		VehiclePositions: *responseBody.VehiclePositionResponse.VehiclePositions,
	}, nil
}
