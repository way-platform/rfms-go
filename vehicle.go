package rfms

// Vehicle is an rFMS vehicle.
type Vehicle struct {
	// AuthorizedPaths the client is authorized to call.
	AuthorizedPaths []string `json:"authorizedPaths,omitempty"`

	// BodyType is the type of body on the chassis. rFMS standard values CITY_BUS, INTERCITY_BUS, COACH. This is used mainly for buses.
	BodyType string `json:"bodyType,omitempty"`

	// Brand The vehicle brand. rFMS standard values VOLVO TRUCKS, SCANIA, DAIMLER, IVECO, DAF, MAN, RENAULT TRUCKS, VDL, VOLVO BUSES, IVECO BUS, IRISBUS
	Brand string `json:"brand,omitempty"`

	// ChassisType The chassis type of the vehicle. OEM specific value. This is used mainly for buses
	ChassisType string `json:"chassisType,omitempty"`

	// CustomerVehicleName The customer's name for the vehicle.
	CustomerVehicleName string `json:"customerVehicleName,omitempty"`

	// DoorConfiguration The door configuration. The door order definition is OEM specific. E.g. [1, 2, 2] means the bus has 3 doors: 1 front door, double doors for door 2 and 3. This is used mainly for buses.
	DoorConfiguration []int `json:"doorConfiguration,omitempty"`

	// EmissionLevel The emission level this vehicle supports. Possible values:
	//  European Union, Heavy-Duty Truck and Bus Engines:
	//  EURO_III, EURO_III_EEV, EURO_IV, EURO_V, EURO_VI
	//  European Union, Nonroad Engines:
	//  EURO_STAGE_III, EURO_STAGE_IV, EURO_STAGE_V
	//  United_States, Heavy-Duty Truck and Bus Engines:
	//  EPA_2004, EPA_2007, EPA_2010, EPA_2015_NOX10, EPA_2015_NOX05, EPA_2015_NOX02
	//  United_States, Nonroad Engines:
	//  EPA_TIER_2, EPA_TIER_3, EPA_TIER_4_2008, EPA_TIER_4_2013
	//  Brazil, Heavy-Duty Truck and Bus Engines:
	//  PROCONVE_P5, PROCONVE_P6, PROCONVE_P7
	//  Brazil, Nonroad Engines:
	//  PROCONVE_MARI
	EmissionLevel string `json:"emissionLevel,omitempty"`

	// GearboxType The type of gearbox the vehicle is equipped with. rFMS standard values MANUAL, AUTOMATIC, SEMI_AUTOMATIC, NO_GEAR (e.g electrical)
	GearboxType string `json:"gearboxType,omitempty"`

	// HasRampOrLift If the vehicle is equipped with a ramp or not. This is used mainly for buses.
	HasRampOrLift bool `json:"hasRampOrLift,omitempty"`

	// Model Indicates the model of the vehicle. OEM specific value.
	Model string `json:"model,omitempty"`

	// NoOfAxles Number of axles on the vehicle. This is used mainly for buses
	NoOfAxles int `json:"noOfAxles,omitempty"`

	// PossibleFuelType The possible fuel types supported by this vehicle, formatted as the HEX id number according to SPN 5837. This does NOT indicate which fuel type that is presently being used.
	PossibleFuelType []string `json:"possibleFuelType,omitempty"`

	// ProductionDate Indicates when the vehicle was produced.
	ProductionDate Date `json:"productionDate,omitempty"`

	// TachographType The type of tachograph in the vehicle. rFMS standard values MTCO, DTCO, TSU, DTCO_G1, DTCO_G2, NONE
	//  DTCO - Digital tachograph, unknown generation
	//  DTCO_G1 - Digital tachograph generation 1
	//  DTCO_G2 - Digital tachograph generation 2
	//  NONE - No tachograph in the vehicle
	//  MTCO - Modular tachograph
	//  TSU - Tachograph simulator
	TachographType string `json:"tachographType,omitempty"`

	// TellTaleCode This parameter indicates how the tell tales shall be interpreted, the code is unique for each OEM. One OEM can have different interpretations  depending on vehicle type.
	TellTaleCode string `json:"tellTaleCode,omitempty"`

	// TotalBatteryPackCapacityWh Total battery pack capacity in watt hours.
	TotalBatteryPackCapacityWh int64 `json:"totalBatteryPackCapacity,omitempty"`

	// TotalFuelTankCapacityGaseousKg Total gas tank capacity for all tanks in kilograms.
	TotalFuelTankCapacityGaseousKg int64 `json:"totalFuelTankCapacityGaseous,omitempty"`

	// TotalFuelTankVolumeMl Total fuel tank volume for all tanks in milliltres.
	TotalFuelTankVolumeMl int64 `json:"totalFuelTankVolume,omitempty"`

	// Type Indicates the type of vehicle. rFMS standard values TRUCK, BUS, VAN
	Type string `json:"type,omitempty"`

	// Vin vehicle identification number. See ISO 3779 (17 characters)
	Vin string `json:"vin"`
}

// A date in the format YYYY-MM-DD.
type Date struct {
	// Day of the month where first day of the month is 1
	Day int32 `json:"day,omitempty"`
	// Month of the year, where January is value 1
	Month int32 `json:"month,omitempty"`
	// Year of the date.
	Year int32 `json:"year,omitempty"`
}

// IsZero returns true if the date is not set.
func (d Date) IsZero() bool {
	return d.Day == 0 && d.Month == 0 && d.Year == 0
}
