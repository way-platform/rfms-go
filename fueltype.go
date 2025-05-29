package rfms

// Known fuel type values (based on J1939 SPN5837).
const (
	FuelTypeNotAvailable = "00"
	FuelTypeGasoline = "01"
	FuelTypeMethanol = "02"
	FuelTypeEthanol = "03"
	FuelTypeDiesel = "04"
	FuelTypeLiquefiedPetroleumGas = "05"
	FuelTypeCompressedNaturalGas = "06"
	FuelTypePropane = "07"
	FuelTypeBatteryElectric = "08"
	FuelTypeHydrogenFuelCell = "1D"
)
