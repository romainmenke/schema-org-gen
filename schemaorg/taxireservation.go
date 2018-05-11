package schemaorg

import "encoding/json"

// TaxiReservation see : https://schema.org/TaxiReservation
type TaxiReservation struct {

typeContext

Reservation

// PartySize see : https://schema.org/partySize
// Number of people the reservation should accommodate.
PartySize interface{} `json:"partySize"` // types : Integer QuantitativeValue

// PickupLocation see : https://schema.org/pickupLocation
// Where a taxi will pick up a passenger or a rental car can be picked up.
PickupLocation *Place `json:"pickupLocation"`

// PickupTime see : https://schema.org/pickupTime
// When a taxi will pickup a passenger or a rental car can be picked up.
PickupTime interface{} `json:"pickupTime"`

}

func (v *TaxiReservation) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "TaxiReservation"

	return json.Marshal(v)
}
