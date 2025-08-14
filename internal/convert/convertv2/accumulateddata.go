package convertv2

import (
	rfmsv2oapi "github.com/way-platform/rfms-go/internal/openapi/rfmsv2oapi"
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/connect/rfms/v5"
)

// accumulatedData converts an rFMS v2 accumulated data to proto.
func accumulatedData(input *rfmsv2oapi.AccumulatedType) *rfmsv5.AccumulatedData {
	var output rfmsv5.AccumulatedData
	if input.AccelerationClass != nil {
		for _, fromToAccelerationClass := range input.AccelerationClass.Value {
			output.SetAccelerationClassMps2(append(output.GetAccelerationClassMps2(), fromToClass(&fromToAccelerationClass)))
		}
	}
	if input.AccelerationDuringBrakeClass != nil {
		for _, fromToAccelerationDuringBrakeClass := range input.AccelerationDuringBrakeClass.Value {
			output.SetAccelerationDuringBrakeClassMps2(append(output.GetAccelerationDuringBrakeClassMps2(), fromToClass(&fromToAccelerationDuringBrakeClass)))
		}
	}
	if input.AccelerationPedalPositionClass != nil {
		for _, fromToAccelerationPedalPositionClass := range input.AccelerationPedalPositionClass.Value {
			output.SetAccelerationPedalPositionClassPercent(append(output.GetAccelerationPedalPositionClassPercent(), fromToClass(&fromToAccelerationPedalPositionClass)))
		}
	}
	if input.BrakePedalCounterSpeedOverZero != nil {
		output.SetBrakePedalSpeedOverZeroCount(int32(*input.BrakePedalCounterSpeedOverZero))
	}
	if input.ChairliftCounter != nil {
		output.SetChairliftCount(int32(*input.ChairliftCounter))
	}
	if input.CurrentGearClass != nil {
		for _, labelCurrentGearClass := range input.CurrentGearClass.Value {
			output.SetCurrentGearClass(append(output.GetCurrentGearClass(), labelClass(&labelCurrentGearClass)))
		}
	}
	if input.DistanceBrakePedalActiveSpeedOverZero != nil {
		output.SetBrakePedalSpeedOverZeroDistanceM(float64(*input.DistanceBrakePedalActiveSpeedOverZero))
	}
	if input.DistanceCruiseControlActive != nil {
		output.SetCruiseControlActiveDistanceM(float64(*input.DistanceCruiseControlActive))
	}
	if input.DurationCruiseControlActive != nil {
		output.SetCruiseControlActiveDurationS(float64(*input.DurationCruiseControlActive))
	}
	if input.DurationWheelbaseSpeedOverZero != nil {
		output.SetWheelBasedSpeedOverZeroDurationS(float64(*input.DurationWheelbaseSpeedOverZero))
	}
	if input.DurationWheelbaseSpeedZero != nil {
		output.SetWheelBasedSpeedZeroDurationS(float64(*input.DurationWheelbaseSpeedZero))
	}
	if input.DrivingWithoutTorqueClass != nil {
		for _, labelDrivingWithoutTorqueClass := range input.DrivingWithoutTorqueClass.Value {
			output.SetDrivingWithoutTorqueClass(append(output.GetDrivingWithoutTorqueClass(), labelClass(&labelDrivingWithoutTorqueClass)))
		}
	}
	if input.DurationCruiseControlActive != nil {
		output.SetCruiseControlActiveDurationS(float64(*input.DurationCruiseControlActive))
	}
	if input.DurationWheelbaseSpeedOverZero != nil {
		output.SetWheelBasedSpeedOverZeroDurationS(float64(*input.DurationWheelbaseSpeedOverZero))
	}
	if input.DurationWheelbaseSpeedZero != nil {
		output.SetWheelBasedSpeedZeroDurationS(float64(*input.DurationWheelbaseSpeedZero))
	}
	if input.EngineSpeedClass != nil {
		for _, fromToEngineSpeedClass := range input.EngineSpeedClass.Value {
			output.SetEngineSpeedClassRpm(append(output.GetEngineSpeedClassRpm(), fromToClass(&fromToEngineSpeedClass)))
		}
	}
	if input.EngineTorqueClass != nil {
		for _, fromToEngineTorqueClass := range input.EngineTorqueClass.Value {
			output.SetEngineTorqueClassPercent(append(output.GetEngineTorqueClassPercent(), fromToClassCombustion(&fromToEngineTorqueClass)))
		}
	}
	if input.FuelConsumptionCruiseControlActive != nil {
		output.SetCruiseControlActiveFuelConsumptionMl(float64(*input.FuelConsumptionCruiseControlActive))
	}
	if input.FuelWheelbaseSpeedOverZero != nil {
		output.SetWheelBasedSpeedOverZeroFuelMl(float64(*input.FuelWheelbaseSpeedOverZero))
	}
	if input.FuelWheelbaseSpeedOverZero != nil {
		output.SetWheelBasedSpeedOverZeroFuelMl(float64(*input.FuelWheelbaseSpeedOverZero))
	}
	if input.FuelWheelbaseSpeedZero != nil {
		output.SetWheelBasedSpeedZeroFuelMl(float64(*input.FuelWheelbaseSpeedZero))
	}
	if input.HighAccelerationClass != nil {
		for _, fromToHighAccelerationClass := range input.HighAccelerationClass.Value {
			output.SetHighAccelerationClassMps2(append(output.GetHighAccelerationClassMps2(), fromToClass(&fromToHighAccelerationClass)))
		}
	}
	if input.KneelingCounter != nil {
		output.SetKneelingCount(int32(*input.KneelingCounter))
	}
	if input.PramRequestCounter != nil {
		output.SetPramRequestCount(int32(*input.PramRequestCounter))
	}
	if input.PtoActiveClass != nil {
		for _, labelPtoActiveClass := range input.PtoActiveClass.Value {
			output.SetPtoActiveClass(append(output.GetPtoActiveClass(), labelClass(&labelPtoActiveClass)))
		}
	}
	if input.RetarderTorqueClass != nil {
		for _, fromToRetarderTorqueClass := range input.RetarderTorqueClass.Value {
			output.SetRetarderTorqueClassPercent(append(output.GetRetarderTorqueClassPercent(), fromToClass(&fromToRetarderTorqueClass)))
		}
	}
	if input.SelectedGearClass != nil {
		for _, labelSelectedGearClass := range input.SelectedGearClass.Value {
			output.SetSelectedGearClass(append(output.GetSelectedGearClass(), labelClass(&labelSelectedGearClass)))
		}
	}
	if input.StopRequestCounter != nil {
		output.SetStopRequestCount(int32(*input.StopRequestCounter))
	}
	if input.VehicleSpeedClass != nil {
		for _, fromToVehicleSpeedClass := range input.VehicleSpeedClass.Value {
			output.SetVehicleSpeedClassKmh(append(output.GetVehicleSpeedClassKmh(), fromToClass(&fromToVehicleSpeedClass)))
		}
	}
	return &output
}

func fromToClass(input *rfmsv2oapi.FromToClassType) *rfmsv5.AccumulatedData_FromToClass {
	var output rfmsv5.AccumulatedData_FromToClass
	if input.From != nil {
		output.SetFrom(*input.From)
	}
	if input.To != nil {
		output.SetTo(*input.To)
	}
	if input.Meters != nil {
		output.SetDistanceM(float64(*input.Meters))
	}
	if input.MilliLitres != nil {
		output.SetFuelConsumptionMl(float64(*input.MilliLitres))
	}
	if input.Seconds != nil {
		output.SetDurationS(float64(*input.Seconds))
	}
	return &output
}

func fromToClassCombustion(input *rfmsv2oapi.FromToClassType) *rfmsv5.AccumulatedData_FromToClassCombustion {
	var output rfmsv5.AccumulatedData_FromToClassCombustion
	if input.From != nil {
		output.SetFrom(*input.From)
	}
	if input.To != nil {
		output.SetTo(*input.To)
	}
	if input.Meters != nil {
		output.SetDistanceM(float64(*input.Meters))
	}
	if input.MilliLitres != nil {
		output.SetFuelConsumptionMl(float64(*input.MilliLitres))
	}
	if input.Seconds != nil {
		output.SetDurationS(float64(*input.Seconds))
	}
	return &output
}

func labelClass(input *rfmsv2oapi.LabelClassType) *rfmsv5.AccumulatedData_LabelClass {
	var output rfmsv5.AccumulatedData_LabelClass
	if input.Label != nil {
		output.SetLabel(*input.Label)
	}
	if input.Meters != nil {
		output.SetDistanceM(float64(*input.Meters))
	}
	if input.Seconds != nil {
		output.SetDurationS(float64(*input.Seconds))
	}
	if input.MilliLitres != nil {
		output.SetFuelConsumptionMl(float64(*input.MilliLitres))
	}
	return &output
}
