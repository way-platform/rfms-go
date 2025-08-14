package convert

import (
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/connect/rfms/v5"
)

// FuelType converts an rFMS fuel type to proto.
func FuelType(input string) rfmsv5.FuelType {
	switch input {
	case "00":
		return rfmsv5.FuelType_FUEL_TYPE_UNSPECIFIED
	case "01":
		return rfmsv5.FuelType_GASOLINE
	case "02":
		return rfmsv5.FuelType_METHANOL
	case "03":
		return rfmsv5.FuelType_ETHANOL
	case "04":
		return rfmsv5.FuelType_DIESEL
	case "05":
		return rfmsv5.FuelType_LIQUEFIED_PETROLEUM_GAS
	case "06":
		return rfmsv5.FuelType_COMPRESSED_NATURAL_GAS
	case "07":
		return rfmsv5.FuelType_PROPANE
	case "08":
		return rfmsv5.FuelType_BATTERY_ELECTRIC
	case "1D":
		return rfmsv5.FuelType_HYDROGEN_FUEL_CELL
	case "1E":
		return rfmsv5.FuelType_HYDROGEN_INTERNAL_COMBUSTION_ENGINE
	case "1F":
		return rfmsv5.FuelType_KEROSENE
	case "20":
		return rfmsv5.FuelType_HEAVY_FUEL_OIL
	default:
		return rfmsv5.FuelType_FUEL_TYPE_UNKNOWN
	}
}
