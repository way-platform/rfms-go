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
func ConvertVehiclePosition(v2 *rfmsv2.VehiclePosition) *rfmsv4.VehiclePosition {
	var v4 rfmsv4.VehiclePosition
	v4.CreatedDateTime = rfmsv4.Time(v2.CreatedDateTime)
	v4.GNSSPosition = convertGNSSPosition(v2.GNSSPosition)
	v4.ReceivedDateTime = rfmsv4.Time(v2.ReceivedDateTime)
	v4.TachographSpeed = v2.TachographSpeed
	v4.TriggerType = *convertTrigger(&v2.Trigger)
	v4.VIN = v2.VIN
	v4.WheelBasedSpeed = v2.WheelBasedSpeed
	return &v4
}

// ConvertVehicleStatus converts an [rfmsv2.VehicleStatus] to an [rfmsv4.VehicleStatus].
func ConvertVehicleStatus(v2 *rfmsv2.VehicleStatus) *rfmsv4.VehicleStatus {
	var v4 rfmsv4.VehicleStatus
	v4.AccumulatedData = convertAccumulatedData(v2.AccumulatedData)
	v4.CreatedDateTime = rfmsv4.Time(v2.CreatedDateTime)
	v4.DoorStatus = convertDoorStatus(v2.DoorStatus)
	v4.Driver1ID = convertDriverID(v2.Driver1ID)
	v4.EngineTotalFuelUsed = v2.EngineTotalFuelUsed
	v4.GrossCombinationVehicleWeight = v2.GrossCombinationVehicleWeight
	v4.HrTotalVehicleDistance = v2.HRTotalVehicleDistance
	v4.ReceivedDateTime = rfmsv4.Time(v2.ReceivedDateTime)
	v4.SnapshotData = convertSnapshotData(v2.SnapshotData)
	v4.Status2OfDoors = convertStatus2OfDoors(v2.Status2OfDoors)
	v4.TotalEngineHours = v2.TotalEngineHours
	v4.TriggerType = *convertTrigger(&v2.Trigger)
	v4.UptimeData = convertUptimeData(v2.UptimeData)
	v4.VIN = v2.VIN
	return &v4
}

func convertUptimeData(v2 *rfmsv2.UptimeData) *rfmsv4.UptimeData {
	var v4 rfmsv4.UptimeData
	v4.AlternatorInfo = convertAlternatorInfo(v2.AlternatorInfo)
	v4.BellowPressureFrontAxleLeft = v2.BellowPressureFrontAxleLeft
	v4.BellowPressureFrontAxleRight = v2.BellowPressureFrontAxleRight
	v4.BellowPressureRearAxleLeft = v2.BellowPressureRearAxleLeft
	v4.BellowPressureRearAxleRight = v2.BellowPressureRearAxleRight
	v4.DurationAtLeastOneDoorOpen = v2.DurationAtLeastOneDoorOpen
	v4.EngineCoolantTemperature = v2.EngineCoolantTemperature
	v4.ServiceBrakeAirPressureCircuit1 = v2.ServiceBrakeAirPressureCircuit1
	v4.ServiceBrakeAirPressureCircuit2 = v2.ServiceBrakeAirPressureCircuit2
	v4.ServiceDistance = v2.ServiceDistance
	v4.TellTaleInfo = convertTellTaleInfo(v2.TellTaleInfo)
	return &v4
}

func convertTellTaleInfo(v2 []rfmsv2.TellTaleInfo) []rfmsv4.TellTaleInfo {
	var result []rfmsv4.TellTaleInfo
	return result
}

func convertAlternatorInfo(v2 []rfmsv2.AlternatorInfo) *rfmsv4.AlternatorInfo {
	var result rfmsv4.AlternatorInfo
	// TODO: Implement me.
	return &result
}

func convertStatus2OfDoors(v2 *rfmsv2.Status2OfDoors) *rfmsv4.Status2OfDoors {
	var v4 rfmsv4.Status2OfDoors
	// TODO: Implement me.
	return &v4
}

func convertSnapshotData(v2 *rfmsv2.SnapshotData) *rfmsv4.SnapshotData {
	var v4 rfmsv4.SnapshotData
	// TODO: Implement me.
	return &v4
}

func convertDriverID(v2 *rfmsv2.DriverID) *rfmsv4.DriverID {
	var v4 rfmsv4.DriverID
	// TODO: Implement me.
	return &v4
}

func convertDoorStatus(v2 []rfmsv2.DoorStatus) []rfmsv4.DoorStatus {
	var v4 []rfmsv4.DoorStatus
	// TODO: Implement me.
	return v4
}

func convertAccumulatedData(v2 *rfmsv2.AccumulatedData) *rfmsv4.AccumulatedData {
	var v4 rfmsv4.AccumulatedData
	// TODO: Implement me.
	return &v4
}

func convertDate(date *rfmsv2.Date) *rfmsv4.Date {
	var result rfmsv4.Date
	result.Day = date.Day
	result.Month = date.Month
	result.Year = date.Year
	return &result
}

func convertGNSSPosition(v2 *rfmsv2.GNSSPosition) *rfmsv4.GNSSPosition {
	var v4 rfmsv4.GNSSPosition
	// TODO: Implement me.
	return &v4
}

func convertTrigger(v2 *rfmsv2.Trigger) *rfmsv4.Trigger {
	var v4 rfmsv4.Trigger
	// TODO: Implement me.
	return &v4
}
