package convertv2

import (
	"github.com/way-platform/rfms-go/internal/convert"
	rfmsv2oapi "github.com/way-platform/rfms-go/internal/openapi/rfmsv2oapi"
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/connect/rfms/v5"
)

// uptimeData converts an rFMS v2 uptime data to proto.
func uptimeData(input *rfmsv2oapi.UptimeType) *rfmsv5.UptimeData {
	var output rfmsv5.UptimeData
	if input.EngineCoolantTemperature != nil {
		output.SetEngineCoolantTemperatureC(*input.EngineCoolantTemperature)
	}
	if input.ServiceDistance != nil {
		output.SetServiceDistanceM(float64(*input.ServiceDistance))
	}
	if input.ServiceBrakeAirPressureCircuit1 != nil {
		output.SetServiceBrakeAirPressureCircuit1Pa(float64(*input.ServiceBrakeAirPressureCircuit1))
	}
	if input.ServiceBrakeAirPressureCircuit2 != nil {
		output.SetServiceBrakeAirPressureCircuit2Pa(float64(*input.ServiceBrakeAirPressureCircuit2))
	}
	if input.DurationAtLeastOneDoorOpen != nil {
		output.SetAtLeastOneDoorOpenDurationS(float64(*input.DurationAtLeastOneDoorOpen))
	}
	if input.BellowPressureFrontAxleLeft != nil {
		output.SetBellowPressureFrontAxleLeftPa(float64(*input.BellowPressureFrontAxleLeft))
	}
	if input.BellowPressureFrontAxleRight != nil {
		output.SetBellowPressureFrontAxleRightPa(float64(*input.BellowPressureFrontAxleRight))
	}
	if input.BellowPressureRearAxleLeft != nil {
		output.SetBellowPressureRearAxleLeftPa(float64(*input.BellowPressureRearAxleLeft))
	}
	if input.BellowPressureRearAxleRight != nil {
		output.SetBellowPressureRearAxleRightPa(float64(*input.BellowPressureRearAxleRight))
	}
	if len(input.AlternatorInfo) > 0 {
		for _, alternatorInfo := range input.AlternatorInfo {
			var alternator rfmsv5.UptimeData_Alternator
			if alternatorInfo.AlternatorNumber != nil {
				alternator.SetNumber(int32(*alternatorInfo.AlternatorNumber))
			}
			if alternatorInfo.AlternatorStatus != nil {
				alternator.SetState(convert.AlternatorState(string(*alternatorInfo.AlternatorStatus)))
				if alternator.GetState() == rfmsv5.UptimeData_Alternator_STATE_UNKNOWN {
					alternator.SetUnknownState(string(*alternatorInfo.AlternatorStatus))
				}
			}
			output.SetAlternators(append(output.GetAlternators(), &alternator))
		}
	}
	if input.DurationAtLeastOneDoorOpen != nil {
		output.SetAtLeastOneDoorOpenDurationS(float64(*input.DurationAtLeastOneDoorOpen))
	}
	if len(input.TellTaleInfo) > 0 {
		for _, tellTaleInfo := range input.TellTaleInfo {
			output.SetTellTales(append(output.GetTellTales(), tellTale(&tellTaleInfo)))
		}
	}
	return &output
}
