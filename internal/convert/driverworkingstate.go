package convert

import (
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/connect/rfms/v5"
)

// DriverWorkingState converts an rFMS driver working state to proto.
func DriverWorkingState(input string) rfmsv5.DriverWorkingState {
	switch input {
	case "REST":
		return rfmsv5.DriverWorkingState_REST
	case "AVAILABLE":
		return rfmsv5.DriverWorkingState_AVAILABLE
	case "DRIVER_AVAILABLE":
		return rfmsv5.DriverWorkingState_AVAILABLE
	case "WORK":
		return rfmsv5.DriverWorkingState_WORK
	case "DRIVE":
		return rfmsv5.DriverWorkingState_DRIVE
	case "ERROR":
		return rfmsv5.DriverWorkingState_DRIVER_WORKING_STATE_ERROR
	case "NOT_AVAILABLE":
		return rfmsv5.DriverWorkingState_DRIVER_WORKING_STATE_NOT_AVAILABLE
	default:
		return rfmsv5.DriverWorkingState_DRIVER_WORKING_STATE_UNKNOWN
	}
}
