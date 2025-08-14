package convert

import (
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/connect/rfms/v5"
)

// TriggerType converts an rFMS trigger type to proto.
func TriggerType(input string) rfmsv5.Trigger_Type {
	switch input {
	case "TIMER":
		return rfmsv5.Trigger_TIMER
	case "IGNITION_ON":
		return rfmsv5.Trigger_IGNITION_ON
	case "IGNITION_OFF":
		return rfmsv5.Trigger_IGNITION_OFF
	case "PTO_ENABLED":
		return rfmsv5.Trigger_PTO_ENABLED
	case "PTO_DISABLED":
		return rfmsv5.Trigger_PTO_DISABLED
	case "DRIVER_LOGIN":
		return rfmsv5.Trigger_DRIVER_LOGIN
	case "DRIVER_LOGOUT":
		return rfmsv5.Trigger_DRIVER_LOGOUT
	case "TELL_TALE":
		return rfmsv5.Trigger_TELL_TALE
	case "ENGINE_ON":
		return rfmsv5.Trigger_ENGINE_ON
	case "ENGINE_OFF":
		return rfmsv5.Trigger_ENGINE_OFF
	case "DRIVER_1_WORKING_STATE_CHANGED":
		return rfmsv5.Trigger_DRIVER_1_WORKING_STATE_CHANGED
	case "DRIVER_2_WORKING_STATE_CHANGED":
		return rfmsv5.Trigger_DRIVER_2_WORKING_STATE_CHANGED
	case "DISTANCE_TRAVELLED":
		return rfmsv5.Trigger_DISTANCE_TRAVELLED
	case "FUEL_TYPE_CHANGE":
		return rfmsv5.Trigger_FUEL_TYPE_CHANGE
	case "PARKING_BRAKE_SWITCH_CHANGE":
		return rfmsv5.Trigger_PARKING_BRAKE_SWITCH_CHANGE
	case "BATTERY_PACK_CHARGING_STATUS_CHANGE":
		return rfmsv5.Trigger_BATTERY_PACK_CHARGING_STATUS_CHANGE
	case "BATTERY_PACK_CHARGING_CONNECTION_STATUS_CHANGE":
		return rfmsv5.Trigger_BATTERY_PACK_CHARGING_CONNECTION_STATUS_CHANGE
	case "TRAILER_CONNECTED":
		return rfmsv5.Trigger_TRAILER_CONNECTED
	case "TRAILER_DISCONNECTED":
		return rfmsv5.Trigger_TRAILER_DISCONNECTED
	case "ALARM":
		return rfmsv5.Trigger_ALARM
	default:
		return rfmsv5.Trigger_TYPE_UNKNOWN
	}
}
