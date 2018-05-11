package schemaorg

import "encoding/json"

// RentalCarReservation see : https://schema.org/RentalCarReservation
type RentalCarReservation struct {

typeContext

Reservation

// DropoffLocation see : /dropoffLocation
// Where a rental car can be dropped off.
DropoffLocation *Place `json:"dropoffLocation"`

// DropoffTime see : /dropoffTime
// When a rental car can be dropped off.
DropoffTime interface{} `json:"dropoffTime"`

// PickupLocation see : /pickupLocation
// Where a taxi will pick up a passenger or a rental car can be picked up.
PickupLocation *Place `json:"pickupLocation"`

// PickupTime see : /pickupTime
// When a taxi will pickup a passenger or a rental car can be picked up.
PickupTime interface{} `json:"pickupTime"`

}

func (v *RentalCarReservation) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "RentalCarReservation"

	return json.Marshal(v)
}
