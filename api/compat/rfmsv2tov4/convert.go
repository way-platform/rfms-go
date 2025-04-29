package rfmsv2tov4

import (
	"github.com/way-platform/rfms-go/v4/api/rfmsv2"
	"github.com/way-platform/rfms-go/v4/api/rfmsv4"
)

// ConvertVehicle converts an [rfmsv2.Vehicle] to an [rfmsv4.Vehicle].
func ConvertVehicle(v2 *rfmsv2.Vehicle) *rfmsv4.Vehicle {
	var v4 rfmsv4.Vehicle
	v4.BodyType = v2.BodyType
	v4.Brand = v2.Brand
	v4.ChassisType = v2.ChassisType
	v4.CustomerVehicleName = v2.CustomerVehicleName
	v4.DoorConfiguration = v2.DoorConfiguration
	v4.EmissionLevel = v2.EmissionLevel
	v4.GearboxType = v2.GearboxType
	v4.HasRampOrLift = v2.HasRampOrLift
	v4.Model = v2.Model
	v4.NoOfAxles = v2.NoOfAxles
	v4.PossibleFuelType = v2.PossibleFuelType
	v4.ProductionDate = convertDate(v2.ProductionDate)
	v4.TachographType = v2.TachographType
	v4.TellTaleCode = v2.TellTaleCode
	v4.TotalFuelTankVolume = v2.TotalFuelTankVolume
	v4.Type = v2.Type
	v4.VIN = v2.VIN
	return &v4
}

// ConvertVehiclePosition converts an [rfmsv2.VehiclePosition] to an [rfmsv4.VehiclePosition].
func ConvertVehiclePosition(position *rfmsv2.VehiclePosition) *rfmsv4.VehiclePosition {
	var result rfmsv4.VehiclePosition
	// TODO: Implement me.
	return &result
}

// ConvertVehicleStatus converts an [rfmsv2.VehicleStatus] to an [rfmsv4.VehicleStatus].
func ConvertVehicleStatus(status *rfmsv2.VehicleStatus) *rfmsv4.VehicleStatus {
	var result rfmsv4.VehicleStatus
	// TODO: Implement me.
	return &result
}

func convertDate(date *rfmsv2.Date) *rfmsv4.Date {
	var result rfmsv4.Date
	result.Day = date.Day
	result.Month = date.Month
	result.Year = date.Year
	return &result
}
