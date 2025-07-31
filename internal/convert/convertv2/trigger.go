package convertv2

import (
	"github.com/way-platform/rfms-go/internal/convert"
	"github.com/way-platform/rfms-go/internal/openapi/rfmsv2oapi"
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/rfms/v5"
)

func trigger(input *rfmsv2oapi.TriggerType) *rfmsv5.Trigger {
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
		output.SetAdditionalInfo(input.TriggerInfo)
	}
	if input.PtoID != nil {
		output.SetPtoId(*input.PtoID)
	}
	if input.DriverID != nil {
		output.SetDriverInfo(driverIdentification(input.DriverID))
	}
	if len(input.TellTaleInfo) > 0 {
		output.SetTellTaleInfo(tellTale(&input.TellTaleInfo[0]))
	}
	return &output
}
