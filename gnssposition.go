package rfms

// GNSSPosition contains the position data.
type GNSSPosition struct {
	// Altitude The altitude of the vehicle. Where 0 is sea level, negative values below sealevel and positive above sealevel. Unit in meters.
	Altitude *int `json:"altitude,omitempty"`

	// Heading The direction of the vehicle (0-359)
	Heading any `json:"heading,omitempty"`

	// Latitude Latitude (WGS84 based)
	Latitude float64 `json:"latitude"`

	// Longitude Longitude (WGS84 based)
	Longitude float64 `json:"longitude"`

	// PositionDateTime The time of the position data in iso8601 format.
	PositionDateTime Time `json:"positionDateTime"`

	// Speed The GNSS(e.g. GPS)-speed in km/h
	Speed *float64 `json:"speed,omitempty"`
}
