package convert

import (
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/rfms/v5"
)

// VehicleGearboxType converts an rFMS gearbox type to proto.
func VehicleGearboxType(input string) rfmsv5.Vehicle_GearboxType {
	switch input {
	case "MANUAL":
		return rfmsv5.Vehicle_MANUAL
	case "AUTOMATIC":
		return rfmsv5.Vehicle_AUTOMATIC
	case "SEMI_AUTOMATIC":
		return rfmsv5.Vehicle_SEMI_AUTOMATIC
	case "NO_GEAR":
		return rfmsv5.Vehicle_NO_GEAR
	case "AMT":
		return rfmsv5.Vehicle_AMT
	default:
		return rfmsv5.Vehicle_GEARBOX_TYPE_UNKNOWN
	}
}
