package rfmsv4oapi

import (
	"strconv"
	"strings"
)

// Int is a custom JSON unmarshaller for int values that may be sent as string.
type Int int

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (i *Int) UnmarshalJSON(data []byte) error {
	value, err := strconv.Atoi(strings.Trim(string(data), `"`))
	if err != nil {
		return err
	}
	*i = Int(value)
	return nil
}

// MarshalJSON implements the [json.Marshaler] interface.
func (i Int) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(int(i))), nil
}
