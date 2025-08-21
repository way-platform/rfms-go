package convertv4

import (
	"time"

	"github.com/way-platform/rfms-go/internal/convert"
	rfmsv4oapi "github.com/way-platform/rfms-go/internal/openapi/rfmsv4oapi"
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/connect/rfms/v5"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// snapshotData converts an rFMS v4 snapshot data to proto.
func snapshotData(input *rfmsv4oapi.SnapshotDataObject) *rfmsv5.SnapshotData {
	var output rfmsv5.SnapshotData
	if input.GNSSPosition != nil {
		output.SetGnssPosition(gnssPosition(input.GNSSPosition))
	}
	if input.TachographSpeed != nil {
		output.SetTachographSpeedKmh(*input.TachographSpeed)
	}
	if input.WheelBasedSpeed != nil {
		output.SetWheelBasedSpeedKmh(*input.WheelBasedSpeed)
	}
	if input.AmbientAirTemperature != nil {
		output.SetAmbientAirTemperatureC(*input.AmbientAirTemperature)
	}
	if input.EngineSpeed != nil {
		output.SetEngineSpeedRpm(*input.EngineSpeed)
	}
	if input.ElectricMotorSpeed != nil {
		output.SetElectricMotorSpeedRpm(*input.ElectricMotorSpeed)
	}
	if input.FuelType != nil {
		output.SetFuelType(convert.FuelType(string(*input.FuelType)))
		if output.GetFuelType() == rfmsv5.FuelType_FUEL_TYPE_UNKNOWN {
			output.SetUnknownFuelType(string(*input.FuelType))
		}
	}
	if input.FuelLevel1 != nil {
		output.SetFuelLevel1Percent(*input.FuelLevel1)
	}
	if input.FuelLevel2 != nil {
		output.SetFuelLevel2Percent(*input.FuelLevel2)
	}
	if input.CatalystFuelLevel != nil {
		output.SetCatalystFuelLevelPercent(*input.CatalystFuelLevel)
	}
	if input.Driver1WorkingState != nil {
		output.SetDriver1WorkingState(convert.DriverWorkingState(string(*input.Driver1WorkingState)))
		if output.GetDriver1WorkingState() == rfmsv5.DriverWorkingState_DRIVER_WORKING_STATE_UNKNOWN {
			output.SetUnknownDriver1WorkingState(string(*input.Driver1WorkingState))
		}
	}
	if input.Driver2WorkingState != nil {
		output.SetDriver2WorkingState(convert.DriverWorkingState(string(*input.Driver2WorkingState)))
		if output.GetDriver2WorkingState() == rfmsv5.DriverWorkingState_DRIVER_WORKING_STATE_UNKNOWN {
			output.SetUnknownDriver2WorkingState(string(*input.Driver2WorkingState))
		}
	}
	if input.Driver2ID != nil {
		output.SetDriver2Id(driverIdentification(input.Driver2ID))
	}
	if input.CatalystFuelLevel != nil {
		output.SetCatalystFuelLevelPercent(*input.CatalystFuelLevel)
	}
	if input.ParkingBrakeSwitch != nil {
		output.SetParkingBrakeSwitch(*input.ParkingBrakeSwitch)
	}
	if input.HybridBatteryPackRemainingCharge != nil {
		output.SetBatteryPackRemainingChargePercent(*input.HybridBatteryPackRemainingCharge)
	}
	if input.BatteryPackChargingStatus != nil {
		output.SetBatteryPackChargingState(convert.ChargingState(string(*input.BatteryPackChargingStatus)))
		if output.GetBatteryPackChargingState() == rfmsv5.ChargingState_CHARGING_STATE_UNKNOWN {
			output.SetUnknownBatteryPackChargingState(string(*input.BatteryPackChargingStatus))
		}
	}
	if input.BatteryPackChargingConnectionStatus != nil {
		output.SetBatteryPackChargingConnectionState(convert.ChargingConnectionState(string(*input.BatteryPackChargingConnectionStatus)))
		if output.GetBatteryPackChargingConnectionState() == rfmsv5.ChargingConnectionState_CHARGING_CONNECTION_STATE_UNKNOWN {
			output.SetUnknownBatteryPackChargingConnectionState(string(*input.BatteryPackChargingConnectionStatus))
		}
	}
	if input.BatteryPackChargingDevice != nil {
		output.SetBatteryPackChargingDevice(convert.ChargingDevice(string(*input.BatteryPackChargingDevice)))
		if output.GetBatteryPackChargingDevice() == rfmsv5.ChargingDevice_CHARGING_DEVICE_UNKNOWN {
			output.SetUnknownBatteryPackChargingDevice(string(*input.BatteryPackChargingDevice))
		}
	}
	if input.BatteryPackChargingPower != nil {
		output.SetBatteryPackChargingPowerW(*input.BatteryPackChargingPower)
	}
	if input.EstimatedTimeBatteryPackChargingCompleted != nil {
		output.SetBatteryPackEstimatedChargingCompletedTime(
			timestamppb.New(time.Time(*input.EstimatedTimeBatteryPackChargingCompleted).UTC()),
		)
	}
	if input.EstimatedDistanceToEmpty != nil {
		output.SetEstimatedDistanceToEmpty(&rfmsv5.SnapshotData_EstimatedDistanceToEmpty{})
		if input.EstimatedDistanceToEmpty.Total != nil {
			output.GetEstimatedDistanceToEmpty().SetTotalM(float64(*input.EstimatedDistanceToEmpty.Total))
		}
		if input.EstimatedDistanceToEmpty.Fuel != nil {
			output.GetEstimatedDistanceToEmpty().SetFuelM(float64(*input.EstimatedDistanceToEmpty.Fuel))
		}
		if input.EstimatedDistanceToEmpty.Gas != nil {
			output.GetEstimatedDistanceToEmpty().SetGasM(float64(*input.EstimatedDistanceToEmpty.Gas))
		}
		if input.EstimatedDistanceToEmpty.BatteryPack != nil {
			output.GetEstimatedDistanceToEmpty().SetBatteryPackM(float64(*input.EstimatedDistanceToEmpty.BatteryPack))
		}
	}
	for _, inputVehicleAxle := range input.VehicleAxles {
		var outputVehicleAxle rfmsv5.VehicleAxle
		if inputVehicleAxle.VehicleAxlePosition != nil {
			outputVehicleAxle.SetPosition(*inputVehicleAxle.VehicleAxlePosition)
		}
		if inputVehicleAxle.VehicleAxleLoad != nil {
			outputVehicleAxle.SetLoadKg(float64(*inputVehicleAxle.VehicleAxleLoad))
		}
		output.SetVehicleAxles(append(output.GetVehicleAxles(), &outputVehicleAxle))
	}
	for _, inputTrailer := range input.Trailers {
		var outputTrailer rfmsv5.Trailer
		if inputTrailer.TrailerNo != nil {
			outputTrailer.SetNumber(int32(*inputTrailer.TrailerNo))
		}
		if inputTrailer.TrailerIdentificationData != nil {
			outputTrailer.SetIdentificationData(string(*inputTrailer.TrailerIdentificationData))
		}
		if inputTrailer.TrailerVIN != nil {
			outputTrailer.SetVin(string(*inputTrailer.TrailerVIN))
		}
		if inputTrailer.CustomerTrailerName != nil {
			outputTrailer.SetCustomerName(string(*inputTrailer.CustomerTrailerName))
		}
		if inputTrailer.TrailerType != nil {
			outputTrailer.SetType(convert.TrailerType(string(*inputTrailer.TrailerType)))
		}
		if inputTrailer.TrailerAxleLoadSum != nil {
			outputTrailer.SetAxleLoadSumKg(float64(*inputTrailer.TrailerAxleLoadSum))
		}
		for _, inputAxle := range inputTrailer.TrailerAxles {
			var trailerAxle rfmsv5.Trailer_Axle
			if inputAxle.TrailerAxlePosition != nil {
				trailerAxle.SetPosition(int32(*inputAxle.TrailerAxlePosition))
			}
			if inputAxle.TrailerAxleLoad != nil {
				trailerAxle.SetLoadKg(float64(*inputAxle.TrailerAxleLoad))
			}
			outputTrailer.SetAxles(append(outputTrailer.GetAxles(), &trailerAxle))
		}
		output.SetTrailers(append(output.GetTrailers(), &outputTrailer))
	}
	return &output
}
