package convert

import (
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/rfms/v5"
)

// VehicleEmissionLevel converts an rFMS emission level to proto.
func VehicleEmissionLevel(input string) rfmsv5.Vehicle_EmissionLevel {
	switch input {
	case "EURO_III":
		return rfmsv5.Vehicle_EURO_III
	case "EURO_III_EEV":
		return rfmsv5.Vehicle_EURO_III_EEV
	case "EURO_IV":
		return rfmsv5.Vehicle_EURO_IV
	case "EURO_V":
		return rfmsv5.Vehicle_EURO_V
	case "EURO_VI":
		return rfmsv5.Vehicle_EURO_VI
	case "EURO_VII":
		return rfmsv5.Vehicle_EURO_VII
	case "EURO_STAGE_III":
		return rfmsv5.Vehicle_EURO_STAGE_III
	case "EURO_STAGE_IV":
		return rfmsv5.Vehicle_EURO_STAGE_IV
	case "EURO_STAGE_V":
		return rfmsv5.Vehicle_EURO_STAGE_V
	case "EPA_2004":
		return rfmsv5.Vehicle_EPA_2004
	case "EPA_2007":
		return rfmsv5.Vehicle_EPA_2007
	case "EPA_2010":
		return rfmsv5.Vehicle_EPA_2010
	case "EPA_2015_NOX10":
		return rfmsv5.Vehicle_EPA_2015_NOX10
	case "EPA_2015_NOX05":
		return rfmsv5.Vehicle_EPA_2015_NOX05
	case "EPA_2015_NOX02":
		return rfmsv5.Vehicle_EPA_2015_NOX02
	case "EPA_TIER_2":
		return rfmsv5.Vehicle_EPA_TIER_2
	case "EPA_TIER_3":
		return rfmsv5.Vehicle_EPA_TIER_3
	case "EPA_TIER_4_2008":
		return rfmsv5.Vehicle_EPA_TIER_4_2008
	case "EPA_TIER_4_2013":
		return rfmsv5.Vehicle_EPA_TIER_4_2013
	case "PROCONVE_P5":
		return rfmsv5.Vehicle_PROCONVE_P5
	case "PROCONVE_P6":
		return rfmsv5.Vehicle_PROCONVE_P6
	case "PROCONVE_P7":
		return rfmsv5.Vehicle_PROCONVE_P7
	case "PROCONVE_MARI":
		return rfmsv5.Vehicle_PROCONVE_MARI
	default:
		return rfmsv5.Vehicle_EMISSION_LEVEL_UNKNOWN
	}
}
