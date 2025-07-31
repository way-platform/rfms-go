package convert

import (
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/rfms/v5"
)

// TriggerContext converts an rFMS trigger context to proto.
func TriggerContext(input string) rfmsv5.Trigger_Context {
	switch input {
	case "RFMS":
		return rfmsv5.Trigger_RFMS
	default:
		return rfmsv5.Trigger_CONTEXT_UNKNOWN
	}
}
