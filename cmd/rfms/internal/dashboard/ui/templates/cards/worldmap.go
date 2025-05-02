package cards

import "github.com/way-platform/rfms-go/v4/api/rfmsv4"

// WorldMap is the data for the world map card.
type WorldMap struct {
	// VehiclePositions is the list of vehicle positions.
	VehiclePositions []rfmsv4.VehiclePosition
}
