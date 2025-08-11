package convert

import (
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/rfms/v5"
)

// TachographType converts an rFMS tachograph type to proto.
func TachographType(input string) rfmsv5.TachographType {
	switch input {
	case "MTCO":
		return rfmsv5.TachographType_MTCO
	case "DTCO":
		return rfmsv5.TachographType_DTCO
	case "DTCO_G1":
		return rfmsv5.TachographType_DTCO_G1
	case "DTCO_G2":
		return rfmsv5.TachographType_DTCO_G2
	case "TSU":
		return rfmsv5.TachographType_TSU
	case "NONE":
		return rfmsv5.TachographType_NONE
	case "smart STONERIDGE": // Scania
		return rfmsv5.TachographType_STONERIDGE_SMART
	case "smart 2 STONERIDGE": // Scania
		return rfmsv5.TachographType_STONERIDGE_SMART2
	default:
		return rfmsv5.TachographType_TACHOGRAPH_TYPE_UNKNOWN
	}
}
