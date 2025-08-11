package convert

import (
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/rfms/v5"
)

// DoorOpenState converts an rFMS bus door open state to proto.
func DoorOpenState(input string) rfmsv5.VehicleStatus_Door_OpenState {
	switch input {
	case "OPEN":
		return rfmsv5.VehicleStatus_Door_OPEN
	case "CLOSED":
		return rfmsv5.VehicleStatus_Door_CLOSED
	case "ERROR":
		return rfmsv5.VehicleStatus_Door_OPEN_STATE_ERROR
	case "NOT_AVAILABLE":
		return rfmsv5.VehicleStatus_Door_OPEN_STATE_NOT_AVAILABLE
	default:
		return rfmsv5.VehicleStatus_Door_OPEN_STATE_UNKNOWN
	}
}
