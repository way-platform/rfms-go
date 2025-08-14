package convert

import (
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/connect/rfms/v5"
)

// IgnitionState converts an rFMS ignition state to proto.
func IgnitionState(input string) rfmsv5.IgnitionState {
	switch input {
	case "OFF":
		return rfmsv5.IgnitionState_OFF
	case "ON":
		return rfmsv5.IgnitionState_ON
	case "NOT_AVAILABLE":
		return rfmsv5.IgnitionState_IGNITION_STATE_NOT_AVAILABLE
	default:
		return rfmsv5.IgnitionState_IGNITION_STATE_UNKNOWN
	}
}
