package schemaorggo

import "encoding/json"

// RentalCarReservation see : https://schema.org/RentalCarReservation
type RentalCarReservation struct {
	Reservation

	typeContext

	// DropoffLocation see : https://schema.org/dropoffLocation
	// Where a rental car can be dropped off.
	DropoffLocation *Place `json:"dropoffLocation,omitempty"`

	// DropoffTime see : https://schema.org/dropoffTime
	// When a rental car can be dropped off.
	DropoffTime DateTime `json:"dropoffTime,omitempty"`

	// PickupLocation see : https://schema.org/pickupLocation
	// Where a taxi will pick up a passenger or a rental car can be picked up.
	PickupLocation *Place `json:"pickupLocation,omitempty"`

	// PickupTime see : https://schema.org/pickupTime
	// When a taxi will pickup a passenger or a rental car can be picked up.
	PickupTime DateTime `json:"pickupTime,omitempty"`
}

func (v RentalCarReservation) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "RentalCarReservation"

	return json.Marshal(v)
}

func (v *RentalCarReservation) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "RentalCarReservation"

	return json.Marshal(*v)
}
