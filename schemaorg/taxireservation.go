package schemaorg

import "encoding/json"

// TaxiReservation see : https://schema.org/TaxiReservation
type TaxiReservation struct {

typeContext

Reservation

// PartySize see : /partySize
// Number of people the reservation should accommodate.
PartySize interface{} `json:"partySize"` // types : Integer QuantitativeValue

// PickupLocation see : /pickupLocation
// Where a taxi will pick up a passenger or a rental car can be picked up.
PickupLocation *Place `json:"pickupLocation"`

// PickupTime see : /pickupTime
// When a taxi will pickup a passenger or a rental car can be picked up.
PickupTime interface{} `json:"pickupTime"`

}

func (v *TaxiReservation) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "TaxiReservation"

	return json.Marshal(v)
}
