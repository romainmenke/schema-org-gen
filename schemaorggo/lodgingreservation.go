package schemaorggo

import "encoding/json"

// LodgingReservation see : https://schema.org/LodgingReservation
type LodgingReservation struct {
	Reservation

	typeContext

	// CheckinTime see : https://schema.org/checkinTime
	// The earliest someone may check into a lodging establishment.
	// types : DateTime
	CheckinTime DateTime `json:"checkinTime,omitempty"`

	// CheckoutTime see : https://schema.org/checkoutTime
	// The latest someone may check out of a lodging establishment.
	// types : DateTime
	CheckoutTime DateTime `json:"checkoutTime,omitempty"`

	// LodgingUnitDescription see : https://schema.org/lodgingUnitDescription
	// A full description of the lodging unit.
	// types : Text
	LodgingUnitDescription string `json:"lodgingUnitDescription,omitempty"`

	// LodgingUnitType see : https://schema.org/lodgingUnitType
	// Textual description of the unit type (including suite vs. room, size of bed, etc.).
	// types : QualitativeValue Text
	LodgingUnitType interface{} `json:"lodgingUnitType,omitempty"`

	// NumAdults see : https://schema.org/numAdults
	// The number of adults staying in the unit.
	// types : Integer QuantitativeValue
	NumAdults interface{} `json:"numAdults,omitempty"`

	// NumChildren see : https://schema.org/numChildren
	// The number of children staying in the unit.
	// types : Integer QuantitativeValue
	NumChildren interface{} `json:"numChildren,omitempty"`
}

func (v LodgingReservation) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "LodgingReservation"

	return json.Marshal(v)
}

func (v *LodgingReservation) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "LodgingReservation"

	return json.Marshal(*v)
}
