package pages

import "github.com/way-platform/rfms-go/cmd/rfms/v4/internal/dashboard/ui/templates/cards"

// Index is the data for the index page.
type Index struct {
	Title            string
	VehicleCountCard cards.VehicleCount
	WorldMapCard     cards.WorldMap
}
