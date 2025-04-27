package rfms

import (
	"strings"
	"time"
)

// Time is a time.Time that can be unmarshalled from an ISO 8601 string.
type Time time.Time

// UnmarshalJSON unmarshals the time from an ISO 8601 string.
func (t *Time) UnmarshalJSON(data []byte) error {
	parsed, err := time.Parse(time.RFC3339, string(strings.Trim(string(data), `Z"`)+"Z"))
	if err != nil {
		return err
	}
	*t = Time(parsed)
	return nil
}
