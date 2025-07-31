package convertv2

import (
	"github.com/way-platform/rfms-go/internal/convert"
	rfmsv2oapi "github.com/way-platform/rfms-go/internal/openapi/rfmsv2oapi"
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/rfms/v5"
)

// tellTale converts an rFMS v2 tell tale info to proto.
func tellTale(input *rfmsv2oapi.TellTaleInfoType) *rfmsv5.TellTale {
	var output rfmsv5.TellTale
	if input.State != nil {
		output.SetState(convert.TellTaleState(string(*input.State)))
		if output.GetState() == rfmsv5.TellTale_STATE_UNKNOWN {
			output.SetUnknownState(string(*input.State))
		}
	}
	if input.TellTale != nil {
		output.SetType(convert.TellTaleType(string(*input.TellTale)))
		if output.GetType() == rfmsv5.TellTale_TYPE_UNKNOWN {
			output.SetUnknownType(string(*input.TellTale))
		}
	}
	return &output
}
