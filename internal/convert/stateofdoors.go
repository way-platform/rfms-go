package convert

import (
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/rfms/v5"
)

// StateOfDoors converts a rFMS state of doors to proto.
func StateOfDoors(input string) rfmsv5.VehicleStatus_StateOfDoors {
	switch input {
	case "ALL_DOORS_DISABLED":
		return rfmsv5.VehicleStatus_ALL_DOORS_DISABLED
	case "AT_LEAST_ONE_DOOR_ENABLED":
		return rfmsv5.VehicleStatus_AT_LEAST_ONE_DOOR_ENABLED
	case "ERROR":
		return rfmsv5.VehicleStatus_STATE_OF_DOORS_ERROR
	case "NOT_AVAILABLE":
		return rfmsv5.VehicleStatus_STATE_OF_DOORS_NOT_AVAILABLE
	default:
		return rfmsv5.VehicleStatus_STATE_OF_DOORS_UNKNOWN
	}
}
