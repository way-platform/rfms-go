package rfms

// Known fuel type values (based on J1939 SPN5837).
const (
	// 0x00 Not available (NONE).
	FuelTypeNotAvailable = "00"
	// 0x01 Gasoline/petrol (GAS).
	FuelTypeGasoline = "01"
	// 0x02 Methanol (METH).
	FuelTypeMethanol = "02"
	// 0x03 Ethanol (ETH).
	FuelTypeEthanol = "03"
	// 0x04 Diesel (DSL).
	FuelTypeDiesel = "04"
	// 0x05 Liquefied Petroleum Gas (LPG).
	FuelTypeLiquefiedPetroleumGas = "05"
	// 0x06 Compressed Natural Gas (CNG).
	FuelTypeCompressedNaturalGas = "06"
	// 0x07 Propane (PROP).
	FuelTypePropane = "07"
	// 0x08 Battery/electric (ELEC).
	FuelTypeBatteryElectric = "08"
	// 0x1D Fuel Cell Utilizing Hydrogen.
	FuelTypeHydrogenFuelCell = "1D"
	// 0x1E Hydrogen Internal Combustion Engine.
	FuelTypeHydrogenInternalCombustionEngine = "1E"
	// 0x1F Kerosene.
	FuelTypeKerosene = "1F"
	// 0x20 Heavy Fuel Oil.
	FuelTypeKerosene = "20"
)
