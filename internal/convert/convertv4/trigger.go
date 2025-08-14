package convertv4

import (
	"github.com/way-platform/rfms-go/internal/convert"
	rfmsv4oapi "github.com/way-platform/rfms-go/internal/openapi/rfmsv4oapi"
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/connect/rfms/v5"
)

// trigger converts an rFMS v4 trigger to proto.
func trigger(input *rfmsv4oapi.TriggerObject) *rfmsv5.Trigger {
	var output rfmsv5.Trigger
	if input.TriggerType != nil {
		output.SetType(convert.TriggerType(*input.TriggerType))
		if output.GetType() == rfmsv5.Trigger_TYPE_UNKNOWN {
			output.SetUnknownType(*input.TriggerType)
		}
	}
	if input.Context != nil {
		output.SetContext(convert.TriggerContext(*input.Context))
		if output.GetContext() == rfmsv5.Trigger_CONTEXT_UNKNOWN {
			output.SetUnknownContext(*input.Context)
		}
	}
	if len(input.TriggerInfo) > 0 {
		output.SetTriggerInfo(input.TriggerInfo)
	}
	if input.PtoID != nil {
		output.SetPtoId(*input.PtoID)
	}
	if input.DriverID != nil {
		output.SetDriverId(driverIdentification(input.DriverID))
	}
	if input.TellTaleInfo != nil {
		output.SetTellTaleInfo(tellTale(input.TellTaleInfo))
	}
	// TODO: Handle ChargingStatusInfo and ChargingConnectionStatusInfo when needed
	return &output
}
