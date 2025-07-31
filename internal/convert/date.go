package convert

import (
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/rfms/v5"
)

// Date converts an rFMS date to proto.
func Date(year, month, day *int32) *rfmsv5.Date {
	var result rfmsv5.Date
	if year != nil {
		result.SetYear(*year)
	}
	if month != nil {
		result.SetMonth(*month)
	}
	if day != nil {
		result.SetDay(*day)
	}
	return &result
}
