package schemaorggo

import "encoding/json"

// LodgingReservation see : https://schema.org/LodgingReservation
type LodgingReservation struct {
	Reservation

	typeContext

	// CheckinTime see : https://schema.org/checkinTime
	// The earliest someone may check into a lodging establishment.
	CheckinTime DateTime `json:"checkinTime"`

	// CheckoutTime see : https://schema.org/checkoutTime
	// The latest someone may check out of a lodging establishment.
	CheckoutTime DateTime `json:"checkoutTime"`

	// LodgingUnitDescription see : https://schema.org/lodgingUnitDescription
	// A full description of the lodging unit.
	LodgingUnitDescription string `json:"lodgingUnitDescription"`

	// LodgingUnitType see : https://schema.org/lodgingUnitType
	// Textual description of the unit type (including suite vs. room, size of bed, etc.).
	LodgingUnitType interface{} `json:"lodgingUnitType"` // types : QualitativeValue Text

	// NumAdults see : https://schema.org/numAdults
	// The number of adults staying in the unit.
	NumAdults interface{} `json:"numAdults"` // types : Integer QuantitativeValue

	// NumChildren see : https://schema.org/numChildren
	// The number of children staying in the unit.
	NumChildren interface{} `json:"numChildren"` // types : Integer QuantitativeValue

}

func (v *LodgingReservation) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "LodgingReservation"

	return json.Marshal(v)
}
