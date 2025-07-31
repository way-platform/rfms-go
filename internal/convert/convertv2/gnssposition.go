package convertv2

import (
	"github.com/way-platform/rfms-go/internal/openapi/rfmsv2oapi"
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/rfms/v5"
)

func gnssPosition(input *rfmsv2oapi.GNSSPositionType) *rfmsv5.GnssPosition {
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
		output.SetTime(input.PositionDateTime.UnixMicro())
	}
	return &output
}
