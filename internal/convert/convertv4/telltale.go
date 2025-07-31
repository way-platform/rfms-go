package convertv4

import (
	"github.com/way-platform/rfms-go/internal/convert"
	rfmsv4oapi "github.com/way-platform/rfms-go/internal/openapi/rfmsv4oapi"
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/rfms/v5"
)

// tellTale converts an rFMS v4 tell tale to proto.
func tellTale(input *rfmsv4oapi.TellTaleObject) *rfmsv5.TellTale {
	var output rfmsv5.TellTale
	if input.TellTale != nil {
		output.SetType(convert.TellTaleType(string(*input.TellTale)))
		if output.GetType() == rfmsv5.TellTale_TYPE_UNKNOWN {
			output.SetUnknownType(string(*input.TellTale))
		}
	}
	if input.State != nil {
		output.SetState(convert.TellTaleState(string(*input.State)))
		if output.GetState() == rfmsv5.TellTale_STATE_UNKNOWN {
			output.SetUnknownState(string(*input.State))
		}
	}
	if input.OEMTellTale != nil {
		output.SetOemSpecificType(*input.OEMTellTale)
	}
	return &output
}
