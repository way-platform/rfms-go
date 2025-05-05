package server

import (
	"net/http"

	"github.com/way-platform/rfms-go/cmd/rfms/v4/internal/dashboard/ui/templates/cards"
	"github.com/way-platform/rfms-go/cmd/rfms/v4/internal/dashboard/ui/templates/pages"
	"github.com/way-platform/rfms-go/v4/api/rfmsv4"
)

func (s *Server) registerRoutes(mux *http.ServeMux) {
	mux.Handle(s.staticRoute())
	mux.Handle(s.indexRoute())
}

func (s *Server) staticRoute() (string, http.Handler) {
	return "/static/", s.staticHandler
}

func (s *Server) indexRoute() (string, http.Handler) {
	return "/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vehiclesResponse, err := s.client.Vehicles(r.Context(), &rfmsv4.VehiclesRequest{})
		if err != nil {
			s.handleError(w, r, "failed to get vehicles", err)
			return
		}
		vehiclePositionsResponse, err := s.client.VehiclePositions(r.Context(), vehiclesResponse.Vehicles)
		if err != nil {
			s.handleError(w, r, "failed to get vehicle positions", err)
			return
		}
		vehicleStatusesResponse, err := s.client.GetVehicleStatuses(r.Context(), vehiclesResponse.Vehicles)
		if err != nil {
			s.handleError(w, r, "failed to get vehicle statuses", err)
			return
		}
		s.renderTemplate(w, r, "index.html", pages.Index{
			Title: "Fleet Dashboard",
			VehicleCountCard: cards.VehicleCount{
				TotalCount: 97,
			},
			WorldMapCard: cards.WorldMap{
				VehiclePositions: []rfmsv4.VehiclePosition{
					{
						VIN:          "EXAMPLE1",
						GNSSPosition: &rfmsv4.GNSSPosition{Latitude: ptr(40.7128), Longitude: ptr(-74.0060)},
					},
					{
						VIN:          "EXAMPLE2",
						GNSSPosition: &rfmsv4.GNSSPosition{Latitude: ptr(51.5074), Longitude: ptr(-0.1278)},
					},
					{
						VIN:          "EXAMPLE3",
						GNSSPosition: &rfmsv4.GNSSPosition{Latitude: ptr(35.6762), Longitude: ptr(139.6503)},
					},
					{
						VIN:          "EXAMPLE4",
						GNSSPosition: &rfmsv4.GNSSPosition{Latitude: ptr(-33.8688), Longitude: ptr(151.2093)},
					},
					{
						VIN:          "EXAMPLE5",
						GNSSPosition: &rfmsv4.GNSSPosition{Latitude: ptr(52.5200), Longitude: ptr(13.4050)},
					},
					{
						VIN:          "EXAMPLE6",
						GNSSPosition: &rfmsv4.GNSSPosition{Latitude: ptr(48.8566), Longitude: ptr(2.3522)},
					},
					{
						VIN:          "EXAMPLE7",
						GNSSPosition: &rfmsv4.GNSSPosition{Latitude: ptr(31.2304), Longitude: ptr(121.4737)},
					},
					{
						VIN:          "EXAMPLE8",
						GNSSPosition: &rfmsv4.GNSSPosition{Latitude: ptr(-23.5505), Longitude: ptr(-46.6333)},
					},
					{
						VIN:          "EXAMPLE9",
						GNSSPosition: &rfmsv4.GNSSPosition{Latitude: ptr(25.2048), Longitude: ptr(55.2708)},
					},
					{
						VIN:          "EXAMPLE10",
						GNSSPosition: &rfmsv4.GNSSPosition{Latitude: ptr(19.0760), Longitude: ptr(72.8777)},
					},
				},
			},
		})
	})
}

func ptr[T any](v T) *T {
	return &v
}
