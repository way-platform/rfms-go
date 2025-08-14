package convert

import (
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/connect/rfms/v5"
)

// DoorLockState converts an rFMS door lock state to proto.
func DoorLockState(input string) rfmsv5.VehicleStatus_Door_LockState {
	switch input {
	case "LOCKED":
		return rfmsv5.VehicleStatus_Door_LOCKED
	case "UNLOCKED":
		return rfmsv5.VehicleStatus_Door_UNLOCKED
	case "ERROR":
		return rfmsv5.VehicleStatus_Door_LOCK_STATE_ERROR
	case "NOT_AVAILABLE":
		return rfmsv5.VehicleStatus_Door_LOCK_STATE_NOT_AVAILABLE
	default:
		return rfmsv5.VehicleStatus_Door_LOCK_STATE_UNKNOWN
	}
}
