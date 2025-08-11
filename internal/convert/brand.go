package convert

import (
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/rfms/v5"
)

// Brand converts an rFMS brand to proto.
func Brand(input string) rfmsv5.Brand {
	switch input {
	case "VOLVO TRUCKS":
		return rfmsv5.Brand_VOLVO_TRUCKS
	case "SCANIA":
		return rfmsv5.Brand_SCANIA
	case "DAIMLER":
		return rfmsv5.Brand_DAIMLER
	case "IVECO":
		return rfmsv5.Brand_IVECO
	case "DAF":
		return rfmsv5.Brand_DAF
	case "MAN":
		return rfmsv5.Brand_MAN
	case "RENAULT TRUCKS":
		return rfmsv5.Brand_RENAULT_TRUCKS
	case "VDL":
		return rfmsv5.Brand_VDL
	case "VOLVO BUSES":
		return rfmsv5.Brand_VOLVO_BUSES
	case "IVECO BUS":
		return rfmsv5.Brand_IVECO_BUS
	case "HEULIEZ":
		return rfmsv5.Brand_HEULIEZ
	case "VWTB":
		return rfmsv5.Brand_VWTB
	case "KENWORTH":
		return rfmsv5.Brand_KENWORTH
	case "PETERBILT":
		return rfmsv5.Brand_PETERBILT
	case "MACK TRUCKS":
		return rfmsv5.Brand_MACK_TRUCKS
	case "INTERNATIONAL":
		return rfmsv5.Brand_INTERNATIONAL
	case "IC BUS":
		return rfmsv5.Brand_IC_BUS
	default:
		return rfmsv5.Brand_BRAND_UNKNOWN
	}
}
