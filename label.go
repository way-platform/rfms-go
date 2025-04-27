package rfms

// Label defines model for LabelObject.
type Label struct {
	Kilograms   *int64  `json:"kilograms,omitempty"`
	Label       *string `json:"label,omitempty"`
	Meters      *int64  `json:"meters,omitempty"`
	MilliLitres *int64  `json:"milliLitres,omitempty"`
	Seconds     *int64  `json:"seconds,omitempty"`
	Watthours   *int64  `json:"watthours,omitempty"`
}
