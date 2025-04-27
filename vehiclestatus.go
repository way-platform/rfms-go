package rfms

import (
	"encoding/json"
	"time"
)

type VehicleStatus struct {
	// Raw response body.
	Raw json.RawMessage `json:"-"`

	// VIN vehicle identification number. See ISO 3779 (17 characters)
	VIN string `json:"vin"`

	AccumulatedData *AccumulatedData `json:"accumulatedData,omitempty"`

	// CreatedDateTime When the data was retrieved in the vehicle in iso8601 format.
	CreatedDateTime Time `json:"createdDateTime"`

	// DoorStatus Individual status for each door. Bus specific parameter
	DoorStatus *[]struct {
		DoorEnabledStatus *string `json:"DoorEnabledStatus,omitempty"`
		DoorLockStatus    *string `json:"DoorLockStatus,omitempty"`
		DoorNumber        *int    `json:"DoorNumber,omitempty"`
		DoorOpenStatus    *string `json:"DoorOpenStatus,omitempty"`
	} `json:"doorStatus,omitempty"`
	Driver1Id *DriverID `json:"driver1Id,omitempty"`

	// EngineTotalFuelUsed The total fuel the vehicle has used during its lifetime in MilliLitres. At least one of engineTotalFuelUsed, totalFuelUsedGaseous or totalElectricEnergyUsed is mandatory.
	EngineTotalFuelUsed *int64 `json:"engineTotalFuelUsed,omitempty"`

	// GrossCombinationVehicleWeight The full vehicle weight in kg
	GrossCombinationVehicleWeight *int `json:"grossCombinationVehicleWeight,omitempty"`

	// HrTotalVehicleDistance Accumulated distance travelled by the vehicle during its operation in meter
	HrTotalVehicleDistance *int64 `json:"hrTotalVehicleDistance,omitempty"`

	// ReceivedDateTime Reception at Server. To be used for handling of "more data available" in iso8601 format.
	ReceivedDateTime Time          `json:"receivedDateTime"`
	SnapshotData     *SnapshotData `json:"snapshotData,omitempty"`

	// Status2OfDoors Composite indication of all bus door statuses. Bus specific parameter
	Status2OfDoors *string `json:"status2OfDoors,omitempty"`

	// TotalElectricEnergyUsed Total electric energy consumed by the vehicle, excluding when plugged in (vehicle coupler) for charging, (incl. motor, PTO, cooling, etc.) in watt hours. Recuperation is subtracted from the value.  At least one of engineTotalFuelUsed, totalFuelUsedGaseous or totalElectricEnergyUsed is mandatory.
	TotalElectricEnergyUsed *int64 `json:"totalElectricEnergyUsed,omitempty"`

	// TotalElectricMotorHours The total hours the electric motor is ready for propulsion (i.e. crank mode). At least one of totalEngineHours or totalElectricMotorHours is mandatory
	TotalElectricMotorHours *float64 `json:"totalElectricMotorHours,omitempty"`

	// TotalEngineHours The total hours of operation for the vehicle combustion engine. At least one of totalEngineHours or totalElectricMotorHours is Mandatory
	TotalEngineHours *float64 `json:"totalEngineHours,omitempty"`

	// TotalFuelUsedGaseous Total fuel consumed in kg (trip drive fuel + trip PTO governor moving fuel + trip PTO governor non-moving fuel + trip idle fuel) over the life of the engine. At least one of engineTotalFuelUsed, totalFuelUsedGaseous or totalElectricEnergyUsed is mandatory.
	TotalFuelUsedGaseous *int64 `json:"totalFuelUsedGaseous,omitempty"`

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
	TriggerType Trigger     `json:"triggerType"`
	UptimeData  *UptimeData `json:"uptimeData,omitempty"`
}

// AccumulatedData defines model for AccumulatedData.
type AccumulatedData struct {
	// AccelerationClass In m/s2 Minimum 13 classes. ], -1.1] ]-1.1, -0.9] ]-0.9, -0.7] ]-0.7, -0.5] ]-0.5, -0.3] ]-0.3, -0.1] ]-0.1, 0.1[ [0.1, 0.3[ [0.3, 0.5[ [0.5, 0.7[ [0.7, 0.9[ [0.9, 1.1[ [1.1, [
	AccelerationClass *[]FromToClass `json:"accelerationClass,omitempty"`

	// AccelerationDuringBrakeClass In m/s2 Minimum 13 classes. ], -1.1] ]-1.1, -0.9] ]-0.9, -0.7] ]-0.7, -0.5] ]-0.5, -0.3] ]-0.3, -0.1] ]-0.1, 0.1[ [0.1, 0.3[ [0.3, 0.5[ [0.5, 0.7[ [0.7, 0.9[ [0.9, 1.1[ [1.1, [
	AccelerationDuringBrakeClass *[]FromToClass `json:"accelerationDuringBrakeClass,omitempty"`

	// AccelerationPedalPositionClass In percent. Minimum 5 classes [0, 20[ [20, 40[ [40, 60[ [60, 80[ [80, 100]
	AccelerationPedalPositionClass *[]FromToClass `json:"accelerationPedalPositionClass,omitempty"`

	// BrakePedalCounterSpeedOverZero The total number of times the brake pedal has been used while the vehicle was driving.
	BrakePedalCounterSpeedOverZero *int64 `json:"brakePedalCounterSpeedOverZero,omitempty"`

	// BrakePedalPositionClass In percent. Minimum 5 classes [0, 20[ [20, 40[ [40, 60[ [60, 80[ [80, 100]
	BrakePedalPositionClass *[]FromToClass `json:"brakePedalPositionClass,omitempty"`

	// ChairliftCounter The total number of times the chairlift has been outside the bus. This is mainly used for Buses
	ChairliftCounter *int64 `json:"chairliftCounter,omitempty"`

	// CurrentGearClass The currently used gear One class per gear. Neutral is also a gear. Park is also a gear. This is formatted according to SPN 523, supplied as a decimal value. Example 0 = Neutral, 1 = 1:st gear... This is mainly used for Buses.
	CurrentGearClass *[]Label `json:"currentGearClass,omitempty"`

	// DistanceBrakePedalActiveSpeedOverZero The total distance the vehicle has driven where the brake pedal has been used. Unit Meters.
	DistanceBrakePedalActiveSpeedOverZero *int64 `json:"distanceBrakePedalActiveSpeedOverZero,omitempty"`

	// DistanceCruiseControlActive The distance the vehicle has been driven with cruise control active
	DistanceCruiseControlActive *int64 `json:"distanceCruiseControlActive,omitempty"`

	// DrivingWithoutTorqueClass Driving without torque, with gear (clutch is engaged) Labels DRIVING_WITHOUT_TORQUE
	DrivingWithoutTorqueClass *[]Label `json:"drivingWithoutTorqueClass,omitempty"`

	// DurationCruiseControlActive The time the vehicle has been driven with cruise control active
	DurationCruiseControlActive *int64 `json:"durationCruiseControlActive,omitempty"`

	// DurationWheelbasedSpeedOverZero The time the vehicle speed has been over zero.
	DurationWheelbasedSpeedOverZero *int64 `json:"durationWheelbasedSpeedOverZero,omitempty"`

	// DurationWheelbasedSpeedZero The time the vehicle speed has been equal to zero, in seconds. Engine on (RPM>0 or electic motor in crank mode) and no PTO active
	DurationWheelbasedSpeedZero *int64 `json:"durationWheelbasedSpeedZero,omitempty"`

	// ElectricEnergyAux The electric energy the auxiliary systems have consumed, in watt hours. Auxiliary systems are all consumers except electric motor(s) and PTO(s).
	ElectricEnergyAux *int64 `json:"electricEnergyAux,omitempty"`

	// ElectricEnergyConsumptionDuringCruiseActive The electric energy the vehicle has consumed while driven with cruise control active, in watt-hours.
	ElectricEnergyConsumptionDuringCruiseActive *int64 `json:"electricEnergyConsumptionDuringCruiseActive,omitempty"`

	// ElectricEnergyWheelbasedSpeedOverZero The electric energy the vehicle has consumed (including recuperation) while the vehicle speed has been over zero. Electric motor is in crank mode. Unit in watt-hours.
	ElectricEnergyWheelbasedSpeedOverZero *int64 `json:"electricEnergyWheelbasedSpeedOverZero,omitempty"`

	// ElectricEnergyWheelbasedSpeedZero The electric energy the vehicle has consumed while the vehicle speed has been equal to zero. Electric motor is in crank mode and no PTO active. Unit in watt-hours.
	ElectricEnergyWheelbasedSpeedZero *int64 `json:"electricEnergyWheelbasedSpeedZero,omitempty"`

	// ElectricMotorTorqueAtCurrentSpeedClass In percent (Engine Percent Load At Current Speed). Minimum 10 classes [0, 10[ [10, 20[ [20, 30[ [30, 40[ [40, 50[ [50, 60[ [60, 70[ [70, 80[ [80, 90[ [90, 100]
	ElectricMotorTorqueAtCurrentSpeedClass *[]FromToClassElectrical `json:"electricMotorTorqueAtCurrentSpeedClass,omitempty"`

	// ElectricMotorTorqueClass In percent (Actual Engine-Percent Torque). Minimum 10 classes [0, 10[ [10, 20[ [20, 30[ [30, 40[ [40, 50[ [50, 60[ [60, 70[ [70, 80[ [80, 90[ [90, 100]
	ElectricMotorTorqueClass *[]FromToClassElectrical `json:"electricMotorTorqueClass,omitempty"`

	// ElectricPowerRecuperationClass Classes refer to the recuperated electric power in kilowatt Minimum 11 classes [0, 100[ [100, 200[ [200, 300[ ... [900, 1000[ [1000, [
	ElectricPowerRecuperationClass *[]FromToClassElectrical `json:"electricPowerRecuperationClass,omitempty"`

	// EngineSpeedClass Classes refer to the RPM of the combustion engine. Only mandatory if the vehicle has a combustion engine for propulsion. Minimum 10 classes [0, 400[ [400, 800[ [800, 1200[ [1200, 1600[ [1600, 2000[ [2000, 2400[ [2400, 2800[ [2800, 3200[ [3200, 3600[ [3600, [ Note: Engine on (RPM>0 or electric motor in crank mode)
	EngineSpeedClass *[]FromToClass `json:"engineSpeedClass,omitempty"`

	// EngineTorqueAtCurrentSpeedClass In percent based on EEC2 value (Engine Percent Load At Current Speed). Minimum 10 classes [0, 10[ [10, 20[ [20, 30[ [30, 40[ [40, 50[ [50, 60[ [60, 70[ [70, 80[ [80, 90[ [90, 100]
	EngineTorqueAtCurrentSpeedClass *[]FromToClassCombustion `json:"engineTorqueAtCurrentSpeedClass,omitempty"`

	// EngineTorqueClass In percent based on EEC1 value (Actual Engine-Percent Torque). Minimum 10 classes [0, 10[ [10, 20[ [20, 30[ [30, 40[ [40, 50[ [50, 60[ [60, 70[ [70, 80[ [80, 90[ [90, 100]
	EngineTorqueClass *[]FromToClassCombustion `json:"engineTorqueClass,omitempty"`

	// FuelConsumptionDuringCruiseActive The fuel the vehicle has consumed while driven with cruise control active, in millilitres
	FuelConsumptionDuringCruiseActive *int64 `json:"fuelConsumptionDuringCruiseActive,omitempty"`

	// FuelConsumptionDuringCruiseActiveGaseous The gas the vehicle has consumed while driven with cruise control active, in kilograms.
	FuelConsumptionDuringCruiseActiveGaseous *int64 `json:"fuelConsumptionDuringCruiseActiveGaseous,omitempty"`

	// FuelWheelbasedSpeedOverZero The fuel the vehicle has consumed while the vehicle speed has been over zero. Engine on (RPM>0). Unit in millilitres.
	FuelWheelbasedSpeedOverZero *int64 `json:"fuelWheelbasedSpeedOverZero,omitempty"`

	// FuelWheelbasedSpeedOverZeroGaseous The gas the vehicle has consumed while the vehicle speed has been over zero. Engine on (RPM>0). Unit in kilograms.
	FuelWheelbasedSpeedOverZeroGaseous *int64 `json:"fuelWheelbasedSpeedOverZeroGaseous,omitempty"`

	// FuelWheelbasedSpeedZero The fuel the vehicle has consumed while the vehicle speed has been equal to zero. Engine on (RPM>0) and no PTO active. Unit in millilitres.
	FuelWheelbasedSpeedZero *int64 `json:"fuelWheelbasedSpeedZero,omitempty"`

	// FuelWheelbasedSpeedZeroGaseous The gas the vehicle has consumed while the vehicle speed has been equal to zero. Engine on (RPM>0) and no PTO active. Unit in kilograms.
	FuelWheelbasedSpeedZeroGaseous *int64 `json:"fuelWheelbasedSpeedZeroGaseous,omitempty"`

	// HighAccelerationClass In m/s2 Minimum 11 classes ], -3.0] ]-3.0, -2.5] ]-2.5, -2.0] ]-2.0, -1.5] ]-1.5, -1.1] ]-1.1, 1.1[ [1.1, 1.5[ [1.5, 2.0[ [2.0, 2.5[ [2.5, 3.0[ [3.0, [
	HighAccelerationClass *[]FromToClass `json:"highAccelerationClass,omitempty"`

	// KneelingCounter The total number of times the bus has knelt.
	KneelingCounter *int64 `json:"kneelingCounter,omitempty"`

	// PramRequestCounter The total number of pram requests made. This is mainly used for Buses
	PramRequestCounter *int64 `json:"pramRequestCounter,omitempty"`

	// PtoActiveClass Label WHEELBASED_SPEED_ZERO
	//  At least one PTO active during wheelbased speed=0
	//  Counters for time (seconds) and consumption (millilitres, kilograms, watt-hours)
	//  Label WHEELBASED_SPEED_OVER_ZERO
	//  At least one PTO active during wheelbased speed>0
	//  Counters for time (seconds), distance (meter) and consumption (millilitres, kilograms, watt-hours)
	PtoActiveClass *[]Label `json:"ptoActiveClass,omitempty"`

	// RetarderTorqueClass In percent (how the retarder is used as a positive value). Minimum 5 classes ]0, 20[ [20, 40[ [40, 60[ [60, 80[ [80, 100]
	RetarderTorqueClass *[]FromToClass `json:"retarderTorqueClass,omitempty"`

	// SelectedGearClass The currently selected gear One class per gear. Neutral is also a gear. Park is also a gear. This is formatted according to SPN 524, supplied as a decimal value. Example 0 = Neutral, 1 = 1:st gear... This is mainly used for Buses.
	SelectedGearClass *[]Label `json:"selectedGearClass,omitempty"`

	// StopRequestCounter The total number of stop requests made. This is mainly used for Buses
	StopRequestCounter *int64 `json:"stopRequestCounter,omitempty"`

	// VehicleSpeedClass In km/h Minimum 40 classes. [0, 4[ [4, 8[ [8, 12[ [12, 16[ [16, 20[ [20, 24[ ... [156, [ Engine on (RPM>0 or electric motor in crank mode)
	VehicleSpeedClass *[]FromToClass `json:"vehicleSpeedClass,omitempty"`
}

// FromToClass defines model for FromToClass.
type FromToClass struct {
	From        *float64 `json:"from,omitempty"`
	Kilograms   *int64   `json:"kilograms,omitempty"`
	Meters      *int64   `json:"meters,omitempty"`
	MilliLitres *int64   `json:"milliLitres,omitempty"`
	Seconds     *int64   `json:"seconds,omitempty"`
	To          *float64 `json:"to,omitempty"`
	Watthours   *int64   `json:"watthours,omitempty"`
}

// FromToClassCombustion defines model for FromToClassCombustion.
type FromToClassCombustion struct {
	From        *float64 `json:"from,omitempty"`
	Kilograms   *int64   `json:"kilograms,omitempty"`
	Meters      *int64   `json:"meters,omitempty"`
	MilliLitres *int64   `json:"milliLitres,omitempty"`
	Seconds     *int64   `json:"seconds,omitempty"`
	To          *float64 `json:"to,omitempty"`
}

// FromToClassElectrical defines model for FromToClassElectrical.
type FromToClassElectrical struct {
	From      *float64 `json:"from,omitempty"`
	Meters    *int64   `json:"meters,omitempty"`
	Seconds   *int64   `json:"seconds,omitempty"`
	To        *float64 `json:"to,omitempty"`
	Watthours *int64   `json:"watthours,omitempty"`
}

// UptimeData defines model for UptimeData.
type UptimeData struct {
	// AlternatorInfo The alternator status of the up to 4 alternators. Used mainly for buses.
	AlternatorInfo *struct {
		AlternatorNumber *int64  `json:"alternatorNumber,omitempty"`
		AlternatorStatus *string `json:"alternatorStatus,omitempty"`
	} `json:"alternatorInfo,omitempty"`

	// BellowPressureFrontAxleLeft The bellow pressure in the front axle left side in Pascal. Used mainly for buses.
	BellowPressureFrontAxleLeft *int64 `json:"bellowPressureFrontAxleLeft,omitempty"`

	// BellowPressureFrontAxleRight The bellow pressure in the front axle right side in Pascal. Used mainly for buses.
	BellowPressureFrontAxleRight *int64 `json:"bellowPressureFrontAxleRight,omitempty"`

	// BellowPressureRearAxleLeft The bellow pressure in the rear axle left side in Pascal. Used mainly for buses.
	BellowPressureRearAxleLeft *int64 `json:"bellowPressureRearAxleLeft,omitempty"`

	// BellowPressureRearAxleRight The bellow pressure in the rear axle right side in Pascal. Used mainly for buses.
	BellowPressureRearAxleRight *int64 `json:"bellowPressureRearAxleRight,omitempty"`

	// DurationAtLeastOneDoorOpen The total time at least one door has been opened in the bus. (seconds) Used mainly for buses.
	DurationAtLeastOneDoorOpen *int64 `json:"durationAtLeastOneDoorOpen,omitempty"`

	// EngineCoolantTemperature The temperature of the coolant liquid in Celsius
	EngineCoolantTemperature *float64 `json:"engineCoolantTemperature,omitempty"`

	// HvessOutletCoolantTemperature The temperature of the battery pack coolant in Celsius HVESS - High Voltage Energy Storage System
	HvessOutletCoolantTemperature *float64 `json:"hvessOutletCoolantTemperature,omitempty"`

	// HvessTemperature The temperature of the battery pack in Celsius HVESS - High Voltage Energy Storage System
	HvessTemperature *float64 `json:"hvessTemperature,omitempty"`

	// ServiceBrakeAirPressureCircuit1 The air pressure in circuit 1 in Pascal.
	ServiceBrakeAirPressureCircuit1 *int64 `json:"serviceBrakeAirPressureCircuit1,omitempty"`

	// ServiceBrakeAirPressureCircuit2 The air pressure in circuit 2 in Pascal.
	ServiceBrakeAirPressureCircuit2 *int64 `json:"serviceBrakeAirPressureCircuit2,omitempty"`

	// ServiceDistance The distance in meter to the next service
	ServiceDistance *int64 `json:"serviceDistance,omitempty"`

	// TellTaleInfo List of tell tales with the actual status for each tell tale.
	TellTaleInfo []TellTale `json:"tellTaleInfo"`
}

// SnapshotData defines model for SnapshotData.
type SnapshotData struct {
	// AmbientAirTemperature The Ambient air temperature in Celsius
	AmbientAirTemperature *float64 `json:"ambientAirTemperature,omitempty"`

	// BatteryPackChargingConnectionStatus Indicates the charging connection status of the battery pack.
	//  Connecting - A charger is being connected
	//  Connected - A charger is connected
	//  Disconnecting - A charger is being disconnected
	//  Disconnected - No charger is connected
	//  Error - An error occurred when connecting or disconnecting
	//  Not available - Charging connection status is not available
	BatteryPackChargingConnectionStatus *string `json:"batteryPackChargingConnectionStatus,omitempty"`

	// BatteryPackChargingDevice Device used to charge the battery pack. Standard rFMS values taken from ISO 15118 (OEM can have additional values):
	//  ACD - Automatic Connection Device
	//  WPT - Wireless Power Transfer
	//  VEHICLE_COUPLER - manual connection of a flexible cable to an EV
	//  NONE - No device connected
	//  NOT_AVAILABLE - Unknown
	BatteryPackChargingDevice *string `json:"batteryPackChargingDevice,omitempty"`

	// BatteryPackChargingPower Charging power in watts.
	BatteryPackChargingPower *float64 `json:"batteryPackChargingPower,omitempty"`

	// BatteryPackChargingStatus Indicates the charging status of the battery pack. Recuperation is excluded.
	//  Not charging - No charging
	//  Charging - Charging ongoing (AC or DC is unknown)
	//  Charging AC - AC charging ongoing
	//  Charging DC - DC charging ongoing
	//  Error - An error occurred when charging
	//  Not available - Charging status is not available
	BatteryPackChargingStatus *string `json:"batteryPackChargingStatus,omitempty"`

	// CatalystFuelLevel The adblue level percentage
	CatalystFuelLevel *float64 `json:"catalystFuelLevel,omitempty"`

	// Driver1WorkingState Tachograph Working state of the driver
	Driver1WorkingState *string   `json:"driver1WorkingState,omitempty"`
	Driver2Id           *DriverID `json:"driver2Id,omitempty"`

	// Driver2WorkingState Tachograph Working state of the driver
	Driver2WorkingState *string `json:"driver2WorkingState,omitempty"`

	// ElectricMotorSpeed The electric motor speed in rev/min
	ElectricMotorSpeed *float64 `json:"electricMotorSpeed,omitempty"`

	// EngineSpeed The engine (Diesel/gaseous) speed in rev/min
	EngineSpeed *float64 `json:"engineSpeed,omitempty"`

	// EstimatedDistanceToEmpty Estimated distance to empty (tanks and/or battery packs) in meters
	EstimatedDistanceToEmpty *struct {
		// BatteryPack Estimated distance to empty, battery pack, in meters
		BatteryPack *int64 `json:"batteryPack,omitempty"`

		// Fuel Estimated distance to empty, fuel tank, in meters
		Fuel *int64 `json:"fuel,omitempty"`

		// Gas Estimated distance to empty, gas tank, in meters
		Gas *int64 `json:"gas,omitempty"`

		// Total Estimated distance to empty, summarizing fuel, gas and battery in meters
		Total *int64 `json:"total,omitempty"`
	} `json:"estimatedDistanceToEmpty,omitempty"`

	// EstimatedTimeBatteryPackChargingCompleted Estimated time when charging has reached the target level.
	EstimatedTimeBatteryPackChargingCompleted *time.Time `json:"estimatedTimeBatteryPackChargingCompleted,omitempty"`

	// FuelLevel1 The fuel level percentage
	FuelLevel1 *float64 `json:"fuelLevel1,omitempty"`

	// FuelLevel2 Ratio of volume of fuel to the total volume of fuel storage container, in percent. When Fuel Level 2 is not used, Fuel Level 1 represents the total fuel in all fuel storage containers.  When Fuel Level 2 is used, Fuel Level 1 represents the fuel level in the primary or left-side fuel storage container.
	FuelLevel2 *float64 `json:"fuelLevel2,omitempty"`

	// FuelType Type of fuel currently being utilized by the vehicle acc. SPN 5837
	FuelType     *string       `json:"fuelType,omitempty"`
	GnssPosition *GNSSPosition `json:"gnssPosition,omitempty"`

	// HybridBatteryPackRemainingCharge Indicates the hybrid battery pack remaining charge.
	//  0% means no charge remaining,
	//  100% means full charge remaining.
	//  Is used as well for full electrical vehicles
	HybridBatteryPackRemainingCharge *float64 `json:"hybridBatteryPackRemainingCharge,omitempty"`

	// ParkingBrakeSwitch Switch signal which indicates when the parking brake is set. In general the switch actuated by the operator's park brake control, whether a pedal, lever or other control mechanism
	//  true - parking brake set
	//  false - parking brake not set
	ParkingBrakeSwitch *bool `json:"parkingBrakeSwitch,omitempty"`

	// TachographSpeed The Tacho speed
	TachographSpeed *float64 `json:"tachographSpeed,omitempty"`

	// Trailers List of trailers connected to the truck.
	Trailers *[]struct {
		// CustomerTrailerName The customer's name for the trailer
		CustomerTrailerName *string `json:"customerTrailerName,omitempty"`

		// TrailerAxleLoadSum The sum of the static vertical loads of the trailer axles in kilograms. The load is sent in the EBS22 message of ISO 11992-2.
		TrailerAxleLoadSum *int `json:"trailerAxleLoadSum,omitempty"`

		// TrailerAxles A list of trailer axles
		TrailerAxles *[]struct {
			// TrailerAxleLoad The static vertical load of a trailer axle in kilograms. The load is sent in the RGE22 message of ISO11992-2.
			TrailerAxleLoad *float32 `json:"trailerAxleLoad,omitempty"`

			// TrailerAxlePosition Axle position from 1 to 15, 1 being in the front closest to the truck, according to ISO 11992-2.
			TrailerAxlePosition *int `json:"trailerAxlePosition,omitempty"`
		} `json:"trailerAxles,omitempty"`

		// TrailerIdentificationData The identification data sent by the trailer to the truck in the RGE23 message of ISO 11992-2. An alternative source is the DID (Data identifier definition) record VIN, as specified in ISO 11992-4. Even though both ISO 11992-2 and ISO 11992-4 specifies this as a VIN, the actual data sent from a trailer is not always the true VIN of the trailer.
		TrailerIdentificationData *string `json:"trailerIdentificationData,omitempty"`

		// TrailerNo Trailer number from 1 to 5, 1 being closest to the truck, according to ISO 11992-2.
		TrailerNo *int `json:"trailerNo,omitempty"`

		// TrailerType Indicates the type of the trailer. The type is sent in the EBS24 message of  ISO 11992-2.
		TrailerType *string `json:"trailerType,omitempty"`

		// TrailerVin The vehicle identification number of the trailer. See ISO 3779 (17 characters) If the trailerIdentificationData is reporting a true VIN, trailerVin will have the same value. If it is possible to map the trailerIdentificationData to a true VIN using other sources, the value can be provided here.
		TrailerVin *string `json:"trailerVin,omitempty"`
	} `json:"trailers,omitempty"`

	// VehicleAxles A list of vehicle axles
	VehicleAxles *[]struct {
		// VehicleAxleLoad The static vertical load of a vehicle axle in kilograms.
		VehicleAxleLoad *float32 `json:"vehicleAxleLoad,omitempty"`

		// VehicleAxlePosition Axle position from 1 to 15, 1 being in the front of the truck
		VehicleAxlePosition *int `json:"vehicleAxlePosition,omitempty"`
	} `json:"vehicleAxles,omitempty"`

	// WheelBasedSpeed The vehicle wheelbased speed
	WheelBasedSpeed *float64 `json:"wheelBasedSpeed,omitempty"`
}
