package convert

import (
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/connect/rfms/v5"
)

// VehicleBodyType converts an rFMS body type to proto.
func VehicleBodyType(input string) rfmsv5.Vehicle_BodyType {
	switch input {
	case "CITY_BUS":
		return rfmsv5.Vehicle_CITY_BUS
	case "INTERCITY_BUS":
		return rfmsv5.Vehicle_INTERCITY_BUS
	case "COACH":
		return rfmsv5.Vehicle_COACH
	default:
		return rfmsv5.Vehicle_BODY_TYPE_UNKNOWN
	}
}
