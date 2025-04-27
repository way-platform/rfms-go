package rfms

// Trigger This description is placed here due to limitations of describing references in OpenAPI
//
//	Property __driverId__:
//	The driver id of driver. (independant whether it is driver or Co-driver)
//	This is only set if the TriggerType = DRIVER_LOGIN, DRIVER_LOGOUT, DRIVER_1_WORKING_STATE_CHANGED or DRIVER_2_WORKING_STATE_CHANGED
//	For DRIVER_LOGIN it is the id of the driver that logged in
//	For DRIVER_LOGOUT it is the id of the driver that logged out
//	For DRIVER_1_WORKING_STATE_CHANGED it is the id of driver 1
//	For DRIVER_2_WORKING_STATE_CHANGED it is the id of driver 2
//	Property __tellTaleInfo__:
//	The tell tale(s) that triggered this message.
//	This is only set if the TriggerType = TELL_TALE
type Trigger struct {
	// ChargingConnectionStatusInfo Additional information can be provided if the trigger type is BATTERY_PACK_CHARGING_CONNECTION_STATUS_CHANGE.
	ChargingConnectionStatusInfo *ChargingConnectionStatusInfo `json:"chargingConnectionStatusInfo,omitempty"`

	// ChargingStatusInfo Additional information can be provided if the trigger type is BATTERY_PACK_CHARGING_STATUS_CHANGE.
	ChargingStatusInfo *ChargingStatusInfo `json:"chargingStatusInfo,omitempty"`

	// Context The context defines if this is part of the standard or OEM specific. rFMS standard values VOLVO TRUCKS, SCANIA, DAIMLER, IVECO, DAF, MAN, RENAULT TRUCKS, VDL, VOLVO BUSES, IVECO BUS, IRISBUS If the Trigger is defined in the rFMS standard, the Context = RFMS
	Context  string    `json:"context"`
	DriverId *DriverID `json:"driverId,omitempty"`

	// PtoId The id of a PTO. This is only set if the TriggerType = PTO_ENABLED or PTO_DISABLED
	PtoId        *string   `json:"ptoId,omitempty"`
	TellTaleInfo *TellTale `json:"tellTaleInfo,omitempty"`

	// TriggerInfo Additional TriggerInfo content for OEM specific triggers E.g. TRAILER_ATTACHED_TRIGGER [id of trailer]
	TriggerInfo *[]string `json:"triggerInfo,omitempty"`

	// TriggerType Trigger types for Context=RFMS:
	//  TIMER - Data was sent due to a timer trigger. (Timer value set outside rFMS scope)
	//  IGNITION_ON - Data was sent due to an ignition on
	//  IGNITION_OFF - Data was sent due to an ignition off
	//  PTO_ENABLED - Data was sent due to that a PTO was enabled, will be sent for each PTO that gets enabled
	//  PTO_DISABLED - Data was sent due to that a PTO was disabled, will be sent for each PTO that gets disabled.
	//  DRIVER_LOGIN - Data was sent due to a successful driver login.
	//  DRIVER_LOGOUT - Data was sent due to a driver logout
	//  TELL_TALE - Data was sent due to that at least one tell tale changed state
	//  ENGINE_ON - Data was sent due to an engine on. For electric motor crank is on
	//  ENGINE_OFF - Data was sent due to an engine off. For electric motor crank is off
	//  DRIVER_1_WORKING_STATE_CHANGED - Data was sent due to that driver 1 changed working state
	//  DRIVER_2_WORKING_STATE_CHANGED - Data was sent due to that driver 2 changed working state
	//  DISTANCE_TRAVELLED - Data was sent due to that a set distance was travelled. (Distance set outside rFMS scope)
	//  FUEL_TYPE_CHANGE - Data was sent due to that the type of fuel currently being utilized by the vehicle changed
	//  PARKING_BRAKE_SWITCH_CHANGE - Data was sent due to that the parking brake state has changed
	//  BATTERY_PACK_CHARGING_STATUS_CHANGE - Data was sent due to a change in the battery pack charging status.
	//  BATTERY_PACK_CHARGING_CONNECTION_STATUS_CHANGE - Data was sent due to a change in the battery pack charging connection status.
	//  TRAILER_CONNECTED - One or several trailers were connected
	//  TRAILER_DISCONNECTED - One or several trailers were disconnected
	TriggerType string `json:"triggerType"`
}

// TellTale defines model for TellTale.
type TellTale struct {
	// State The current state of the tell tale.
	State    string `json:"state"`
	TellTale string `json:"tellTale"`
	// OemTellTale The OemTellTale is only set when the TellTale == OEM_SPECIFIC_TELL_TALE. This is an OEM specific string defining a tell tale in the OEM context.
	OemTellTale *string `json:"oemTellTale,omitempty"`
}

type ChargingConnectionStatusInfo struct {
	// Event CONNECTING - Vehicle is being connected to a charger
	//  CONNECTED - Vehicle is connected to a charger
	//  DISCONNECTING - Vehicle is being disconnected from the charger
	//  DISCONNECTED - Vehicle is not connected to a charger
	//  ERROR - An error occurred
	Event *string `json:"event,omitempty"`

	// EventDetail Details regarding the event. Content is OEM specific
	EventDetail *string `json:"eventDetail,omitempty"`
}

type ChargingStatusInfo struct {
	// Event CHARGING_STARTED - Charging has started
	//  CHARGING_COMPLETED - Charging is completed
	//  CHARGING_INTERRUPTED - Charging has been interrupted (no error)
	//  ERROR - An error occurred when charging
	//  ESTIMATED_COMPLETION_TIME_CHANGED - The estimated time for completed charging has changed. (Threshold is outside scope of rFMS)
	//  TIMER - A predefined time has passed since last charge status update. (Frequency is outside the scope of rFMS)
	//  CHARGING_LEVEL - The charging level has reached a predefined level. (Charging levels are outside the scope of rFMS)
	Event *string `json:"event,omitempty"`

	// EventDetail Details regarding the event. Content is OEM specific
	EventDetail *string `json:"eventDetail,omitempty"`
}
