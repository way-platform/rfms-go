package convertv4

import (
	rfmsv4oapi "github.com/way-platform/rfms-go/internal/openapi/rfmsv4oapi"
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/rfms/v5"
)

// accumulatedData converts an rFMS v4 accumulated data to proto.
func accumulatedData(input *rfmsv4oapi.AccumulatedDataObject) *rfmsv5.AccumulatedData {
	var output rfmsv5.AccumulatedData
	if input.DurationWheelbasedSpeedOverZero != nil {
		output.SetWheelBasedSpeedOverZeroDurationS(float64(*input.DurationWheelbasedSpeedOverZero))
	}
	if input.DistanceCruiseControlActive != nil {
		output.SetCruiseControlActiveDistanceM(float64(*input.DistanceCruiseControlActive))
	}
	if input.DurationCruiseControlActive != nil {
		output.SetCruiseControlActiveDurationS(float64(*input.DurationCruiseControlActive))
	}
	if input.FuelConsumptionDuringCruiseActive != nil {
		output.SetCruiseControlActiveFuelConsumptionMl(float64(*input.FuelConsumptionDuringCruiseActive))
	}
	if input.FuelConsumptionDuringCruiseActiveGaseous != nil {
		output.SetCruiseControlActiveFuelConsumptionGaseousKg(float64(*input.FuelConsumptionDuringCruiseActiveGaseous))
	}
	if input.ElectricEnergyConsumptionDuringCruiseActive != nil {
		output.SetCruiseControlActiveElectricEnergyConsumptionWh(float64(*input.ElectricEnergyConsumptionDuringCruiseActive))
	}
	if input.DurationWheelbasedSpeedZero != nil {
		output.SetWheelBasedSpeedZeroDurationS(float64(*input.DurationWheelbasedSpeedZero))
	}
	if input.FuelWheelbasedSpeedZero != nil {
		output.SetWheelBasedSpeedZeroFuelMl(float64(*input.FuelWheelbasedSpeedZero))
	}
	if input.FuelWheelbasedSpeedZeroGaseous != nil {
		output.SetWheelBasedSpeedZeroFuelGaseousKg(float64(*input.FuelWheelbasedSpeedZeroGaseous))
	}
	if input.ElectricEnergyWheelbasedSpeedZero != nil {
		output.SetWheelBasedSpeedZeroElectricEnergyConsumptionWh(float64(*input.ElectricEnergyWheelbasedSpeedZero))
	}
	if input.FuelWheelbasedSpeedOverZero != nil {
		output.SetWheelBasedSpeedOverZeroFuelMl(float64(*input.FuelWheelbasedSpeedOverZero))
	}
	if input.FuelWheelbasedSpeedOverZeroGaseous != nil {
		output.SetWheelBasedSpeedOverZeroFuelGaseousKg(float64(*input.FuelWheelbasedSpeedOverZeroGaseous))
	}
	if input.ElectricEnergyWheelbasedSpeedOverZero != nil {
		output.SetWheelBasedSpeedOverZeroElectricEnergyConsumptionWh(float64(*input.ElectricEnergyWheelbasedSpeedOverZero))
	}
	if input.ElectricEnergyAux != nil {
		output.SetAuxElectricEnergyConsumptionWh(float64(*input.ElectricEnergyAux))
	}
	if input.PtoActiveClass != nil {
		for _, labelPtoActiveClass := range input.PtoActiveClass {
			output.SetPtoActiveClass(append(output.GetPtoActiveClass(), labelClass(&labelPtoActiveClass)))
		}
	}
	if input.BrakePedalCounterSpeedOverZero != nil {
		output.SetBrakePedalSpeedOverZeroCount(int32(*input.BrakePedalCounterSpeedOverZero))
	}
	if input.DistanceBrakePedalActiveSpeedOverZero != nil {
		output.SetBrakePedalSpeedOverZeroDistanceM(float64(*input.DistanceBrakePedalActiveSpeedOverZero))
	}
	if input.AccelerationPedalPositionClass != nil {
		for _, fromToAccelerationPedalPositionClass := range input.AccelerationPedalPositionClass {
			output.SetAccelerationPedalPositionClassPercent(append(output.GetAccelerationPedalPositionClassPercent(), fromToClass(&fromToAccelerationPedalPositionClass)))
		}
	}
	if input.BrakePedalPositionClass != nil {
		for _, fromToBrakePedalPositionClass := range input.BrakePedalPositionClass {
			output.SetBrakePedalPositionClassPercent(append(output.GetBrakePedalPositionClassPercent(), fromToClass(&fromToBrakePedalPositionClass)))
		}
	}
	if input.AccelerationClass != nil {
		for _, fromToAccelerationClass := range input.AccelerationClass {
			output.SetAccelerationClassMps2(append(output.GetAccelerationClassMps2(), fromToClass(&fromToAccelerationClass)))
		}
	}
	if input.HighAccelerationClass != nil {
		for _, fromToHighAccelerationClass := range input.HighAccelerationClass {
			output.SetHighAccelerationClassMps2(append(output.GetHighAccelerationClassMps2(), fromToClass(&fromToHighAccelerationClass)))
		}
	}
	if input.RetarderTorqueClass != nil {
		for _, fromToRetarderTorqueClass := range input.RetarderTorqueClass {
			output.SetRetarderTorqueClassPercent(append(output.GetRetarderTorqueClassPercent(), fromToClass(&fromToRetarderTorqueClass)))
		}
	}
	if input.DrivingWithoutTorqueClass != nil {
		for _, labelDrivingWithoutTorqueClass := range input.DrivingWithoutTorqueClass {
			output.SetDrivingWithoutTorqueClass(append(output.GetDrivingWithoutTorqueClass(), labelClass(&labelDrivingWithoutTorqueClass)))
		}
	}
	if input.EngineTorqueClass != nil {
		for _, fromToEngineTorqueClass := range input.EngineTorqueClass {
			output.SetEngineTorqueClassPercent(append(output.GetEngineTorqueClassPercent(), fromToClassCombustion(&fromToEngineTorqueClass)))
		}
	}
	if input.ElectricMotorTorqueClass != nil {
		for _, fromToElectricMotorTorqueClass := range input.ElectricMotorTorqueClass {
			output.SetElectricMotorTorqueClassPercent(append(output.GetElectricMotorTorqueClassPercent(), fromToClassElectrical(&fromToElectricMotorTorqueClass)))
		}
	}
	if input.EngineTorqueAtCurrentSpeedClass != nil {
		for _, fromToEngineTorqueAtCurrentSpeedClass := range input.EngineTorqueAtCurrentSpeedClass {
			output.SetEngineTorqueAtCurrentSpeedClassPercent(append(output.GetEngineTorqueAtCurrentSpeedClassPercent(), fromToClassCombustion(&fromToEngineTorqueAtCurrentSpeedClass)))
		}
	}
	if input.ElectricMotorTorqueAtCurrentSpeedClass != nil {
		for _, fromToElectricMotorTorqueAtCurrentSpeedClass := range input.ElectricMotorTorqueAtCurrentSpeedClass {
			output.SetElectricMotorTorqueAtCurrentSpeedClassPercent(append(output.GetElectricMotorTorqueAtCurrentSpeedClassPercent(), fromToClassElectrical(&fromToElectricMotorTorqueAtCurrentSpeedClass)))
		}
	}
	if input.VehicleSpeedClass != nil {
		for _, fromToVehicleSpeedClass := range input.VehicleSpeedClass {
			output.SetVehicleSpeedClassKmh(append(output.GetVehicleSpeedClassKmh(), fromToClass(&fromToVehicleSpeedClass)))
		}
	}
	if input.EngineSpeedClass != nil {
		for _, fromToEngineSpeedClass := range input.EngineSpeedClass {
			output.SetEngineSpeedClassRpm(append(output.GetEngineSpeedClassRpm(), fromToClass(&fromToEngineSpeedClass)))
		}
	}
	if input.AccelerationDuringBrakeClass != nil {
		for _, fromToAccelerationDuringBrakeClass := range input.AccelerationDuringBrakeClass {
			output.SetAccelerationDuringBrakeClassMps2(append(output.GetAccelerationDuringBrakeClassMps2(), fromToClass(&fromToAccelerationDuringBrakeClass)))
		}
	}
	if input.SelectedGearClass != nil {
		for _, labelSelectedGearClass := range input.SelectedGearClass {
			output.SetSelectedGearClass(append(output.GetSelectedGearClass(), labelClass(&labelSelectedGearClass)))
		}
	}
	if input.CurrentGearClass != nil {
		for _, labelCurrentGearClass := range input.CurrentGearClass {
			output.SetCurrentGearClass(append(output.GetCurrentGearClass(), labelClass(&labelCurrentGearClass)))
		}
	}
	if input.ChairliftCounter != nil {
		output.SetChairliftCount(int32(*input.ChairliftCounter))
	}
	if input.StopRequestCounter != nil {
		output.SetStopRequestCount(int32(*input.StopRequestCounter))
	}
	if input.KneelingCounter != nil {
		output.SetKneelingCount(int32(*input.KneelingCounter))
	}
	if input.PramRequestCounter != nil {
		output.SetPramRequestCount(int32(*input.PramRequestCounter))
	}
	if input.ElectricPowerRecuperationClass != nil {
		for _, fromToElectricPowerRecuperationClass := range input.ElectricPowerRecuperationClass {
			output.SetElectricPowerRecuperationClassKw(append(output.GetElectricPowerRecuperationClassKw(), fromToClassElectrical(&fromToElectricPowerRecuperationClass)))
		}
	}
	return &output
}

func labelClass(input *rfmsv4oapi.LabelObject) *rfmsv5.AccumulatedData_LabelClass {
	var output rfmsv5.AccumulatedData_LabelClass
	if input.Label != nil {
		output.SetLabel(*input.Label)
	}
	if input.Seconds != nil {
		output.SetDurationS(float64(*input.Seconds))
	}
	if input.Meters != nil {
		output.SetDistanceM(float64(*input.Meters))
	}
	if input.MilliLitres != nil {
		output.SetFuelConsumptionMl(float64(*input.MilliLitres))
	}
	if input.Kilograms != nil {
		output.SetFuelConsumptionGaseousKg(float64(*input.Kilograms))
	}
	if input.Watthours != nil {
		output.SetElectricEnergyConsumptionWh(float64(*input.Watthours))
	}
	return &output
}

func fromToClass(input *rfmsv4oapi.FromToClassObject) *rfmsv5.AccumulatedData_FromToClass {
	var output rfmsv5.AccumulatedData_FromToClass
	if input.From != nil {
		output.SetFrom(*input.From)
	}
	if input.To != nil {
		output.SetTo(*input.To)
	}
	if input.Seconds != nil {
		output.SetDurationS(float64(*input.Seconds))
	}
	if input.Meters != nil {
		output.SetDistanceM(float64(*input.Meters))
	}
	if input.MilliLitres != nil {
		output.SetFuelConsumptionMl(float64(*input.MilliLitres))
	}
	if input.Kilograms != nil {
		output.SetFuelConsumptionGaseousKg(float64(*input.Kilograms))
	}
	if input.Watthours != nil {
		output.SetElectricEnergyConsumptionWh(float64(*input.Watthours))
	}
	return &output
}

func fromToClassCombustion(input *rfmsv4oapi.FromToClassObjectCombustion) *rfmsv5.AccumulatedData_FromToClassCombustion {
	var output rfmsv5.AccumulatedData_FromToClassCombustion
	if input.From != nil {
		output.SetFrom(*input.From)
	}
	if input.To != nil {
		output.SetTo(*input.To)
	}
	if input.Seconds != nil {
		output.SetDurationS(float64(*input.Seconds))
	}
	if input.Meters != nil {
		output.SetDistanceM(float64(*input.Meters))
	}
	if input.MilliLitres != nil {
		output.SetFuelConsumptionMl(float64(*input.MilliLitres))
	}
	if input.Kilograms != nil {
		output.SetFuelConsumptionGaseousKg(float64(*input.Kilograms))
	}
	return &output
}

func fromToClassElectrical(input *rfmsv4oapi.FromToClassObjectElectrical) *rfmsv5.AccumulatedData_FromToClassElectrical {
	var output rfmsv5.AccumulatedData_FromToClassElectrical
	if input.From != nil {
		output.SetFrom(*input.From)
	}
	if input.To != nil {
		output.SetTo(*input.To)
	}
	if input.Seconds != nil {
		output.SetDurationS(float64(*input.Seconds))
	}
	if input.Meters != nil {
		output.SetDistanceM(float64(*input.Meters))
	}
	if input.Watthours != nil {
		output.SetElectricEnergyConsumptionWh(float64(*input.Watthours))
	}
	return &output
}
