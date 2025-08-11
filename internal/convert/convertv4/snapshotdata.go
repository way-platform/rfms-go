package convertv4

import (
	"github.com/way-platform/rfms-go/internal/convert"
	rfmsv4oapi "github.com/way-platform/rfms-go/internal/openapi/rfmsv4oapi"
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/rfms/v5"
)

// snapshotData converts an rFMS v4 snapshot data to proto.
func snapshotData(input *rfmsv4oapi.SnapshotDataObject) *rfmsv5.SnapshotData {
	var output rfmsv5.SnapshotData
	if input.GNSSPosition != nil {
		output.SetGnssPosition(gnssPosition(input.GNSSPosition))
	}
	if input.AmbientAirTemperature != nil {
		output.SetAmbientAirTemperatureC(*input.AmbientAirTemperature)
	}
	if input.EngineSpeed != nil {
		output.SetEngineSpeedRpm(*input.EngineSpeed)
	}
	if input.FuelLevel1 != nil {
		output.SetFuelLevel1Percent(*input.FuelLevel1)
	}
	if input.TachographSpeed != nil {
		output.SetTachographSpeedKmh(*input.TachographSpeed)
	}
	if input.WheelBasedSpeed != nil {
		output.SetWheelBasedSpeedKmh(*input.WheelBasedSpeed)
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
	return &output
}
