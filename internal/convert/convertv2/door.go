package convertv2

import (
	"github.com/way-platform/rfms-go/internal/convert"
	rfmsv2oapi "github.com/way-platform/rfms-go/internal/openapi/rfmsv2oapi"
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/connect/rfms/v5"
)

// door converts an rFMS v2 door status to proto.
func door(input *rfmsv2oapi.DoorStatusType) *rfmsv5.VehicleStatus_Door {
	var output rfmsv5.VehicleStatus_Door
	if input.DoorNumber != nil {
		output.SetNumber(int32(*input.DoorNumber))
	}
	if input.DoorEnabledStatus != nil {
		output.SetEnabledState(convert.DoorEnabledState(string(*input.DoorEnabledStatus)))
		if output.GetEnabledState() == rfmsv5.VehicleStatus_Door_ENABLED_STATE_UNKNOWN {
			output.SetUnknownEnabledState(string(*input.DoorEnabledStatus))
		}
	}
	if input.DoorOpenStatus != nil {
		output.SetOpenState(convert.DoorOpenState(string(*input.DoorOpenStatus)))
		if output.GetOpenState() == rfmsv5.VehicleStatus_Door_OPEN_STATE_UNKNOWN {
			output.SetUnknownOpenState(string(*input.DoorOpenStatus))
		}
	}
	if input.DoorLockStatus != nil {
		output.SetLockState(convert.DoorLockState(string(*input.DoorLockStatus)))
		if output.GetLockState() == rfmsv5.VehicleStatus_Door_LOCK_STATE_UNKNOWN {
			output.SetUnknownLockState(string(*input.DoorLockStatus))
		}
	}
	return &output
}
