package schemaorg

import "encoding/json"

// LodgingReservation see : https://schema.org/LodgingReservation
type LodgingReservation struct {

typeContext

Reservation

// CheckinTime see : /checkinTime
// The earliest someone may check into a lodging establishment.
CheckinTime interface{} `json:"checkinTime"`

// CheckoutTime see : /checkoutTime
// The latest someone may check out of a lodging establishment.
CheckoutTime interface{} `json:"checkoutTime"`

// LodgingUnitDescription see : /lodgingUnitDescription
// A full description of the lodging unit.
LodgingUnitDescription string `json:"lodgingUnitDescription"`

// LodgingUnitType see : /lodgingUnitType
// Textual description of the unit type (including suite vs. room, size of bed, etc.).
LodgingUnitType interface{} `json:"lodgingUnitType"` // types : QualitativeValue Text

// NumAdults see : /numAdults
// The number of adults staying in the unit.
NumAdults interface{} `json:"numAdults"` // types : Integer QuantitativeValue

// NumChildren see : /numChildren
// The number of children staying in the unit.
NumChildren interface{} `json:"numChildren"` // types : Integer QuantitativeValue

}

func (v *LodgingReservation) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "LodgingReservation"

	return json.Marshal(v)
}
