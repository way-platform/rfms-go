package rfms

import (
	"encoding/json"
)

// VehiclePosition contains vehicle position data.
type VehiclePosition struct {
	// Raw response body.
	Raw json.RawMessage `json:"-"`

	// VIN vehicle identification number. See ISO 3779 (17 characters)
	VIN string `json:"vin"`

	// CreatedDateTime When the data was retrieved in the vehicle in iso8601 format.
	CreatedDateTime Time `json:"createdDateTime"`

	// ReceivedDateTime Reception at Server. To be used for handling of "more data available" in iso8601 format.
	ReceivedDateTime Time `json:"receivedDateTime"`

	// GNSSPosition contains the position data.
	GnssPosition *GNSSPosition `json:"gnssPosition,omitempty"`

	// TachographSpeed Tachograph vehicle speed in km/h (Speed of the vehicle registered by the tachograph)
	TachographSpeed *float64 `json:"tachographSpeed,omitempty"`

	// WheelBasedSpeed Wheel-Based Vehicle Speed in km/h (Speed of the vehicle as calculated from wheel or tailshaft speed)
	WheelBasedSpeed *float64 `json:"wheelBasedSpeed,omitempty"`

	// TriggerType This description is placed here due to limitations of describing references in OpenAPI
	//  Property __driverId__:
	//  The driver id of driver. (independant whether it is driver or Co-driver)
	//  This is only set if the TriggerType = DRIVER_LOGIN, DRIVER_LOGOUT, DRIVER_1_WORKING_STATE_CHANGED or DRIVER_2_WORKING_STATE_CHANGED
	//  For DRIVER_LOGIN it is the id of the driver that logged in
	//  For DRIVER_LOGOUT it is the id of the driver that logged out
	//  For DRIVER_1_WORKING_STATE_CHANGED it is the id of driver 1
	//  For DRIVER_2_WORKING_STATE_CHANGED it is the id of driver 2
	//  Property __tellTaleInfo__:
	//  The tell tale(s) that triggered this message.
	//  This is only set if the TriggerType = TELL_TALE
	TriggerType Trigger `json:"triggerType"`
}
