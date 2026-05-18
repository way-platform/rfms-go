package rfms_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"connectrpc.com/connect"
	rfms "github.com/way-platform/rfms-go"
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/connect/rfms/v5"
)

func newTestClient(t *testing.T, srv *httptest.Server) *rfms.Client {
	t.Helper()
	client, err := rfms.NewClient(
		rfms.WithBaseURL(srv.URL),
		rfms.WithVersion(rfms.V4),
		rfms.WithBasicAuth("test", "test"),
		rfms.WithRetryCount(0),
	)
	if err != nil {
		t.Fatal(err)
	}
	return client
}

func TestVehicles(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/vehicles" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		resp := map[string]any{
			"moreDataAvailable": false,
			"vehicleResponse": map[string]any{
				"vehicles": []map[string]any{
					{"vin": "YS2R4X20009876543"},
					{"vin": "WDB96340310987654"},
				},
			},
		}
		_ = json.NewEncoder(w).Encode(resp)
	}))
	defer srv.Close()

	client := newTestClient(t, srv)
	resp, err := client.Vehicles(context.Background(), rfmsv5.VehiclesRequest_builder{}.Build())
	if err != nil {
		t.Fatal(err)
	}
	if got := len(resp.GetVehicles()); got != 2 {
		t.Fatalf("expected 2 vehicles, got %d", got)
	}
	if got := resp.GetVehicles()[0].GetVin(); got != "YS2R4X20009876543" {
		t.Fatalf("expected VIN YS2R4X20009876543, got %s", got)
	}
	if got := resp.GetVehicles()[1].GetVin(); got != "WDB96340310987654" {
		t.Fatalf("expected VIN WDB96340310987654, got %s", got)
	}
	if resp.GetMoreDataAvailable() {
		t.Fatal("expected moreDataAvailable=false")
	}
}

func TestVehiclePositions(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/vehiclepositions" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		resp := map[string]any{
			"moreDataAvailable": false,
			"vehiclePositionResponse": map[string]any{
				"vehiclePositions": []map[string]any{
					{
						"vin": "YS2R4X20009876543",
						"gnssPosition": map[string]any{
							"latitude":  57.7089,
							"longitude": 11.9746,
						},
					},
				},
			},
		}
		_ = json.NewEncoder(w).Encode(resp)
	}))
	defer srv.Close()

	client := newTestClient(t, srv)
	req := rfmsv5.VehiclePositionsRequest_builder{}.Build()
	resp, err := client.VehiclePositions(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	if got := len(resp.GetVehiclePositions()); got != 1 {
		t.Fatalf("expected 1 position, got %d", got)
	}
	pos := resp.GetVehiclePositions()[0]
	if got := pos.GetVin(); got != "YS2R4X20009876543" {
		t.Fatalf("expected VIN YS2R4X20009876543, got %s", got)
	}
}

func TestVehicleStatuses(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/vehiclestatuses" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		resp := map[string]any{
			"moreDataAvailable": false,
			"vehicleStatusResponse": map[string]any{
				"vehicleStatuses": []map[string]any{
					{
						"vin":             "YS2R4X20009876543",
						"createdDateTime": "2024-01-15T10:30:00Z",
					},
				},
			},
		}
		_ = json.NewEncoder(w).Encode(resp)
	}))
	defer srv.Close()

	client := newTestClient(t, srv)
	req := rfmsv5.VehicleStatusesRequest_builder{}.Build()
	resp, err := client.VehicleStatuses(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	if got := len(resp.GetVehicleStatuses()); got != 1 {
		t.Fatalf("expected 1 status, got %d", got)
	}
	status := resp.GetVehicleStatuses()[0]
	if got := status.GetVin(); got != "YS2R4X20009876543" {
		t.Fatalf("expected VIN YS2R4X20009876543, got %s", got)
	}
}

func TestErrorCodeMapping(t *testing.T) {
	tests := []struct {
		statusCode int
		wantCode   connect.Code
	}{
		{http.StatusBadRequest, connect.CodeInvalidArgument},
		{http.StatusUnauthorized, connect.CodeUnauthenticated},
		{http.StatusForbidden, connect.CodePermissionDenied},
		{http.StatusNotFound, connect.CodeNotFound},
		{http.StatusConflict, connect.CodeAlreadyExists},
		{http.StatusTooManyRequests, connect.CodeResourceExhausted},
		{http.StatusNotImplemented, connect.CodeUnimplemented},
		{http.StatusServiceUnavailable, connect.CodeUnavailable},
		{http.StatusGatewayTimeout, connect.CodeDeadlineExceeded},
		{http.StatusInternalServerError, connect.CodeInternal},
		{http.StatusTeapot, connect.CodeUnknown},
	}
	for _, tt := range tests {
		t.Run(http.StatusText(tt.statusCode), func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tt.statusCode)
				_ = json.NewEncoder(w).Encode(map[string]any{
					"error":             "test_error",
					"error_description": "something went wrong",
				})
			}))
			defer srv.Close()

			client := newTestClient(t, srv)
			_, err := client.Vehicles(context.Background(), rfmsv5.VehiclesRequest_builder{}.Build())
			if err == nil {
				t.Fatal("expected error, got nil")
			}
			var connectErr *connect.Error
			if !errors.As(err, &connectErr) {
				t.Fatalf("expected *connect.Error, got %T: %v", err, err)
			}
			if connectErr.Code() != tt.wantCode {
				t.Fatalf("expected code %v, got %v", tt.wantCode, connectErr.Code())
			}
		})
	}
}

func TestRateLimitError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Retry-After", "30")
		w.WriteHeader(http.StatusTooManyRequests)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"error":             "rate_limit_exceeded",
			"error_description": "too many requests",
		})
	}))
	defer srv.Close()

	client := newTestClient(t, srv)
	_, err := client.Vehicles(context.Background(), rfmsv5.VehiclesRequest_builder{}.Build())
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	var connectErr *connect.Error
	if !errors.As(err, &connectErr) {
		t.Fatalf("expected *connect.Error, got %T: %v", err, err)
	}
	if connectErr.Code() != connect.CodeResourceExhausted {
		t.Fatalf("expected CodeResourceExhausted, got %v", connectErr.Code())
	}
}
