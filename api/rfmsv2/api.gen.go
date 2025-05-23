// Package rfmsv2 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.2-0.20250407101136-1883b30943ae DO NOT EDIT.
package rfmsv2

import (
	"time"
)

// Defines values for AlternatorStatus.
const (
	AlternatorStatusCHARGING     AlternatorStatus = "CHARGING"
	AlternatorStatusERROR        AlternatorStatus = "ERROR"
	AlternatorStatusNOTAVAILABLE AlternatorStatus = "NOT_AVAILABLE"
	AlternatorStatusNOTCHARGING  AlternatorStatus = "NOT_CHARGING"
)

// Defines values for DoorEnabledStatus.
const (
	DoorEnabledStatusDISABLED     DoorEnabledStatus = "DISABLED"
	DoorEnabledStatusENABLED      DoorEnabledStatus = "ENABLED"
	DoorEnabledStatusERROR        DoorEnabledStatus = "ERROR"
	DoorEnabledStatusNOTAVAILABLE DoorEnabledStatus = "NOT_AVAILABLE"
)

// Defines values for DoorLockStatus.
const (
	DoorLockStatusERROR        DoorLockStatus = "ERROR"
	DoorLockStatusLOCKED       DoorLockStatus = "LOCKED"
	DoorLockStatusNOTAVAILABLE DoorLockStatus = "NOT_AVAILABLE"
	DoorLockStatusUNLOCKED     DoorLockStatus = "UNLOCKED"
)

// Defines values for DoorOpenStatus.
const (
	DoorOpenStatusCLOSED       DoorOpenStatus = "CLOSED"
	DoorOpenStatusERROR        DoorOpenStatus = "ERROR"
	DoorOpenStatusNOTAVAILABLE DoorOpenStatus = "NOT_AVAILABLE"
	DoorOpenStatusOPEN         DoorOpenStatus = "OPEN"
)

// Defines values for DriverAuthenticationType.
const (
	DriverAuthenticationTypeCOMPANYCARD       DriverAuthenticationType = "COMPANY_CARD"
	DriverAuthenticationTypeCONTROLCARD       DriverAuthenticationType = "CONTROL_CARD"
	DriverAuthenticationTypeDRIVERCARD        DriverAuthenticationType = "DRIVER_CARD"
	DriverAuthenticationTypeMANUFACTURINGCARD DriverAuthenticationType = "MANUFACTURING_CARD"
	DriverAuthenticationTypeMOTIONSENSOR      DriverAuthenticationType = "MOTION_SENSOR"
	DriverAuthenticationTypeRESERVED          DriverAuthenticationType = "RESERVED"
	DriverAuthenticationTypeVEHICLEUNIT       DriverAuthenticationType = "VEHICLE_UNIT"
)

// Defines values for DriverWorkingState.
const (
	DriverWorkingStateDRIVE           DriverWorkingState = "DRIVE"
	DriverWorkingStateDRIVERAVAILABLE DriverWorkingState = "DRIVER_AVAILABLE"
	DriverWorkingStateERROR           DriverWorkingState = "ERROR"
	DriverWorkingStateNOTAVAILABLE    DriverWorkingState = "NOT_AVAILABLE"
	DriverWorkingStateREST            DriverWorkingState = "REST"
	DriverWorkingStateWORK            DriverWorkingState = "WORK"
)

// Defines values for Status2OfDoors.
const (
	Status2OfDoorsALLDOORSDISABLED      Status2OfDoors = "ALL_DOORS_DISABLED"
	Status2OfDoorsATLEASTONEDOORENABLED Status2OfDoors = "AT_LEAST_ONE_DOOR_ENABLED"
	Status2OfDoorsERROR                 Status2OfDoors = "ERROR"
	Status2OfDoorsNOTAVAILABLE          Status2OfDoors = "NOT_AVAILABLE"
)

// Defines values for TellTaleState.
const (
	TellTaleStateINFO         TellTaleState = "INFO"
	TellTaleStateNOTAVAILABLE TellTaleState = "NOT_AVAILABLE"
	TellTaleStateOFF          TellTaleState = "OFF"
	TellTaleStateRED          TellTaleState = "RED"
	TellTaleStateRESERVED4    TellTaleState = "RESERVED_4"
	TellTaleStateRESERVED5    TellTaleState = "RESERVED_5"
	TellTaleStateRESERVED6    TellTaleState = "RESERVED_6"
	TellTaleStateYELLOW       TellTaleState = "YELLOW"
)

// Defines values for TellTaleType.
const (
	TellTaleTypeABSTRAILER                      TellTaleType = "ABS_TRAILER"
	TellTaleTypeACC                             TellTaleType = "ACC"
	TellTaleTypeADBLUELEVEL                     TellTaleType = "ADBLUE_LEVEL"
	TellTaleTypeADVANCEDEMERGENCYBREAKING       TellTaleType = "ADVANCED_EMERGENCY_BREAKING"
	TellTaleTypeAIRBAG                          TellTaleType = "AIRBAG"
	TellTaleTypeAIRFILTERCLOGGED                TellTaleType = "AIR_FILTER_CLOGGED"
	TellTaleTypeANTILOCKBRAKEFAILURE            TellTaleType = "ANTI_LOCK_BRAKE_FAILURE"
	TellTaleTypeARTICULATION                    TellTaleType = "ARTICULATION"
	TellTaleTypeAUXILLARYAIRPRESSURE            TellTaleType = "AUXILLARY_AIR_PRESSURE"
	TellTaleTypeBATTERYCHARGINGCONDITION        TellTaleType = "BATTERY_CHARGING_CONDITION"
	TellTaleTypeBRAKELIGHTS                     TellTaleType = "BRAKE_LIGHTS"
	TellTaleTypeBRAKEMALFUNCTION                TellTaleType = "BRAKE_MALFUNCTION"
	TellTaleTypeBUSSTOPBRAKE                    TellTaleType = "BUS_STOP_BRAKE"
	TellTaleTypeCOOLINGAIRCONDITIONING          TellTaleType = "COOLING_AIR_CONDITIONING"
	TellTaleTypeEBS                             TellTaleType = "EBS"
	TellTaleTypeEBSTRAILER12                    TellTaleType = "EBS_TRAILER_1_2"
	TellTaleTypeENGINECOMPARTMENTTEMPERATURE    TellTaleType = "ENGINE_COMPARTMENT_TEMPERATURE"
	TellTaleTypeENGINECOOLANTLEVEL              TellTaleType = "ENGINE_COOLANT_LEVEL"
	TellTaleTypeENGINECOOLANTTEMPERATURE        TellTaleType = "ENGINE_COOLANT_TEMPERATURE"
	TellTaleTypeENGINEEMISSIONFAILURE           TellTaleType = "ENGINE_EMISSION_FAILURE"
	TellTaleTypeENGINEMILINDICATOR              TellTaleType = "ENGINE_MIL_INDICATOR"
	TellTaleTypeENGINEOIL                       TellTaleType = "ENGINE_OIL"
	TellTaleTypeENGINEOILLEVEL                  TellTaleType = "ENGINE_OIL_LEVEL"
	TellTaleTypeENGINEOILTEMPERATURE            TellTaleType = "ENGINE_OIL_TEMPERATURE"
	TellTaleTypeESCINDICATOR                    TellTaleType = "ESC_INDICATOR"
	TellTaleTypeESCSWITCHEDOFF                  TellTaleType = "ESC_SWITCHED_OFF"
	TellTaleTypeFRONTFOGLIGHT                   TellTaleType = "FRONT_FOG_LIGHT"
	TellTaleTypeFUELFILTERDIFFPRESSURE          TellTaleType = "FUEL_FILTER_DIFF_PRESSURE"
	TellTaleTypeFUELLEVEL                       TellTaleType = "FUEL_LEVEL"
	TellTaleTypeGENERALFAILURE                  TellTaleType = "GENERAL_FAILURE"
	TellTaleTypeHATCHOPEN                       TellTaleType = "HATCH_OPEN"
	TellTaleTypeHAZARDWARNING                   TellTaleType = "HAZARD_WARNING"
	TellTaleTypeHEIGHTCONTROL                   TellTaleType = "HEIGHT_CONTROL"
	TellTaleTypeHIGHBEAMMAINBEAM                TellTaleType = "HIGH_BEAM_MAIN_BEAM"
	TellTaleTypeKNEELING                        TellTaleType = "KNEELING"
	TellTaleTypeLANEDEPARTUREINDICATOR          TellTaleType = "LANE_DEPARTURE_INDICATOR"
	TellTaleTypeLANEDEPARTUREWARNINGSWITCHEDOFF TellTaleType = "LANE_DEPARTURE_WARNING_SWITCHED_OFF"
	TellTaleTypeLOWBEAMDIPPEDBEAM               TellTaleType = "LOW_BEAM_DIPPED_BEAM"
	TellTaleTypeLOWERING                        TellTaleType = "LOWERING"
	TellTaleTypeOEMSPECIFICTELLTALE             TellTaleType = "OEM_SPECIFIC_TELL_TALE"
	TellTaleTypePARKINGBRAKE                    TellTaleType = "PARKING_BRAKE"
	TellTaleTypePARKINGHEATER                   TellTaleType = "PARKING_HEATER"
	TellTaleTypePOSITIONLIGHTS                  TellTaleType = "POSITION_LIGHTS"
	TellTaleTypePRAMREQUEST                     TellTaleType = "PRAM_REQUEST"
	TellTaleTypePROVISIONINGHANDICAPPEDPERSON   TellTaleType = "PROVISIONING_HANDICAPPED_PERSON"
	TellTaleTypeRAISING                         TellTaleType = "RAISING"
	TellTaleTypeREARFOGLIGHT                    TellTaleType = "REAR_FOG_LIGHT"
	TellTaleTypeRETARDER                        TellTaleType = "RETARDER"
	TellTaleTypeSEATBELT                        TellTaleType = "SEAT_BELT"
	TellTaleTypeSERVICECALLFORMAINTENANCE       TellTaleType = "SERVICE_CALL_FOR_MAINTENANCE"
	TellTaleTypeSTEERINGFAILURE                 TellTaleType = "STEERING_FAILURE"
	TellTaleTypeSTEERINGFLUIDLEVEL              TellTaleType = "STEERING_FLUID_LEVEL"
	TellTaleTypeSTOPREQUEST                     TellTaleType = "STOP_REQUEST"
	TellTaleTypeTACHOGRAPHINDICATOR             TellTaleType = "TACHOGRAPH_INDICATOR"
	TellTaleTypeTIREMALFUNCTION                 TellTaleType = "TIRE_MALFUNCTION"
	TellTaleTypeTRAILERCONNECTED                TellTaleType = "TRAILER_CONNECTED"
	TellTaleTypeTRANSMISSIONFLUIDTEMPERATURE    TellTaleType = "TRANSMISSION_FLUID_TEMPERATURE"
	TellTaleTypeTRANSMISSIONMALFUNCTION         TellTaleType = "TRANSMISSION_MALFUNCTION"
	TellTaleTypeTURNSIGNALS                     TellTaleType = "TURN_SIGNALS"
	TellTaleTypeWINDSCREENWASHERFLUID           TellTaleType = "WINDSCREEN_WASHER_FLUID"
	TellTaleTypeWORNBRAKELININGS                TellTaleType = "WORN_BRAKE_LININGS"
)

// Defines values for TriggerType.
const (
	TriggerTypeDISTANCETRAVELLED          TriggerType = "DISTANCE_TRAVELLED"
	TriggerTypeDRIVER1WORKINGSTATECHANGED TriggerType = "DRIVER_1_WORKING_STATE_CHANGED"
	TriggerTypeDRIVER2WORKINGSTATECHANGED TriggerType = "DRIVER_2_WORKING_STATE_CHANGED"
	TriggerTypeDRIVERLOGIN                TriggerType = "DRIVER_LOGIN"
	TriggerTypeDRIVERLOGOUT               TriggerType = "DRIVER_LOGOUT"
	TriggerTypeENGINEOFF                  TriggerType = "ENGINE_OFF"
	TriggerTypeENGINEON                   TriggerType = "ENGINE_ON"
	TriggerTypeIGNITIONOFF                TriggerType = "IGNITION_OFF"
	TriggerTypeIGNITIONON                 TriggerType = "IGNITION_ON"
	TriggerTypePTODISABLED                TriggerType = "PTO_DISABLED"
	TriggerTypePTOENABLED                 TriggerType = "PTO_ENABLED"
	TriggerTypeTELLTALE                   TriggerType = "TELL_TALE"
	TriggerTypeTIMER                      TriggerType = "TIMER"
)

// AccumulatedData defines model for AccumulatedData.
type AccumulatedData struct {
	AccelerationClass                     *FromToClasses `json:"AccelerationClass,omitempty"`
	AccelerationDuringBrakeClass          *FromToClasses `json:"AccelerationDuringBrakeClass,omitempty"`
	AccelerationPedalPositionClass        *FromToClasses `json:"AccelerationPedalPositionClass,omitempty"`
	BrakePedalCounterSpeedOverZero        *int64         `json:"BrakePedalCounterSpeedOverZero,omitempty"`
	ChairliftCounter                      *int64         `json:"ChairliftCounter,omitempty"`
	CurrentGearClass                      *LabelClasses  `json:"CurrentGearClass,omitempty"`
	DistanceBrakePedalActiveSpeedOverZero *int64         `json:"DistanceBrakePedalActiveSpeedOverZero,omitempty"`
	DistanceCruiseControlActive           *int64         `json:"DistanceCruiseControlActive,omitempty"`
	DrivingWithoutTorqueClass             *LabelClasses  `json:"DrivingWithoutTorqueClass,omitempty"`
	DurationCruiseControlActive           *int64         `json:"DurationCruiseControlActive,omitempty"`
	DurationWheelbaseSpeedOverZero        *int64         `json:"DurationWheelbaseSpeedOverZero,omitempty"`
	DurationWheelbaseSpeedZero            *int64         `json:"DurationWheelbaseSpeedZero,omitempty"`
	EngineSpeedClass                      *FromToClasses `json:"EngineSpeedClass,omitempty"`
	EngineTorqueAtCurrentSpeedClass       *FromToClasses `json:"EngineTorqueAtCurrentSpeedClass,omitempty"`
	EngineTorqueClass                     *FromToClasses `json:"EngineTorqueClass,omitempty"`
	FuelConsumptionCruiseControlActive    *int64         `json:"FuelConsumptionCruiseControlActive,omitempty"`
	FuelWheelbaseSpeedOverZero            *int64         `json:"FuelWheelbaseSpeedOverZero,omitempty"`
	FuelWheelbaseSpeedZero                *int64         `json:"FuelWheelbaseSpeedZero,omitempty"`
	HighAccelerationClass                 *FromToClasses `json:"HighAccelerationClass,omitempty"`
	KneelingCounter                       *int64         `json:"KneelingCounter,omitempty"`
	PramRequestCounter                    *int64         `json:"PramRequestCounter,omitempty"`
	PtoActiveClass                        *LabelClasses  `json:"PtoActiveClass,omitempty"`
	RetarderTorqueClass                   *FromToClasses `json:"RetarderTorqueClass,omitempty"`
	SelectedGearClass                     *LabelClasses  `json:"SelectedGearClass,omitempty"`
	StopRequestCounter                    *int64         `json:"StopRequestCounter,omitempty"`
	VehicleSpeedClass                     *FromToClasses `json:"VehicleSpeedClass,omitempty"`
}

// AlternatorInfo defines model for AlternatorInfo.
type AlternatorInfo struct {
	AlternatorNumber *int64           `json:"AlternatorNumber,omitempty"`
	AlternatorStatus AlternatorStatus `json:"AlternatorStatus,omitempty"`
}

// AlternatorStatus defines model for AlternatorStatus.
type AlternatorStatus string

// Date defines model for Date.
type Date struct {
	Day   *int32 `json:"Day,omitempty"`
	Month *int32 `json:"Month,omitempty"`
	Year  *int32 `json:"Year,omitempty"`
}

// DoorEnabledStatus defines model for DoorEnabledStatus.
type DoorEnabledStatus string

// DoorLockStatus defines model for DoorLockStatus.
type DoorLockStatus string

// DoorOpenStatus defines model for DoorOpenStatus.
type DoorOpenStatus string

// DoorStatus defines model for DoorStatus.
type DoorStatus struct {
	DoorEnabledStatus DoorEnabledStatus `json:"DoorEnabledStatus,omitempty"`
	DoorLockStatus    DoorLockStatus    `json:"DoorLockStatus,omitempty"`
	DoorNumber        *int32            `json:"DoorNumber,omitempty"`
	DoorOpenStatus    DoorOpenStatus    `json:"DoorOpenStatus,omitempty"`
}

// DriverAuthenticationType defines model for DriverAuthenticationType.
type DriverAuthenticationType string

// DriverID defines model for DriverId.
type DriverID struct {
	OEMDriverIdentification   *OEMDriverIdentification   `json:"OemDriverIdentification,omitempty"`
	TachoDriverIdentification *TachoDriverIdentification `json:"TachoDriverIdentification,omitempty"`
}

// DriverWorkingState defines model for DriverWorkingState.
type DriverWorkingState string

// FromToClass defines model for FromToClass.
type FromToClass struct {
	From        *float64 `json:"From,omitempty"`
	Meters      *int64   `json:"Meters,omitempty"`
	MilliLitres *int64   `json:"MilliLitres,omitempty"`
	Seconds     *int64   `json:"Seconds,omitempty"`
	To          *float64 `json:"To,omitempty"`
}

// FromToClasses defines model for FromToClasses.
type FromToClasses struct {
	Value []FromToClass `json:"Value,omitempty"`
}

// GNSSPosition defines model for GNSSPosition.
type GNSSPosition struct {
	Altitude         *int32    `json:"Altitude,omitempty"`
	Heading          *int32    `json:"Heading,omitempty"`
	Latitude         *float64  `json:"Latitude,omitempty"`
	Longitude        *float64  `json:"Longitude,omitempty"`
	PositionDateTime time.Time `json:"PositionDateTime,omitzero"`
	Speed            *float64  `json:"Speed,omitempty"`
}

// LabelClass defines model for LabelClass.
type LabelClass struct {
	Label       string `json:"Label,omitempty"`
	Meters      *int64 `json:"Meters,omitempty"`
	MilliLitres *int64 `json:"MilliLitres,omitempty"`
	Seconds     *int64 `json:"Seconds,omitempty"`
}

// LabelClasses defines model for LabelClasses.
type LabelClasses struct {
	Value []LabelClass `json:"Value,omitempty"`
}

// OEMDriverIdentification defines model for OemDriverIdentification.
type OEMDriverIdentification struct {
	IDType                  string `json:"IdType,omitempty"`
	OEMDriverIdentification string `json:"OemDriverIdentification,omitempty"`
}

// SnapshotData defines model for SnapshotData.
type SnapshotData struct {
	AmbientAirTemperature *float64           `json:"AmbientAirTemperature,omitempty"`
	CatalystFuelLevel     *float64           `json:"CatalystFuelLevel,omitempty"`
	Driver1WorkingState   DriverWorkingState `json:"Driver1WorkingState,omitempty"`
	Driver2ID             *DriverID          `json:"Driver2Id,omitempty"`
	Driver2WorkingState   DriverWorkingState `json:"Driver2WorkingState,omitempty"`
	EngineSpeed           *float64           `json:"EngineSpeed,omitempty"`
	FuelLevel1            *float64           `json:"FuelLevel1,omitempty"`
	GNSSPosition          *GNSSPosition      `json:"GNSSPosition,omitempty"`
	TachographSpeed       *float64           `json:"TachographSpeed,omitempty"`
	WheelBasedSpeed       *float64           `json:"WheelBasedSpeed,omitempty"`
}

// Status2OfDoors defines model for Status2OfDoors.
type Status2OfDoors string

// TachoDriverIdentification defines model for TachoDriverIdentification.
type TachoDriverIdentification struct {
	CardIssuingMemberState        string                   `json:"CardIssuingMemberState,omitempty"`
	CardRenewalIndex              string                   `json:"CardRenewalIndex,omitempty"`
	CardReplacementIndex          string                   `json:"CardReplacementIndex,omitempty"`
	DriverAuthenticationEquipment DriverAuthenticationType `json:"DriverAuthenticationEquipment,omitempty"`
	DriverIdentification          string                   `json:"DriverIdentification,omitempty"`
}

// TellTaleInfo defines model for TellTaleInfo.
type TellTaleInfo struct {
	State    TellTaleState `json:"State,omitempty"`
	TellTale TellTaleType  `json:"TellTale,omitempty"`
}

// TellTaleState defines model for TellTaleState.
type TellTaleState string

// TellTaleType defines model for TellTaleType.
type TellTaleType string

// Trigger defines model for Trigger.
type Trigger struct {
	Context      string         `json:"Context,omitempty"`
	DriverID     *DriverID      `json:"DriverId,omitempty"`
	PtoID        string         `json:"PtoId,omitempty"`
	TellTaleInfo []TellTaleInfo `json:"TellTaleInfo,omitempty"`
	TriggerInfo  []string       `json:"TriggerInfo,omitempty"`
	TriggerType  TriggerType    `json:"TriggerType,omitempty"`
}

// TriggerType defines model for TriggerType.
type TriggerType string

// UptimeData defines model for UptimeData.
type UptimeData struct {
	AlternatorInfo                  []AlternatorInfo `json:"AlternatorInfo,omitempty"`
	BellowPressureFrontAxleLeft     *int64           `json:"BellowPressureFrontAxleLeft,omitempty"`
	BellowPressureFrontAxleRight    *int64           `json:"BellowPressureFrontAxleRight,omitempty"`
	BellowPressureRearAxleLeft      *int64           `json:"BellowPressureRearAxleLeft,omitempty"`
	BellowPressureRearAxleRight     *int64           `json:"BellowPressureRearAxleRight,omitempty"`
	DurationAtLeastOneDoorOpen      *int64           `json:"DurationAtLeastOneDoorOpen,omitempty"`
	EngineCoolantTemperature        *float64         `json:"EngineCoolantTemperature,omitempty"`
	ServiceBrakeAirPressureCircuit1 *int64           `json:"ServiceBrakeAirPressureCircuit1,omitempty"`
	ServiceBrakeAirPressureCircuit2 *int64           `json:"ServiceBrakeAirPressureCircuit2,omitempty"`
	ServiceDistance                 *int64           `json:"ServiceDistance,omitempty"`
	TellTaleInfo                    []TellTaleInfo   `json:"TellTaleInfo,omitempty"`
}

// Vehicle defines model for Vehicle.
type Vehicle struct {
	BodyType            string   `json:"BodyType,omitempty"`
	Brand               string   `json:"Brand,omitempty"`
	ChassisType         string   `json:"ChassisType,omitempty"`
	CustomerVehicleName string   `json:"CustomerVehicleName,omitempty"`
	DoorConfiguration   []int32  `json:"DoorConfiguration,omitempty"`
	EmissionLevel       string   `json:"EmissionLevel,omitempty"`
	GearboxType         string   `json:"GearboxType,omitempty"`
	HasRampOrLift       *bool    `json:"HasRampOrLift,omitempty"`
	Model               string   `json:"Model,omitempty"`
	NoOfAxles           *int32   `json:"NoOfAxles,omitempty"`
	PossibleFuelType    []string `json:"PossibleFuelType,omitempty"`
	ProductionDate      *Date    `json:"ProductionDate,omitempty"`
	TachographType      string   `json:"TachographType,omitempty"`
	TellTaleCode        string   `json:"TellTaleCode,omitempty"`
	TotalFuelTankVolume *int32   `json:"TotalFuelTankVolume,omitempty"`
	Type                string   `json:"Type,omitempty"`
	VIN                 string   `json:"VIN,omitempty"`
}

// VehiclePosition defines model for VehiclePosition.
type VehiclePosition struct {
	CreatedDateTime  time.Time     `json:"CreatedDateTime,omitzero"`
	GNSSPosition     *GNSSPosition `json:"GNSSPosition,omitempty"`
	ReceivedDateTime time.Time     `json:"ReceivedDateTime,omitzero"`
	TachographSpeed  *float64      `json:"TachographSpeed,omitempty"`
	TriggerType      *Trigger      `json:"TriggerType,omitempty"`
	VIN              string        `json:"VIN,omitempty"`
	WheelBasedSpeed  *float64      `json:"WheelBasedSpeed,omitempty"`
}

// VehiclePositionsResponse defines model for VehiclePositionsResponse.
type VehiclePositionsResponse struct {
	MoreDataAvailable     bool              `json:"MoreDataAvailable,omitempty"`
	RequestServerDateTime time.Time         `json:"RequestServerDateTime,omitzero"`
	VehiclePosition       []VehiclePosition `json:"VehiclePosition,omitempty"`
}

// VehicleStatus defines model for VehicleStatus.
type VehicleStatus struct {
	AccumulatedData               *AccumulatedData `json:"AccumulatedData,omitempty"`
	CreatedDateTime               time.Time        `json:"CreatedDateTime,omitzero"`
	DoorStatus                    []DoorStatus     `json:"DoorStatus,omitempty"`
	Driver1ID                     *DriverID        `json:"Driver1Id,omitempty"`
	EngineTotalFuelUsed           *int64           `json:"EngineTotalFuelUsed,omitempty"`
	GrossCombinationVehicleWeight *int32           `json:"GrossCombinationVehicleWeight,omitempty"`
	HRTotalVehicleDistance        *int64           `json:"HRTotalVehicleDistance,omitempty"`
	ReceivedDateTime              time.Time        `json:"ReceivedDateTime,omitzero"`
	SnapshotData                  *SnapshotData    `json:"SnapshotData,omitempty"`
	Status2OfDoors                Status2OfDoors   `json:"Status2OfDoors,omitempty"`
	TotalEngineHours              *float64         `json:"TotalEngineHours,omitempty"`
	TriggerType                   *Trigger         `json:"TriggerType,omitempty"`
	UptimeData                    *UptimeData      `json:"UptimeData,omitempty"`
	VIN                           string           `json:"Vin,omitempty"`
}

// VehicleStatusesResponse defines model for VehicleStatusesResponse.
type VehicleStatusesResponse struct {
	MoreDataAvailable     bool            `json:"MoreDataAvailable,omitempty"`
	RequestServerDateTime time.Time       `json:"RequestServerDateTime,omitzero"`
	VehicleStatus         []VehicleStatus `json:"VehicleStatus,omitempty"`
}

// VehiclesResponse defines model for VehiclesResponse.
type VehiclesResponse struct {
	MoreDataAvailable bool      `json:"MoreDataAvailable,omitempty"`
	Vehicle           []Vehicle `json:"Vehicle,omitempty"`
}
