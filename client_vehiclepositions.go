package rfms

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/way-platform/rfms-go/internal/convert/convertv2"
	"github.com/way-platform/rfms-go/internal/convert/convertv4"
	"github.com/way-platform/rfms-go/internal/openapi/rfmsv2oapi"
	"github.com/way-platform/rfms-go/internal/openapi/rfmsv4oapi"
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/connect/rfms/v5"
)

// VehiclePositionsRequest is the request for the [Client.VehiclePositions] method.
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

// VehiclePositionsResponse is the response for the [Client.VehiclePositions] method.
type VehiclePositionsResponse struct {
	// VehiclePositions in the response.
	VehiclePositions []*rfmsv5.VehiclePosition `json:"vehiclePositions"`
	// MoreDataAvailable indicates if there is more data available.
	MoreDataAvailable bool `json:"moreDataAvailable"`
	// RequestServerDateTime is the server time when the request was received.
	RequestServerDateTime time.Time `json:"requestServerDateTime,omitzero"`
}

// VehiclePositions implements the rFMS API method "GET /vehiclepositions".
func (c *Client) VehiclePositions(ctx context.Context, request VehiclePositionsRequest) (_ VehiclePositionsResponse, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("rFMS %s vehicle positions: %w", c.config.apiVersion, err)
		}
	}()
	switch c.config.apiVersion {
	case V2_1:
		return c.vehiclePositionsV2(ctx, request)
	case V4:
		return c.vehiclePositionsV4(ctx, request)
	default:
		return VehiclePositionsResponse{}, fmt.Errorf("unsupported API version")
	}
}

func (c *Client) vehiclePositionsV2(ctx context.Context, request VehiclePositionsRequest) (VehiclePositionsResponse, error) {
	httpRequest, err := c.newRequest(ctx, http.MethodGet, "/vehiclepositions", nil)
	if err != nil {
		return VehiclePositionsResponse{}, fmt.Errorf("create request: %w", err)
	}
	httpRequest.Header.Set("Accept", "application/vnd.fmsstandard.com.vehiclepositions.v2.1+json; UTF-8")
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
	if request.LatestOnly {
		q.Set("latestOnly", "true")
	}
	if request.TriggerFilter != "" {
		q.Set("triggerFilter", request.TriggerFilter)
	}
	httpRequest.URL.RawQuery = q.Encode()
	httpResponse, err := c.do(httpRequest)
	if err != nil {
		return VehiclePositionsResponse{}, fmt.Errorf("send request: %w", err)
	}
	defer httpResponse.Body.Close()
	if httpResponse.StatusCode != http.StatusOK {
		return VehiclePositionsResponse{}, newHTTPError(httpResponse)
	}
	data, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return VehiclePositionsResponse{}, fmt.Errorf("read response body: %w", err)
	}
	var response rfmsv2oapi.VehiclePositions
	if err := json.Unmarshal(data, &response); err != nil {
		return VehiclePositionsResponse{}, fmt.Errorf("unmarshal response body: %w", err)
	}
	var result VehiclePositionsResponse
	result.MoreDataAvailable = response.MoreDataAvailable != nil && *response.MoreDataAvailable
	for _, vehiclePosition := range response.VehiclePosition {
		result.VehiclePositions = append(result.VehiclePositions, convertv2.VehiclePosition(&vehiclePosition))
	}
	if response.RequestServerDateTime != nil {
		result.RequestServerDateTime = *response.RequestServerDateTime
	}
	return result, nil
}

func (c *Client) vehiclePositionsV4(ctx context.Context, request VehiclePositionsRequest) (VehiclePositionsResponse, error) {
	httpRequest, err := c.newRequest(ctx, http.MethodGet, "/vehiclepositions", nil)
	if err != nil {
		return VehiclePositionsResponse{}, fmt.Errorf("create request: %w", err)
	}
	httpRequest.Header.Set("Accept", "application/json; rfms=vehiclepositions.v4.0")
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
	if request.LatestOnly {
		q.Set("latestOnly", "true")
	}
	if request.TriggerFilter != "" {
		q.Set("triggerFilter", request.TriggerFilter)
	}
	httpRequest.URL.RawQuery = q.Encode()
	httpResponse, err := c.do(httpRequest)
	if err != nil {
		return VehiclePositionsResponse{}, fmt.Errorf("send request: %w", err)
	}
	defer httpResponse.Body.Close()
	if httpResponse.StatusCode != http.StatusOK {
		return VehiclePositionsResponse{}, newHTTPError(httpResponse)
	}
	data, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return VehiclePositionsResponse{}, fmt.Errorf("read response body: %w", err)
	}
	var response rfmsv4oapi.VehiclePositionResponseObject
	if err := json.Unmarshal(data, &response); err != nil {
		return VehiclePositionsResponse{}, fmt.Errorf("unmarshal vresponse body: %w", err)
	}
	var result VehiclePositionsResponse
	result.MoreDataAvailable = response.MoreDataAvailable != nil && *response.MoreDataAvailable
	for _, vehiclePosition := range response.VehiclePositionResponse.VehiclePositions {
		result.VehiclePositions = append(result.VehiclePositions, convertv4.VehiclePosition(&vehiclePosition))
	}
	if response.RequestServerDateTime != nil {
		result.RequestServerDateTime = time.Time(*response.RequestServerDateTime)
	}
	return result, nil
}
