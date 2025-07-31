package convertv2

import (
	"github.com/way-platform/rfms-go/internal/openapi/rfmsv2oapi"
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/rfms/v5"
)

// VehiclePosition converts an rFMS v2 vehicle position to proto.
func VehiclePosition(input *rfmsv2oapi.VehiclePositionType) *rfmsv5.VehiclePosition {
	var output rfmsv5.VehiclePosition
	if input.VIN != nil {
		output.SetVin(*input.VIN)
	}
	if input.TriggerType != nil {
		output.SetTrigger(trigger(input.TriggerType))
	}
	if input.CreatedDateTime != nil {
		output.SetCreateTime(input.CreatedDateTime.UnixMicro())
	}
	if input.ReceivedDateTime != nil {
		output.SetReceiveTime(input.ReceivedDateTime.UnixMicro())
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
