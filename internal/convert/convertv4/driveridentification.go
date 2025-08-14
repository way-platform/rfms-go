package convertv4

import (
	rfmsv4oapi "github.com/way-platform/rfms-go/internal/openapi/rfmsv4oapi"
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/connect/rfms/v5"
)

// driverIdentification converts an rFMS v4 driver ID to proto.
func driverIdentification(input *rfmsv4oapi.DriverIDObject) *rfmsv5.DriverIdentification {
	var output rfmsv5.DriverIdentification
	if input.TachoDriverIdentification != nil {
		var tacho rfmsv5.DriverIdentification_Tacho
		if input.TachoDriverIdentification.DriverIdentification != nil {
			tacho.SetDriverId(*input.TachoDriverIdentification.DriverIdentification)
		}
		if input.TachoDriverIdentification.CardIssuingMemberState != nil {
			tacho.SetCardIssuingMemberState(*input.TachoDriverIdentification.CardIssuingMemberState)
		}
		if input.TachoDriverIdentification.CardRenewalIndex != nil {
			tacho.SetCardRenewalIndex(*input.TachoDriverIdentification.CardRenewalIndex)
		}
		if input.TachoDriverIdentification.CardReplacementIndex != nil {
			tacho.SetCardReplacementIndex(*input.TachoDriverIdentification.CardReplacementIndex)
		}
		// TODO: Convert DriverAuthenticationEquipment enum if needed
		output.SetTacho(&tacho)
	}
	if input.OEMDriverIdentification != nil {
		var oem rfmsv5.DriverIdentification_Oem
		if input.OEMDriverIdentification.OEMDriverIdentification != nil {
			oem.SetDriverId(*input.OEMDriverIdentification.OEMDriverIdentification)
		}
		if input.OEMDriverIdentification.IDType != nil {
			oem.SetIdType(*input.OEMDriverIdentification.IDType)
		}
		output.SetOem(&oem)
	}
	return &output
}
