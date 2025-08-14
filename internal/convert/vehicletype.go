package convert

import (
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/connect/rfms/v5"
)

// VehicleType converts an rFMS vehicle type to proto.
func VehicleType(input string) rfmsv5.Vehicle_Type {
	switch input {
	case "TRUCK":
		return rfmsv5.Vehicle_TRUCK
	case "BUS":
		return rfmsv5.Vehicle_BUS
	case "VAN":
		return rfmsv5.Vehicle_VAN
	default:
		return rfmsv5.Vehicle_TYPE_UNKNOWN
	}
}
