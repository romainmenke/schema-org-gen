package schemaorggo

import "encoding/json"

// ReservationPackage see : https://schema.org/ReservationPackage
type ReservationPackage struct {
	Reservation

	typeContext

	// SubReservation see : https://schema.org/subReservation
	// The individual reservations included in the package. Typically a repeated property.
	// types : Reservation
	SubReservation *Reservation `json:"subReservation,omitempty"`
}

func (v ReservationPackage) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ReservationPackage"

	return json.Marshal(v)
}

func (v *ReservationPackage) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "ReservationPackage"

	return json.Marshal(*v)
}
