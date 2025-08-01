package convertv2

import (
	"github.com/way-platform/rfms-go/internal/convert"
	"github.com/way-platform/rfms-go/internal/openapi/rfmsv2oapi"
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/rfms/v5"
)

// VehicleStatus converts an rFMS v2 vehicle status to proto.
func VehicleStatus(input *rfmsv2oapi.VehicleStatusType) *rfmsv5.VehicleStatus {
	var output rfmsv5.VehicleStatus
	if input.VIN != nil {
		output.SetVin(*input.VIN)
	}
	if input.TriggerType != nil {
		output.SetTrigger(trigger(input.TriggerType))
	}
	if input.CreatedDateTime != nil {
		output.SetCreateTime(input.CreatedDateTime.UnixMicro())
	}
	if input.ReceivedDateTime != nil {
		output.SetReceiveTime(input.ReceivedDateTime.UnixMicro())
	}
	if input.HRTotalVehicleDistance != nil {
		output.SetTotalVehicleDistanceM(*input.HRTotalVehicleDistance)
	}
	if input.TotalEngineHours != nil {
		output.SetTotalEngineHours(*input.TotalEngineHours)
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
	if input.Status2OfDoors != nil {
		output.SetStateOfDoors(convert.StateOfDoors(string(*input.Status2OfDoors)))
		if output.GetStateOfDoors() == rfmsv5.VehicleStatus_STATE_OF_DOORS_UNKNOWN {
			output.SetUnknownStateOfDoors(string(*input.Status2OfDoors))
		}
	}
	if len(input.DoorStatus) > 0 {
		for _, doorStatus := range input.DoorStatus {
			output.SetDoors(append(output.GetDoors(), door(&doorStatus)))
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
