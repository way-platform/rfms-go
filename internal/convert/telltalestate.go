package convert

import (
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/rfms/v5"
)

// TellTaleState converts an rFMS tell tale state to proto.
func TellTaleState(input string) rfmsv5.TellTale_State {
	switch input {
	case "RED":
		return rfmsv5.TellTale_RED
	case "YELLOW":
		return rfmsv5.TellTale_YELLOW
	case "INFO":
		return rfmsv5.TellTale_INFO
	case "OFF":
		return rfmsv5.TellTale_OFF
	case "NOT_AVAILABLE":
		return rfmsv5.TellTale_STATE_NOT_AVAILABLE
	default:
		return rfmsv5.TellTale_STATE_UNKNOWN
	}
}
