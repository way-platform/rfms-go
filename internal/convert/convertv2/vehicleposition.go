package convertv2

import (
	"time"

	"github.com/way-platform/rfms-go/internal/openapi/rfmsv2oapi"
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/connect/rfms/v5"
	"google.golang.org/protobuf/types/known/timestamppb"
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
		output.SetCreateTime(timestamppb.New(time.Time(*input.CreatedDateTime).UTC()))
	}
	if input.ReceivedDateTime != nil {
		output.SetReceiveTime(timestamppb.New(time.Time(*input.ReceivedDateTime).UTC()))
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
