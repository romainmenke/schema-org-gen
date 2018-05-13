package schemaorggo

import "encoding/json"

// FlightReservation see : https://schema.org/FlightReservation
type FlightReservation struct {
	Reservation

	typeContext

	// BoardingGroup see : https://schema.org/boardingGroup
	// The airline-specific indicator of boarding order / preference.
	// types : Text
	BoardingGroup []string `json:"boardingGroup,omitempty"`

	// PassengerPriorityStatus see : https://schema.org/passengerPriorityStatus
	// The priority status assigned to a passenger for security or boarding (e.g. FastTrack or Priority).
	// types : QualitativeValue Text
	PassengerPriorityStatus []interface{} `json:"passengerPriorityStatus,omitempty"`

	// PassengerSequenceNumber see : https://schema.org/passengerSequenceNumber
	// The passenger&#39;s sequence number as assigned by the airline.
	// types : Text
	PassengerSequenceNumber []string `json:"passengerSequenceNumber,omitempty"`

	// SecurityScreening see : https://schema.org/securityScreening
	// The type of security screening the passenger is subject to.
	// types : Text
	SecurityScreening []string `json:"securityScreening,omitempty"`
}

func (v FlightReservation) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Reservation.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.BoardingGroup
		if len(v.BoardingGroup) == 1 {
			value = v.BoardingGroup[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["boardingGroup"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PassengerPriorityStatus
		if len(v.PassengerPriorityStatus) == 1 {
			value = v.PassengerPriorityStatus[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["passengerPriorityStatus"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PassengerSequenceNumber
		if len(v.PassengerSequenceNumber) == 1 {
			value = v.PassengerSequenceNumber[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["passengerSequenceNumber"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SecurityScreening
		if len(v.SecurityScreening) == 1 {
			value = v.SecurityScreening[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["securityScreening"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v FlightReservation) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "FlightReservation"

	return data, nil
}

func (v FlightReservation) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
