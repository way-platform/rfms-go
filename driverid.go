package rfms

// DriverID contains driver identification.
type DriverID struct {
	// OEM contains OEM specific driver identification.
	OEM *OEMDriverIdentification `json:"oemDriverIdentification,omitempty"`
	// Tacho contains tachograph driver identification.
	Tacho *TachoDriverIdentification `json:"tachoDriverIdentification,omitempty"`
}

// OEMDriverIdentification contains OEM-sourced driver identification.
type OEMDriverIdentification struct {
	// IdType Contains an optional id type (e.g. pin, USB, encrypted EU id...)
	IdType *string `json:"idType,omitempty"`

	// OemDriverIdentification An OEM specific driver id.
	OemDriverIdentification *string `json:"oemDriverIdentification,omitempty"`
}

// TachoDriverIdentification contains tachograph-sourced driver identification.
type TachoDriverIdentification struct {
	// DriverIdentification The unique identification of a driver in a Member State. This fields is formatted according the definition for driverIdentification in COMMISSION REGULATION (EC) No 1360/2002 Annex 1b
	DriverIdentification string `json:"driverIdentification"`

	// CardIssuingMemberState The country alpha code of the Member State having issued the card. This fields is formatted according the definition for NationAlpha in COMMISSION REGULATION (EC) No 1360/2002 Annex 1b
	CardIssuingMemberState string `json:"cardIssuingMemberState"`

	// CardRenewalIndex A card renewal index. This fields is formatted according the definition for CardRenewalIndex (chap 2.25) in: COMMISSION REGULATION (EC) No 1360/2002 Annex 1b
	CardRenewalIndex *string `json:"cardRenewalIndex,omitempty"`

	// CardReplacementIndex A card replacement index. This fields is formatted according the definition for CardReplacementIndex (chap 2.26) in: COMMISSION REGULATION (EC) No 1360/2002 Annex 1b
	CardReplacementIndex *string `json:"cardReplacementIndex,omitempty"`

	// DriverAuthenticationEquipment Code to distinguish different types of equipment for the tachograph application. See description of the field 'DriverAuthenticationEquipment' in COMMISSION REGULATION (EC) No 1360/2002 Annex 1b
	DriverAuthenticationEquipment *string `json:"driverAuthenticationEquipment,omitempty"`
}
