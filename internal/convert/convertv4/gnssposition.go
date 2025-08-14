package convertv4

import (
	"time"

	"github.com/way-platform/rfms-go/internal/openapi/rfmsv4oapi"
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/connect/rfms/v5"
)

// gnssPosition converts an rFMS v4 GNSS position to proto.
func gnssPosition(input *rfmsv4oapi.GNSSPositionObject) *rfmsv5.GnssPosition {
	var output rfmsv5.GnssPosition
	if input.Latitude != nil {
		output.SetLatitude(*input.Latitude)
	}
	if input.Longitude != nil {
		output.SetLongitude(*input.Longitude)
	}
	if input.Heading != nil {
		output.SetHeadingDeg(float64(*input.Heading))
	}
	if input.Altitude != nil {
		output.SetAltitudeM(float64(*input.Altitude))
	}
	if input.Speed != nil {
		output.SetSpeedKmh(*input.Speed)
	}
	if input.PositionDateTime != nil {
		output.SetTime(time.Time(*input.PositionDateTime).UTC().Format(time.RFC3339Nano))
	}
	return &output
}
