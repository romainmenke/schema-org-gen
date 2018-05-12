package schemaorggo

import "encoding/json"

// ReservationPackage see : https://schema.org/ReservationPackage
type ReservationPackage struct {
	Reservation

	typeContext

	// SubReservation see : https://schema.org/subReservation
	// The individual reservations included in the package. Typically a repeated property.
	SubReservation *Reservation `json:"subReservation"`
}

func (v *ReservationPackage) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ReservationPackage"

	return json.Marshal(v)
}
