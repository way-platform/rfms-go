package convert

import (
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/rfms/v5"
)

// VehicleBrand converts an rFMS brand to proto.
func VehicleBrand(input string) rfmsv5.Vehicle_Brand {
	switch input {
	case "VOLVO TRUCKS":
		return rfmsv5.Vehicle_VOLVO_TRUCKS
	case "SCANIA":
		return rfmsv5.Vehicle_SCANIA
	case "DAIMLER":
		return rfmsv5.Vehicle_DAIMLER
	case "IVECO":
		return rfmsv5.Vehicle_IVECO
	case "DAF":
		return rfmsv5.Vehicle_DAF
	case "MAN":
		return rfmsv5.Vehicle_MAN
	case "RENAULT TRUCKS":
		return rfmsv5.Vehicle_RENAULT_TRUCKS
	case "VDL":
		return rfmsv5.Vehicle_VDL
	case "VOLVO BUSES":
		return rfmsv5.Vehicle_VOLVO_BUSES
	case "IVECO BUS":
		return rfmsv5.Vehicle_IVECO_BUS
	case "HEULIEZ":
		return rfmsv5.Vehicle_HEULIEZ
	case "VWTB":
		return rfmsv5.Vehicle_VWTB
	case "KENWORTH":
		return rfmsv5.Vehicle_KENWORTH
	case "PETERBILT":
		return rfmsv5.Vehicle_PETERBILT
	case "MACK TRUCKS":
		return rfmsv5.Vehicle_MACK_TRUCKS
	case "INTERNATIONAL":
		return rfmsv5.Vehicle_INTERNATIONAL
	case "IC BUS":
		return rfmsv5.Vehicle_IC_BUS
	default:
		return rfmsv5.Vehicle_BRAND_UNKNOWN
	}
}
