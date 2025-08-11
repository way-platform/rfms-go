package convertv4

import (
	"time"

	"github.com/way-platform/rfms-go/internal/openapi/rfmsv4oapi"
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/rfms/v5"
)

// VehiclePosition converts an rFMS v4 vehicle position to proto.
func VehiclePosition(input *rfmsv4oapi.VehiclePositionObject) *rfmsv5.VehiclePosition {
	var output rfmsv5.VehiclePosition
	if input.VIN != nil {
		output.SetVin(string(*input.VIN))
	}
	if input.TriggerType != nil {
		output.SetTrigger(trigger(input.TriggerType))
	}
	if input.CreatedDateTime != nil {
		output.SetCreateTime(time.Time(*input.CreatedDateTime).UTC().Format(time.RFC3339Nano))
	}
	if input.ReceivedDateTime != nil {
		output.SetReceiveTime(time.Time(*input.ReceivedDateTime).UTC().Format(time.RFC3339Nano))
	}
	if input.GNSSPosition != nil {
		output.SetGnssPosition(gnssPosition(input.GNSSPosition))
	}
	if input.WheelBasedSpeed != nil {
		output.SetWheelBasedSpeedKmh(*input.WheelBasedSpeed)
	}
	if input.TachographSpeed != nil {
		output.SetTachographSpeedKmh(*input.TachographSpeed)
	}
	return &output
}
