package convert

import (
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/connect/rfms/v5"
)

// DoorEnabledState converts an rFMS door enabled state to proto.
func DoorEnabledState(input string) rfmsv5.VehicleStatus_Door_EnabledState {
	switch input {
	case "ENABLED":
		return rfmsv5.VehicleStatus_Door_ENABLED
	case "DISABLED":
		return rfmsv5.VehicleStatus_Door_DISABLED
	case "ERROR":
		return rfmsv5.VehicleStatus_Door_ENABLED_STATE_ERROR
	case "NOT_AVAILABLE":
		return rfmsv5.VehicleStatus_Door_ENABLED_STATE_NOT_AVAILABLE
	default:
		return rfmsv5.VehicleStatus_Door_ENABLED_STATE_UNKNOWN
	}
}
