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
	"google.golang.org/protobuf/types/known/timestamppb"
)

// VehicleStatuses implements the rFMS API method "GET /vehiclestatuses".
func (c *Client) VehicleStatuses(
	ctx context.Context,
	request *rfmsv5.VehicleStatusesRequest,
) (_ *rfmsv5.VehicleStatusesResponse, err error) {
	switch c.config.apiVersion {
	case V2_1:
		return c.vehicleStatusesV2(ctx, request)
	case V4:
		return c.vehicleStatusesV4(ctx, request)
	default:
		return nil, fmt.Errorf("unsupported API version")
	}
}

func (c *Client) vehicleStatusesV2(
	ctx context.Context,
	request *rfmsv5.VehicleStatusesRequest,
) (_ *rfmsv5.VehicleStatusesResponse, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("rFMS v2 vehicle statuses: %w", err)
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
	if len(request.GetContentFilter()) > 0 {
		query.Set("contentFilter", strings.Join(request.GetContentFilter(), ","))
	}
	if len(request.GetTriggerFilter()) > 0 {
		query.Set("triggerFilter", strings.Join(request.GetTriggerFilter(), ","))
	}
	if request.GetLatestOnly() {
		query.Set("latestOnly", "true")
	}
	path := "/vehiclestatuses"
	if len(query) > 0 {
		path += "?" + query.Encode()
	}
	fullURL := c.config.baseURL + path
	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("http request: %w", err)
	}
	httpRequest.Header.Set("User-Agent", getUserAgent())
	httpRequest.Header.Set("Accept", "application/vnd.fmsstandard.com.Vehiclestatuses.v2.1+json")
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
	var response rfmsv2oapi.VehicleStatuses
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, fmt.Errorf("unmarshal v2 response body: %w", err)
	}
	resp := &rfmsv5.VehicleStatusesResponse{}
	statuses := make([]*rfmsv5.VehicleStatus, 0, len(response.VehicleStatus))
	for _, vehicleStatus := range response.VehicleStatus {
		statuses = append(statuses, convertv2.VehicleStatus(&vehicleStatus))
	}
	resp.SetVehicleStatuses(statuses)
	if response.MoreDataAvailable != nil {
		resp.SetMoreDataAvailable(*response.MoreDataAvailable)
	}
	if response.RequestServerDateTime != nil {
		resp.SetRequestServerDateTime(timestamppb.New(*response.RequestServerDateTime))
	}
	return resp, nil
}

func (c *Client) vehicleStatusesV4(
	ctx context.Context,
	request *rfmsv5.VehicleStatusesRequest,
) (_ *rfmsv5.VehicleStatusesResponse, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("rFMS v4 vehicle statuses: %w", err)
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
	if len(request.GetContentFilter()) > 0 {
		query.Set("contentFilter", strings.Join(request.GetContentFilter(), ","))
	}
	if len(request.GetTriggerFilter()) > 0 {
		query.Set("triggerFilter", strings.Join(request.GetTriggerFilter(), ","))
	}
	if request.GetLatestOnly() {
		query.Set("latestOnly", "true")
	}
	path := "/vehiclestatuses"
	if len(query) > 0 {
		path += "?" + query.Encode()
	}
	fullURL := c.config.baseURL + path
	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("http request: %w", err)
	}
	httpRequest.Header.Set("User-Agent", getUserAgent())
	httpRequest.Header.Set("Accept", "application/json; rfms=vehiclestatuses.v4.0")
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
	var response rfmsv4oapi.VehicleStatusResponseObject
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, fmt.Errorf("unmarshal v4 response body: %w", err)
	}
	resp := &rfmsv5.VehicleStatusesResponse{}
	statuses := make([]*rfmsv5.VehicleStatus, 0, len(response.VehicleStatusResponse.VehicleStatuses))
	for _, vehicleStatus := range response.VehicleStatusResponse.VehicleStatuses {
		statuses = append(statuses, convertv4.VehicleStatus(&vehicleStatus))
	}
	resp.SetVehicleStatuses(statuses)
	if response.MoreDataAvailable != nil {
		resp.SetMoreDataAvailable(*response.MoreDataAvailable)
	}
	if response.RequestServerDateTime != nil {
		resp.SetRequestServerDateTime(timestamppb.New(time.Time(*response.RequestServerDateTime)))
	}
	return resp, nil
}
