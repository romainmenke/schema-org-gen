package schemaorggo

import "encoding/json"

// FlightReservation see : https://schema.org/FlightReservation
type FlightReservation struct {
	Reservation

	typeContext

	// BoardingGroup see : https://schema.org/boardingGroup
	// The airline-specific indicator of boarding order / preference.
	BoardingGroup string `json:"boardingGroup,omitempty"`

	// PassengerPriorityStatus see : https://schema.org/passengerPriorityStatus
	// The priority status assigned to a passenger for security or boarding (e.g. FastTrack or Priority).
	PassengerPriorityStatus interface{} `json:"passengerPriorityStatus,omitempty"` // types : QualitativeValue Text

	// PassengerSequenceNumber see : https://schema.org/passengerSequenceNumber
	// The passenger's sequence number as assigned by the airline.
	PassengerSequenceNumber string `json:"passengerSequenceNumber,omitempty"`

	// SecurityScreening see : https://schema.org/securityScreening
	// The type of security screening the passenger is subject to.
	SecurityScreening string `json:"securityScreening,omitempty"`
}

func (v FlightReservation) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "FlightReservation"

	return json.Marshal(v)
}

func (v *FlightReservation) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "FlightReservation"

	return json.Marshal(*v)
}
