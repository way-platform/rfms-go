package convert

import (
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/connect/rfms/v5"
)

// AlternatorState converts an alternator state to proto.
func AlternatorState(input string) rfmsv5.UptimeData_Alternator_State {
	switch input {
	case "NOT_CHARGING":
		return rfmsv5.UptimeData_Alternator_NOT_CHARGING
	case "CHARGING":
		return rfmsv5.UptimeData_Alternator_CHARGING
	case "ERROR":
		return rfmsv5.UptimeData_Alternator_STATE_ERROR
	case "NOT_AVAILABLE":
		return rfmsv5.UptimeData_Alternator_STATE_NOT_AVAILABLE
	default:
		return rfmsv5.UptimeData_Alternator_STATE_UNKNOWN
	}
}
