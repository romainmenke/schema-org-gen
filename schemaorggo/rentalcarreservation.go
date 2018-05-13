package schemaorggo

import "encoding/json"

// RentalCarReservation see : https://schema.org/RentalCarReservation
type RentalCarReservation struct {
	Reservation

	typeContext

	// DropoffLocation see : https://schema.org/dropoffLocation
	// Where a rental car can be dropped off.
	// types : Place
	DropoffLocation []*Place `json:"dropoffLocation,omitempty"`

	// DropoffTime see : https://schema.org/dropoffTime
	// When a rental car can be dropped off.
	// types : DateTime
	DropoffTime []DateTime `json:"dropoffTime,omitempty"`

	// PickupLocation see : https://schema.org/pickupLocation
	// Where a taxi will pick up a passenger or a rental car can be picked up.
	// types : Place
	PickupLocation []*Place `json:"pickupLocation,omitempty"`

	// PickupTime see : https://schema.org/pickupTime
	// When a taxi will pickup a passenger or a rental car can be picked up.
	// types : DateTime
	PickupTime []DateTime `json:"pickupTime,omitempty"`
}

func (v RentalCarReservation) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Reservation.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.DropoffLocation
		if len(v.DropoffLocation) == 1 {
			value = v.DropoffLocation[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["dropoffLocation"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DropoffTime
		if len(v.DropoffTime) == 1 {
			value = v.DropoffTime[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["dropoffTime"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PickupLocation
		if len(v.PickupLocation) == 1 {
			value = v.PickupLocation[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["pickupLocation"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PickupTime
		if len(v.PickupTime) == 1 {
			value = v.PickupTime[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["pickupTime"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v RentalCarReservation) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "RentalCarReservation"

	return data, nil
}

func (v RentalCarReservation) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
