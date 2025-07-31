package convertv2

import (
	"github.com/way-platform/rfms-go/internal/convert"
	"github.com/way-platform/rfms-go/internal/openapi/rfmsv2oapi"
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/rfms/v5"
)

// Vehicle converts an rFMS v2 vehicle to proto.
func Vehicle(input *rfmsv2oapi.VehicleType) *rfmsv5.Vehicle {
	var output rfmsv5.Vehicle
	if input.VIN != nil {
		output.SetVin(*input.VIN)
	}
	if input.BodyType != nil {
		output.SetBodyType(convert.VehicleBodyType(*input.BodyType))
		if output.GetBodyType() == rfmsv5.Vehicle_BODY_TYPE_UNKNOWN {
			output.SetUnknownBodyType(string(*input.BodyType))
		}
	}
	if input.Brand != nil {
		output.SetBrand(convert.VehicleBrand(*input.Brand))
		if output.GetBrand() == rfmsv5.Vehicle_BRAND_UNKNOWN {
			output.SetUnknownBrand(*input.Brand)
		}
	}
	if input.ChassisType != nil {
		output.SetChassisType(*input.ChassisType)
	}
	if input.CustomerVehicleName != nil {
		output.SetCustomerVehicleName(*input.CustomerVehicleName)
	}
	if len(input.DoorConfiguration) > 0 {
		for _, door := range input.DoorConfiguration {
			output.SetDoorConfiguration(append(output.GetDoorConfiguration(), door))
		}
	}
	if input.EmissionLevel != nil {
		output.SetEmissionLevel(convert.VehicleEmissionLevel(*input.EmissionLevel))
		if output.GetEmissionLevel() == rfmsv5.Vehicle_EMISSION_LEVEL_UNKNOWN {
			output.SetUnknownEmissionLevel(*input.EmissionLevel)
		}
	}
	if input.GearboxType != nil {
		output.SetGearboxType(convert.VehicleGearboxType(*input.GearboxType))
		if output.GetGearboxType() == rfmsv5.Vehicle_GEARBOX_TYPE_UNKNOWN {
			output.SetUnknownGearboxType(string(*input.GearboxType))
		}
	}
	if input.HasRampOrLift != nil {
		output.SetHasRampOrLift(*input.HasRampOrLift)
	}
	if input.Model != nil {
		output.SetModel(*input.Model)
	}
	if input.NoOfAxles != nil {
		output.SetAxleCount(*input.NoOfAxles)
	}
	if len(input.PossibleFuelType) > 0 {
		for _, possibleFuelType := range input.PossibleFuelType {
			fuelType := convert.FuelType(possibleFuelType)
			output.SetPossibleFuelTypes(append(output.GetPossibleFuelTypes(), fuelType))
			if fuelType == rfmsv5.FuelType_FUEL_TYPE_UNKNOWN {
				output.SetUnknownPossibleFuelTypes(append(output.GetUnknownPossibleFuelTypes(), fuelType))
			}
		}
	}
	if input.ProductionDate != nil {
		output.SetProductionDate(convert.Date(input.ProductionDate.Year, input.ProductionDate.Month, input.ProductionDate.Day))
	}
	if input.TachographType != nil {
		output.SetTachographType(convert.VehicleTachographType(*input.TachographType))
		if output.GetTachographType() == rfmsv5.Vehicle_TACHOGRAPH_TYPE_UNKNOWN {
			output.SetUnknownTachographType(string(*input.TachographType))
		}
	}
	if input.TellTaleCode != nil {
		output.SetTellTaleCode(*input.TellTaleCode)
	}
	if input.TotalFuelTankVolume != nil {
		output.SetTotalFuelTankVolumeMl(float64(*input.TotalFuelTankVolume))
	}
	if input.Type != nil {
		output.SetType(convert.VehicleType(*input.Type))
		if output.GetType() == rfmsv5.Vehicle_TYPE_UNKNOWN {
			output.SetUnknownType(string(*input.Type))
		}
	}
	return &output
}
