package convert

import (
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/rfms/v5"
)

// AuthenticationEquipment converts an rFMS driver authentication equipment to proto.
func AuthenticationEquipment(input string) rfmsv5.DriverIdentification_Tacho_AuthenticationEquipment {
	switch input {
	case "DRIVER_CARD":
		return rfmsv5.DriverIdentification_Tacho_DRIVER_CARD
	case "CONTROL_CARD":
		return rfmsv5.DriverIdentification_Tacho_CONTROL_CARD
	case "COMPANY_CARD":
		return rfmsv5.DriverIdentification_Tacho_COMPANY_CARD
	case "MANUFACTURING_CARD":
		return rfmsv5.DriverIdentification_Tacho_MANUFACTURING_CARD
	case "VEHICLE_UNIT":
		return rfmsv5.DriverIdentification_Tacho_VEHICLE_UNIT
	case "MOTION_SENSOR":
		return rfmsv5.DriverIdentification_Tacho_MOTION_SENSOR
	default:
		return rfmsv5.DriverIdentification_Tacho_AUTHENTICATION_EQUIPMENT_UNKNOWN
	}
}
