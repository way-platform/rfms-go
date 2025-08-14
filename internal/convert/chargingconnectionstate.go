package convert

import (
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/connect/rfms/v5"
)

// ChargingConnectionState converts an rFMS charging connection state to proto.
func ChargingConnectionState(input string) rfmsv5.ChargingConnectionState {
	switch input {
	case "CONNECTING":
		return rfmsv5.ChargingConnectionState_CONNECTING
	case "CONNECTED":
		return rfmsv5.ChargingConnectionState_CONNECTED
	case "DISCONNECTING":
		return rfmsv5.ChargingConnectionState_DISCONNECTING
	case "DISCONNECTED":
		return rfmsv5.ChargingConnectionState_DISCONNECTED
	case "ERROR":
		return rfmsv5.ChargingConnectionState_CHARGING_CONNECTION_STATE_ERROR
	case "NOT_AVAILABLE":
		return rfmsv5.ChargingConnectionState_CHARGING_CONNECTION_STATE_NOT_AVAILABLE
	default:
		return rfmsv5.ChargingConnectionState_CHARGING_CONNECTION_STATE_UNKNOWN
	}
}
