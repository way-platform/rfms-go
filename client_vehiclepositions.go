package rfms

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/way-platform/rfms-go/internal/convert/convertv2"
	"github.com/way-platform/rfms-go/internal/convert/convertv4"
	"github.com/way-platform/rfms-go/internal/openapi/rfmsv2oapi"
	"github.com/way-platform/rfms-go/internal/openapi/rfmsv4oapi"
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/connect/rfms/v5"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// VehiclePositions implements the rFMS API method "GET /vehiclepositions".
func (c *Client) VehiclePositions(
	ctx context.Context,
	request *rfmsv5.VehiclePositionsRequest,
) (_ *rfmsv5.VehiclePositionsResponse, err error) {
	switch c.config.apiVersion {
	case V2_1:
		return c.vehiclePositionsV2(ctx, request)
	case V4:
		return c.vehiclePositionsV4(ctx, request)
	default:
		return nil, fmt.Errorf("unsupported API version")
	}
}

func (c *Client) vehiclePositionsV2(
	ctx context.Context,
	request *rfmsv5.VehiclePositionsRequest,
) (_ *rfmsv5.VehiclePositionsResponse, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("rFMS v2 vehicle positions: %w", err)
		}
	}()
	query := url.Values{}
	if request.GetLastVin() != "" {
		query.Set("lastVin", request.GetLastVin())
	}
	if request.GetDateType() != "" {
		query.Set("datetype", request.GetDateType())
	}
	if request.HasStartTime() {
		query.Set("starttime", rfmsv4oapi.Time(request.GetStartTime().AsTime()).String())
	}
	if request.HasStopTime() {
		query.Set("stoptime", rfmsv4oapi.Time(request.GetStopTime().AsTime()).String())
	}
	if request.GetVin() != "" {
		query.Set("vin", request.GetVin())
	}
	if request.GetLatestOnly() {
		query.Set("latestOnly", "true")
	}
	if request.GetTriggerFilter() != "" {
		query.Set("triggerFilter", request.GetTriggerFilter())
	}
	path := "/vehiclepositions"
	if len(query) > 0 {
		path += "?" + query.Encode()
	}
	fullURL := c.config.baseURL + path
	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("http request: %w", err)
	}
	httpRequest.Header.Set("User-Agent", getUserAgent())
	httpRequest.Header.Set("Accept", "application/vnd.fmsstandard.com.Vehiclepositions.v2.1+json")
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
	var response rfmsv2oapi.VehiclePositions
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, fmt.Errorf("unmarshal response body: %w", err)
	}
	resp := &rfmsv5.VehiclePositionsResponse{}
	positions := make([]*rfmsv5.VehiclePosition, 0, len(response.VehiclePosition))
	for _, vehiclePosition := range response.VehiclePosition {
		positions = append(positions, convertv2.VehiclePosition(&vehiclePosition))
	}
	resp.SetVehiclePositions(positions)
	if response.MoreDataAvailable != nil {
		resp.SetMoreDataAvailable(*response.MoreDataAvailable)
	}
	if response.RequestServerDateTime != nil {
		resp.SetRequestServerDateTime(timestamppb.New(*response.RequestServerDateTime))
	}
	return resp, nil
}

func (c *Client) vehiclePositionsV4(
	ctx context.Context,
	request *rfmsv5.VehiclePositionsRequest,
) (_ *rfmsv5.VehiclePositionsResponse, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("rFMS v4 vehicle positions: %w", err)
		}
	}()
	query := url.Values{}
	if request.GetLastVin() != "" {
		query.Set("lastVin", request.GetLastVin())
	}
	if request.GetDateType() != "" {
		query.Set("datetype", request.GetDateType())
	}
	if request.HasStartTime() {
		query.Set("starttime", rfmsv4oapi.Time(request.GetStartTime().AsTime()).String())
	}
	if request.HasStopTime() {
		query.Set("stoptime", rfmsv4oapi.Time(request.GetStopTime().AsTime()).String())
	}
	if request.GetVin() != "" {
		query.Set("vin", request.GetVin())
	}
	if request.GetLatestOnly() {
		query.Set("latestOnly", "true")
	}
	if request.GetTriggerFilter() != "" {
		query.Set("triggerFilter", request.GetTriggerFilter())
	}
	path := "/vehiclepositions"
	if len(query) > 0 {
		path += "?" + query.Encode()
	}
	fullURL := c.config.baseURL + path
	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("http request: %w", err)
	}
	httpRequest.Header.Set("User-Agent", getUserAgent())
	httpRequest.Header.Set("Accept", "application/json; rfms=vehiclepositions.v4.0")
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
	var response rfmsv4oapi.VehiclePositionResponseObject
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, fmt.Errorf("unmarshal response body: %w", err)
	}
	resp := &rfmsv5.VehiclePositionsResponse{}
	positions := make([]*rfmsv5.VehiclePosition, 0, len(response.VehiclePositionResponse.VehiclePositions))
	for _, vehiclePosition := range response.VehiclePositionResponse.VehiclePositions {
		positions = append(positions, convertv4.VehiclePosition(&vehiclePosition))
	}
	resp.SetVehiclePositions(positions)
	if response.MoreDataAvailable != nil {
		resp.SetMoreDataAvailable(*response.MoreDataAvailable)
	}
	if response.RequestServerDateTime != nil {
		resp.SetRequestServerDateTime(timestamppb.New(time.Time(*response.RequestServerDateTime)))
	}
	return resp, nil
}
