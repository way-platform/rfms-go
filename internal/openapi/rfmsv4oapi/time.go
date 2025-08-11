package rfmsv4oapi

import (
	"strings"
	"time"
)

// Time is a time.Time that can be unmarshalled from an ISO 8601 string.
type Time time.Time

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (t *Time) UnmarshalJSON(data []byte) error {
	parsed, err := time.Parse(time.RFC3339Nano, string(strings.Trim(string(data), `Z"`)+"Z"))
	if err != nil {
		return err
	}
	*t = Time(parsed)
	return nil
}

// MarshalJSON implements the [json.Marshaler] interface.
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.String() + `"`), nil
}

// String returns the time in RFC3339Nano format.
func (t Time) String() string {
	return time.Time(t).UTC().Format(time.RFC3339Nano)
}
