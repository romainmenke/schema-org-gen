package schemaorggo

import "encoding/json"

// ReservationPackage see : https://schema.org/ReservationPackage
type ReservationPackage struct {
	Reservation

	typeContext

	// SubReservation see : https://schema.org/subReservation
	// The individual reservations included in the package. Typically a repeated property.
	// types : Reservation
	SubReservation []*Reservation `json:"subReservation,omitempty"`
}

func (v ReservationPackage) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Reservation.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.SubReservation
		if len(v.SubReservation) == 1 {
			value = v.SubReservation[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["subReservation"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v ReservationPackage) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "ReservationPackage"

	return data, nil
}

func (v ReservationPackage) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
