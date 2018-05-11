package schemaorg

import "encoding/json"

// FlightReservation see : https://schema.org/FlightReservation
type FlightReservation struct {

typeContext

Reservation

// BoardingGroup see : /boardingGroup
// The airline-specific indicator of boarding order / preference.
BoardingGroup string `json:"boardingGroup"`

// PassengerPriorityStatus see : /passengerPriorityStatus
// The priority status assigned to a passenger for security or boarding (e.g. FastTrack or Priority).
PassengerPriorityStatus interface{} `json:"passengerPriorityStatus"` // types : QualitativeValue Text

// PassengerSequenceNumber see : /passengerSequenceNumber
// The passenger's sequence number as assigned by the airline.
PassengerSequenceNumber string `json:"passengerSequenceNumber"`

// SecurityScreening see : /securityScreening
// The type of security screening the passenger is subject to.
SecurityScreening string `json:"securityScreening"`

}

func (v *FlightReservation) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "FlightReservation"

	return json.Marshal(v)
}
