package rfmsv2tov4

import (
	"github.com/way-platform/rfms-go/api/rfmsv2"
	"github.com/way-platform/rfms-go/api/rfmsv4"
)

// ConvertVehiclesResponse converts an [rfmsv2.VehiclesResponse] to an [rfmsv4.VehiclesResponse].
func ConvertVehiclesResponse(v2 *rfmsv2.VehiclesResponse) *rfmsv4.VehiclesResponse {
	var v4 rfmsv4.VehiclesResponse
	v4.MoreDataAvailable = v2.MoreDataAvailable
	v4.VehicleResponse.Vehicles = make([]rfmsv4.Vehicle, 0, len(v2.Vehicle))
	for _, v2Vehicle := range v2.Vehicle {
		v4.VehicleResponse.Vehicles = append(v4.VehicleResponse.Vehicles, *ConvertVehicle(&v2Vehicle))
	}
	return &v4
}

// ConvertVehiclePositionsResponse converts an [rfmsv2.VehiclePositionsResponse] to an [rfmsv4.VehiclePositionsResponse].
func ConvertVehiclePositionsResponse(v2 *rfmsv2.VehiclePositionsResponse) *rfmsv4.VehiclePositionsResponse {
	var v4 rfmsv4.VehiclePositionsResponse
	v4.MoreDataAvailable = v2.MoreDataAvailable
	v4.RequestServerDateTime = rfmsv4.Time(v2.RequestServerDateTime)
	v4.VehiclePositionResponse.VehiclePositions = make([]rfmsv4.VehiclePosition, 0, len(v2.VehiclePosition))
	for _, v2VehiclePosition := range v2.VehiclePosition {
		v4.VehiclePositionResponse.VehiclePositions = append(v4.VehiclePositionResponse.VehiclePositions, *ConvertVehiclePosition(&v2VehiclePosition))
	}
	return &v4
}

// ConvertVehicleStatusesResponse converts an [rfmsv2.VehicleStatusesResponse] to an [rfmsv4.VehicleStatusesResponse].
func ConvertVehicleStatusesResponse(v2 *rfmsv2.VehicleStatusesResponse) *rfmsv4.VehicleStatusesResponse {
	var v4 rfmsv4.VehicleStatusesResponse
	v4.MoreDataAvailable = v2.MoreDataAvailable
	v4.RequestServerDateTime = rfmsv4.Time(v2.RequestServerDateTime)
	v4.VehicleStatusResponse.VehicleStatuses = make([]rfmsv4.VehicleStatus, 0, len(v2.VehicleStatus))
	for _, v2VehicleStatus := range v2.VehicleStatus {
		v4.VehicleStatusResponse.VehicleStatuses = append(v4.VehicleStatusResponse.VehicleStatuses, *ConvertVehicleStatus(&v2VehicleStatus))
	}
	return &v4
}

// ConvertVehicle converts an [rfmsv2.Vehicle] to an [rfmsv4.Vehicle].
func ConvertVehicle(v2 *rfmsv2.Vehicle) *rfmsv4.Vehicle {
	if v2 == nil {
		return nil
	}
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
	if v2 == nil {
		return nil
	}
	var v4 rfmsv4.VehiclePosition
	v4.CreatedDateTime = rfmsv4.Time(v2.CreatedDateTime)
	v4.GNSSPosition = convertGNSSPosition(v2.GNSSPosition)
	v4.ReceivedDateTime = rfmsv4.Time(v2.ReceivedDateTime)
	v4.TachographSpeed = v2.TachographSpeed
	v4.TriggerType = convertTrigger(v2.TriggerType)
	v4.VIN = v2.VIN
	v4.WheelBasedSpeed = v2.WheelBasedSpeed
	return &v4
}

// ConvertVehicleStatus converts an [rfmsv2.VehicleStatus] to an [rfmsv4.VehicleStatus].
func ConvertVehicleStatus(v2 *rfmsv2.VehicleStatus) *rfmsv4.VehicleStatus {
	if v2 == nil {
		return nil
	}
	var v4 rfmsv4.VehicleStatus
	v4.AccumulatedData = convertAccumulatedData(v2.AccumulatedData)
	v4.CreatedDateTime = rfmsv4.Time(v2.CreatedDateTime)
	v4.DoorStatus = make([]rfmsv4.DoorStatus, 0, len(v2.DoorStatus))
	for _, v2DoorStatus := range v2.DoorStatus {
		v4.DoorStatus = append(v4.DoorStatus, *convertDoorStatus(&v2DoorStatus))
	}
	v4.Driver1ID = convertDriverID(v2.Driver1ID)
	v4.EngineTotalFuelUsed = v2.EngineTotalFuelUsed
	v4.GrossCombinationVehicleWeight = v2.GrossCombinationVehicleWeight
	v4.HrTotalVehicleDistance = v2.HRTotalVehicleDistance
	v4.ReceivedDateTime = rfmsv4.Time(v2.ReceivedDateTime)
	v4.SnapshotData = convertSnapshotData(v2.SnapshotData)
	v4.Status2OfDoors = rfmsv4.Status2OfDoors(v2.Status2OfDoors)
	v4.TotalEngineHours = v2.TotalEngineHours
	v4.TriggerType = convertTrigger(v2.TriggerType)
	v4.UptimeData = convertUptimeData(v2.UptimeData)
	v4.VIN = v2.VIN
	return &v4
}

func convertUptimeData(v2 *rfmsv2.UptimeData) *rfmsv4.UptimeData {
	if v2 == nil {
		return nil
	}
	var v4 rfmsv4.UptimeData
	v4.AlternatorInfo = make([]rfmsv4.AlternatorInfo, 0, len(v2.AlternatorInfo))
	for _, v2AlternatorInfo := range v2.AlternatorInfo {
		v4.AlternatorInfo = append(v4.AlternatorInfo, *convertAlternatorInfo(&v2AlternatorInfo))
	}
	v4.BellowPressureFrontAxleLeft = v2.BellowPressureFrontAxleLeft
	v4.BellowPressureFrontAxleRight = v2.BellowPressureFrontAxleRight
	v4.BellowPressureRearAxleLeft = v2.BellowPressureRearAxleLeft
	v4.BellowPressureRearAxleRight = v2.BellowPressureRearAxleRight
	v4.DurationAtLeastOneDoorOpen = v2.DurationAtLeastOneDoorOpen
	v4.EngineCoolantTemperature = v2.EngineCoolantTemperature
	v4.ServiceBrakeAirPressureCircuit1 = v2.ServiceBrakeAirPressureCircuit1
	v4.ServiceBrakeAirPressureCircuit2 = v2.ServiceBrakeAirPressureCircuit2
	v4.ServiceDistance = v2.ServiceDistance
	v4.TellTaleInfo = make([]rfmsv4.TellTaleInfo, 0, len(v2.TellTaleInfo))
	for _, v2TellTaleInfo := range v2.TellTaleInfo {
		v4.TellTaleInfo = append(v4.TellTaleInfo, *convertTellTaleInfo(&v2TellTaleInfo))
	}
	return &v4
}

func convertTellTaleInfo(v2 *rfmsv2.TellTaleInfo) *rfmsv4.TellTaleInfo {
	if v2 == nil {
		return nil
	}
	var result rfmsv4.TellTaleInfo
	result.TellTale = rfmsv4.TellTaleType(v2.TellTale)
	result.State = rfmsv4.TellTaleState(v2.State)
	return &result
}

func convertAlternatorInfo(v2 *rfmsv2.AlternatorInfo) *rfmsv4.AlternatorInfo {
	if v2 == nil {
		return nil
	}
	var v4 rfmsv4.AlternatorInfo
	v4.AlternatorNumber = v2.AlternatorNumber
	v4.AlternatorStatus = rfmsv4.AlternatorStatus(v2.AlternatorStatus)
	return &v4
}

func convertSnapshotData(v2 *rfmsv2.SnapshotData) *rfmsv4.SnapshotData {
	if v2 == nil {
		return nil
	}
	var v4 rfmsv4.SnapshotData
	v4.AmbientAirTemperature = v2.AmbientAirTemperature
	v4.CatalystFuelLevel = v2.CatalystFuelLevel
	v4.Driver1WorkingState = rfmsv4.DriverWorkingState(v2.Driver1WorkingState)
	v4.Driver2ID = convertDriverID(v2.Driver2ID)
	v4.Driver2WorkingState = rfmsv4.DriverWorkingState(v2.Driver2WorkingState)
	v4.EngineSpeed = v2.EngineSpeed
	v4.FuelLevel1 = v2.FuelLevel1
	v4.GNSSPosition = convertGNSSPosition(v2.GNSSPosition)
	v4.TachographSpeed = v2.TachographSpeed
	v4.WheelBasedSpeed = v2.WheelBasedSpeed
	return &v4
}

func convertDriverID(v2 *rfmsv2.DriverID) *rfmsv4.DriverID {
	if v2 == nil {
		return nil
	}
	var v4 rfmsv4.DriverID
	v4.OEMDriverIdentification = convertOEMDriverIdentification(v2.OEMDriverIdentification)
	v4.TachoDriverIdentification = convertTachoDriverIdentification(v2.TachoDriverIdentification)
	return &v4
}

func convertOEMDriverIdentification(v2 *rfmsv2.OEMDriverIdentification) *rfmsv4.OEMDriverIdentification {
	if v2 == nil {
		return nil
	}
	var v4 rfmsv4.OEMDriverIdentification
	v4.IDType = v2.IDType
	v4.OEMDriverIdentification = v2.OEMDriverIdentification
	return &v4
}

func convertTachoDriverIdentification(v2 *rfmsv2.TachoDriverIdentification) *rfmsv4.TachoDriverIdentification {
	if v2 == nil {
		return nil
	}
	var v4 rfmsv4.TachoDriverIdentification
	v4.CardIssuingMemberState = v2.CardIssuingMemberState
	v4.CardRenewalIndex = v2.CardRenewalIndex
	v4.CardReplacementIndex = v2.CardReplacementIndex
	v4.DriverAuthenticationEquipment = rfmsv4.DriverAuthenticationType(v2.DriverAuthenticationEquipment)
	v4.DriverIdentification = v2.DriverIdentification
	return &v4
}

func convertDoorStatus(v2 *rfmsv2.DoorStatus) *rfmsv4.DoorStatus {
	if v2 == nil {
		return nil
	}
	var v4 rfmsv4.DoorStatus
	v4.DoorEnabledStatus = rfmsv4.DoorEnabledStatus(v2.DoorEnabledStatus)
	v4.DoorLockStatus = rfmsv4.DoorLockStatus(v2.DoorLockStatus)
	v4.DoorNumber = v2.DoorNumber
	v4.DoorOpenStatus = rfmsv4.DoorOpenStatus(v2.DoorOpenStatus)
	return &v4
}

func convertAccumulatedData(v2 *rfmsv2.AccumulatedData) *rfmsv4.AccumulatedData {
	if v2 == nil {
		return nil
	}
	var v4 rfmsv4.AccumulatedData
	v4.AccelerationClass = make([]rfmsv4.FromToClass, 0, len(v2.AccelerationClass.Value))
	for _, v2FromToClass := range v2.AccelerationClass.Value {
		v4.AccelerationClass = append(v4.AccelerationClass, *convertFromToClass(&v2FromToClass))
	}
	v4.AccelerationDuringBrakeClass = make([]rfmsv4.FromToClass, 0, len(v2.AccelerationDuringBrakeClass.Value))
	for _, v2FromToClass := range v2.AccelerationDuringBrakeClass.Value {
		v4.AccelerationDuringBrakeClass = append(v4.AccelerationDuringBrakeClass, *convertFromToClass(&v2FromToClass))
	}
	v4.AccelerationPedalPositionClass = make([]rfmsv4.FromToClass, 0, len(v2.AccelerationPedalPositionClass.Value))
	for _, v2FromToClass := range v2.AccelerationPedalPositionClass.Value {
		v4.AccelerationPedalPositionClass = append(v4.AccelerationPedalPositionClass, *convertFromToClass(&v2FromToClass))
	}
	v4.BrakePedalCounterSpeedOverZero = v2.BrakePedalCounterSpeedOverZero
	v4.ChairliftCounter = v2.ChairliftCounter
	v4.CurrentGearClass = make([]rfmsv4.Label, 0, len(v2.CurrentGearClass.Value))
	for _, v2LabelClass := range v2.CurrentGearClass.Value {
		v4.CurrentGearClass = append(v4.CurrentGearClass, *convertLabelClass(&v2LabelClass))
	}
	v4.DistanceBrakePedalActiveSpeedOverZero = v2.DistanceBrakePedalActiveSpeedOverZero
	v4.DistanceCruiseControlActive = v2.DistanceCruiseControlActive
	v4.DrivingWithoutTorqueClass = make([]rfmsv4.Label, 0, len(v2.DrivingWithoutTorqueClass.Value))
	v4.DurationCruiseControlActive = v2.DurationCruiseControlActive
	v4.DurationWheelbasedSpeedOverZero = v2.DurationWheelbaseSpeedOverZero
	v4.DurationWheelbasedSpeedZero = v2.DurationWheelbaseSpeedZero
	v4.EngineSpeedClass = make([]rfmsv4.FromToClass, 0, len(v2.EngineSpeedClass.Value))
	for _, v2FromToClass := range v2.EngineSpeedClass.Value {
		v4.EngineSpeedClass = append(v4.EngineSpeedClass, *convertFromToClass(&v2FromToClass))
	}
	v4.EngineTorqueAtCurrentSpeedClass = make([]rfmsv4.FromToClassCombustion, 0, len(v2.EngineTorqueAtCurrentSpeedClass.Value))
	for _, v2FromToClass := range v2.EngineTorqueAtCurrentSpeedClass.Value {
		v4.EngineTorqueAtCurrentSpeedClass = append(v4.EngineTorqueAtCurrentSpeedClass, *convertFromToClassCombustion(&v2FromToClass))
	}
	v4.EngineTorqueClass = make([]rfmsv4.FromToClassCombustion, 0, len(v2.EngineTorqueClass.Value))
	for _, v2FromToClass := range v2.EngineTorqueClass.Value {
		v4.EngineTorqueClass = append(v4.EngineTorqueClass, *convertFromToClassCombustion(&v2FromToClass))
	}
	v4.FuelConsumptionDuringCruiseActive = v2.FuelConsumptionCruiseControlActive
	v4.FuelWheelbasedSpeedOverZero = v2.FuelWheelbaseSpeedOverZero
	v4.FuelWheelbasedSpeedZero = v2.FuelWheelbaseSpeedZero
	v4.HighAccelerationClass = make([]rfmsv4.FromToClass, 0, len(v2.HighAccelerationClass.Value))
	v4.KneelingCounter = v2.KneelingCounter
	v4.PramRequestCounter = v2.PramRequestCounter
	v4.PtoActiveClass = make([]rfmsv4.Label, 0, len(v2.PtoActiveClass.Value))
	for _, v2LabelClass := range v2.PtoActiveClass.Value {
		v4.PtoActiveClass = append(v4.PtoActiveClass, *convertLabelClass(&v2LabelClass))
	}
	v4.RetarderTorqueClass = make([]rfmsv4.FromToClass, 0, len(v2.RetarderTorqueClass.Value))
	for _, v2FromToClass := range v2.RetarderTorqueClass.Value {
		v4.RetarderTorqueClass = append(v4.RetarderTorqueClass, *convertFromToClass(&v2FromToClass))
	}
	v4.SelectedGearClass = make([]rfmsv4.Label, 0, len(v2.SelectedGearClass.Value))
	for _, v2LabelClass := range v2.SelectedGearClass.Value {
		v4.SelectedGearClass = append(v4.SelectedGearClass, *convertLabelClass(&v2LabelClass))
	}
	v4.StopRequestCounter = v2.StopRequestCounter
	v4.VehicleSpeedClass = make([]rfmsv4.FromToClass, 0, len(v2.VehicleSpeedClass.Value))
	for _, v2FromToClass := range v2.VehicleSpeedClass.Value {
		v4.VehicleSpeedClass = append(v4.VehicleSpeedClass, *convertFromToClass(&v2FromToClass))
	}
	return &v4
}

func convertFromToClass(v2 *rfmsv2.FromToClass) *rfmsv4.FromToClass {
	if v2 == nil {
		return nil
	}
	var v4 rfmsv4.FromToClass
	v4.From = v2.From
	v4.To = v2.To
	v4.Meters = v2.Meters
	v4.MilliLitres = v2.MilliLitres
	v4.Seconds = v2.Seconds
	return &v4
}

func convertFromToClassCombustion(v2 *rfmsv2.FromToClass) *rfmsv4.FromToClassCombustion {
	if v2 == nil {
		return nil
	}
	var v4 rfmsv4.FromToClassCombustion
	v4.From = v2.From
	v4.To = v2.To
	v4.Meters = v2.Meters
	v4.MilliLitres = v2.MilliLitres
	v4.Seconds = v2.Seconds
	return &v4
}

func convertLabelClass(v2 *rfmsv2.LabelClass) *rfmsv4.Label {
	if v2 == nil {
		return nil
	}
	var v4 rfmsv4.Label
	v4.Label = v2.Label
	v4.Meters = v2.Meters
	v4.MilliLitres = v2.MilliLitres
	v4.Seconds = v2.Seconds
	return &v4
}

func convertDate(date *rfmsv2.Date) *rfmsv4.Date {
	if date == nil {
		return nil
	}
	var result rfmsv4.Date
	result.Day = date.Day
	result.Month = date.Month
	result.Year = date.Year
	return &result
}

func convertGNSSPosition(v2 *rfmsv2.GNSSPosition) *rfmsv4.GNSSPosition {
	if v2 == nil {
		return nil
	}
	var v4 rfmsv4.GNSSPosition
	v4.Altitude = v2.Altitude
	if v2.Heading != nil {
		// Account for Scania providing heading as string.
		v4.Heading = ptr(rfmsv4.Int(*v2.Heading))
	}
	v4.Latitude = v2.Latitude
	v4.Longitude = v2.Longitude
	v4.PositionDateTime = rfmsv4.Time(v2.PositionDateTime)
	v4.Speed = v2.Speed
	return &v4
}

func convertTrigger(v2 *rfmsv2.Trigger) *rfmsv4.Trigger {
	if v2 == nil {
		return nil
	}
	var v4 rfmsv4.Trigger
	v4.Context = v2.Context
	v4.DriverID = convertDriverID(v2.DriverID)
	v4.PtoID = v2.PtoID
	if len(v2.TellTaleInfo) > 0 {
		// Trigger only has one TellTaleInfo in v4.
		v4.TellTaleInfo = convertTellTaleInfo(&v2.TellTaleInfo[0])
	}
	v4.TriggerType = rfmsv4.TriggerType(v2.TriggerType)
	v4.TriggerInfo = v2.TriggerInfo
	return &v4
}

func ptr[T any](v T) *T {
	return &v
}
