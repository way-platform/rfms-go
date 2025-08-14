package convert

import (
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/connect/rfms/v5"
)

// ChargingDevice converts an rFMS charging device to proto.
func ChargingDevice(input string) rfmsv5.ChargingDevice {
	switch input {
	case "NONE":
		return rfmsv5.ChargingDevice_CHARGING_DEVICE_NONE
	case "ACD":
		return rfmsv5.ChargingDevice_AUTOMATIC_CONNECTION_DEVICE
	case "WPT":
		return rfmsv5.ChargingDevice_WIRELESS_POWER_TRANSFER
	case "VEHICLE_COUPLER":
		return rfmsv5.ChargingDevice_VEHICLE_COUPLER
	case "NOT_AVAILABLE":
		return rfmsv5.ChargingDevice_CHARGING_DEVICE_NOT_AVAILABLE
	default:
		return rfmsv5.ChargingDevice_CHARGING_DEVICE_UNKNOWN
	}
}
