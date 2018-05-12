package schemaorggo

import "encoding/json"

// TaxiReservation see : https://schema.org/TaxiReservation
type TaxiReservation struct {
	Reservation

	typeContext

	// PartySize see : https://schema.org/partySize
	// Number of people the reservation should accommodate.
	// types : Integer QuantitativeValue
	PartySize interface{} `json:"partySize,omitempty"`

	// PickupLocation see : https://schema.org/pickupLocation
	// Where a taxi will pick up a passenger or a rental car can be picked up.
	// types : Place
	PickupLocation *Place `json:"pickupLocation,omitempty"`

	// PickupTime see : https://schema.org/pickupTime
	// When a taxi will pickup a passenger or a rental car can be picked up.
	// types : DateTime
	PickupTime DateTime `json:"pickupTime,omitempty"`
}

func (v TaxiReservation) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "TaxiReservation"

	return json.Marshal(v)
}

func (v *TaxiReservation) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "TaxiReservation"

	return json.Marshal(*v)
}
