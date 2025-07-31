package convert

import (
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/rfms/v5"
)

// ChargingState converts an rFMS charging state to proto.
func ChargingState(input string) rfmsv5.ChargingState {
	switch input {
	case "NOT_CHARGING":
		return rfmsv5.ChargingState_NOT_CHARGING
	case "CHARGING":
		return rfmsv5.ChargingState_CHARGING
	case "CHARGING_AC":
		return rfmsv5.ChargingState_CHARGING_AC
	case "CHARGING_DC":
		return rfmsv5.ChargingState_CHARGING_DC
	case "ERROR":
		return rfmsv5.ChargingState_CHARGING_STATE_ERROR
	case "NOT_AVAILABLE":
		return rfmsv5.ChargingState_CHARGING_STATE_NOT_AVAILABLE
	default:
		return rfmsv5.ChargingState_CHARGING_STATE_UNKNOWN
	}
}
