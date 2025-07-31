package convert

import (
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/rfms/v5"
)

// VehicleTachographType converts an rFMS tachograph type to proto.
func VehicleTachographType(input string) rfmsv5.Vehicle_TachographType {
	switch input {
	case "MTCO":
		return rfmsv5.Vehicle_MTCO
	case "DTCO":
		return rfmsv5.Vehicle_DTCO
	case "DTCO_G1":
		return rfmsv5.Vehicle_DTCO_G1
	case "DTCO_G2":
		return rfmsv5.Vehicle_DTCO_G2
	case "TSU":
		return rfmsv5.Vehicle_TSU
	case "NONE":
		return rfmsv5.Vehicle_NONE
	case "smart STONERIDGE": // Scania
		return rfmsv5.Vehicle_STONERIDGE_SMART
	case "smart 2 STONERIDGE": // Scania
		return rfmsv5.Vehicle_STONERIDGE_SMART2
	default:
		return rfmsv5.Vehicle_TACHOGRAPH_TYPE_UNKNOWN
	}
}
