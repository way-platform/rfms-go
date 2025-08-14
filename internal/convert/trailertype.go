package convert

import (
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/connect/rfms/v5"
)

// TrailerType converts an rFMS trailer type to a Trailer_Type enum.
func TrailerType(input string) rfmsv5.Trailer_Type {
	switch input {
	case "SEMI_TRAILER":
		return rfmsv5.Trailer_SEMI_TRAILER
	case "CENTRE_AXLE_TRAILER":
		return rfmsv5.Trailer_CENTRE_AXLE_TRAILER
	case "FULL_TRAILER":
		return rfmsv5.Trailer_FULL_TRAILER
	case "CONVERTER_DOLLY":
		return rfmsv5.Trailer_CONVERTER_DOLLY
	case "LINK_TRAILER":
		return rfmsv5.Trailer_LINK_TRAILER
	case "TOWING_SEMI_TRAILER":
		return rfmsv5.Trailer_TOWING_SEMI_TRAILER
	case "TOWING_CENTRE_AXLE_TRAILER":
		return rfmsv5.Trailer_TOWING_CENTRE_AXLE_TRAILER
	case "TOWING_FULL_TRAILER":
		return rfmsv5.Trailer_TOWING_FULL_TRAILER
	case "UNKNOWN":
		return rfmsv5.Trailer_TYPE_NOT_AVAILABLE
	default:
		return rfmsv5.Trailer_TYPE_UNKNOWN
	}
}
