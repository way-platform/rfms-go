package convertv4

import (
	"time"

	"github.com/way-platform/rfms-go/internal/convert"
	"github.com/way-platform/rfms-go/internal/openapi/rfmsv4oapi"
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/rfms/v5"
)

// VehicleStatus converts an rFMS v4 vehicle status to proto.
func VehicleStatus(input *rfmsv4oapi.VehicleStatusObject) *rfmsv5.VehicleStatus {
	var output rfmsv5.VehicleStatus
	if input.VIN != nil {
		output.SetVin(string(*input.VIN))
	}
	if input.TriggerType != nil {
		output.SetTrigger(trigger(input.TriggerType))
	}
	if input.CreatedDateTime != nil {
		output.SetCreateTime(time.Time(*input.CreatedDateTime).UnixMicro())
	}
	if input.ReceivedDateTime != nil {
		output.SetReceiveTime(time.Time(*input.ReceivedDateTime).UnixMicro())
	}
	if input.HrTotalVehicleDistance != nil {
		output.SetTotalVehicleDistanceM(*input.HrTotalVehicleDistance)
	}
	if input.TotalEngineHours != nil {
		output.SetTotalEngineHours(*input.TotalEngineHours)
	}
	if input.TotalElectricMotorHours != nil {
		output.SetTotalElectricMotorHours(*input.TotalElectricMotorHours)
	}
	if input.Driver1ID != nil {
		output.SetDriver1(driverIdentification(input.Driver1ID))
	}
	if input.GrossCombinationVehicleWeight != nil {
		output.SetGrossCombinationVehicleWeightKg(float64(*input.GrossCombinationVehicleWeight))
	}
	if input.EngineTotalFuelUsed != nil {
		output.SetEngineTotalFuelUsedMl(float64(*input.EngineTotalFuelUsed))
	}
	if input.TotalFuelUsedGaseous != nil {
		output.SetTotalFuelUsedGaseousKg(float64(*input.TotalFuelUsedGaseous))
	}
	if input.TotalElectricEnergyUsed != nil {
		output.SetTotalElectricEnergyUsedWh(float64(*input.TotalElectricEnergyUsed))
	}
	if input.Status2OfDoors != nil {
		output.SetStateOfDoors(convert.StateOfDoors(string(*input.Status2OfDoors)))
		if output.GetStateOfDoors() == rfmsv5.VehicleStatus_STATE_OF_DOORS_UNKNOWN {
			output.SetUnknownStateOfDoors(string(*input.Status2OfDoors))
		}
	}
	if len(input.DoorStatus) > 0 {
		for _, doorStatus := range input.DoorStatus {
			var door rfmsv5.VehicleStatus_Door
			if doorStatus.DoorNumber != nil {
				door.SetNumber(*doorStatus.DoorNumber)
			}
			if doorStatus.DoorEnabledStatus != nil {
				door.SetEnabledState(convert.DoorEnabledState(string(*doorStatus.DoorEnabledStatus)))
				if door.GetEnabledState() == rfmsv5.VehicleStatus_Door_ENABLED_STATE_UNKNOWN {
					door.SetUnknownEnabledState(string(*doorStatus.DoorEnabledStatus))
				}
			}
			if doorStatus.DoorOpenStatus != nil {
				door.SetOpenState(convert.DoorOpenState(string(*doorStatus.DoorOpenStatus)))
				if door.GetOpenState() == rfmsv5.VehicleStatus_Door_OPEN_STATE_UNKNOWN {
					door.SetUnknownOpenState(string(*doorStatus.DoorOpenStatus))
				}
			}
			if doorStatus.DoorLockStatus != nil {
				door.SetLockState(convert.DoorLockState(string(*doorStatus.DoorLockStatus)))
				if door.GetLockState() == rfmsv5.VehicleStatus_Door_LOCK_STATE_UNKNOWN {
					door.SetUnknownLockState(string(*doorStatus.DoorLockStatus))
				}
			}
			output.SetDoors(append(output.GetDoors(), &door))
		}
	}
	if input.AccumulatedData != nil {
		output.SetAccumulatedData(accumulatedData(input.AccumulatedData))
	}
	if input.SnapshotData != nil {
		output.SetSnapshotData(snapshotData(input.SnapshotData))
	}
	if input.UptimeData != nil {
		output.SetUptimeData(uptimeData(input.UptimeData))
	}
	return &output
}
