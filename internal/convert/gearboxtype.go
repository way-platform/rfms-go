package convert

import (
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/rfms/v5"
)

// GearboxType converts an rFMS gearbox type to proto.
func GearboxType(input string) rfmsv5.GearboxType {
	switch input {
	case "MANUAL":
		return rfmsv5.GearboxType_MANUAL
	case "AUTOMATIC":
		return rfmsv5.GearboxType_AUTOMATIC
	case "SEMI_AUTOMATIC":
		return rfmsv5.GearboxType_SEMI_AUTOMATIC
	case "NO_GEAR":
		return rfmsv5.GearboxType_NO_GEAR
	case "AMT":
		return rfmsv5.GearboxType_AMT
	default:
		return rfmsv5.GearboxType_GEARBOX_TYPE_UNKNOWN
	}
}
