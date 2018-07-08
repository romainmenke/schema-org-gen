package schemaorggo

import "encoding/json"

// BusTrip see : https://schema.org/BusTrip
type BusTrip struct {
	Trip

	typeContext

	// ArrivalBusStop see : https://schema.org/arrivalBusStop
	// The stop or station from which the bus arrives.
	// types : BusStation BusStop
	ArrivalBusStop []interface{} `json:"arrivalBusStop,omitempty"`

	// BusName see : https://schema.org/busName
	// The name of the bus (e.g. Bolt Express).
	// types : Text
	BusName []string `json:"busName,omitempty"`

	// BusNumber see : https://schema.org/busNumber
	// The unique identifier for the bus.
	// types : Text
	BusNumber []string `json:"busNumber,omitempty"`

	// DepartureBusStop see : https://schema.org/departureBusStop
	// The stop or station from which the bus departs.
	// types : BusStation BusStop
	DepartureBusStop []interface{} `json:"departureBusStop,omitempty"`
}

func (v BusTrip) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Trip.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.ArrivalBusStop
		if len(v.ArrivalBusStop) == 1 {
			value = v.ArrivalBusStop[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["arrivalBusStop"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.BusName
		if len(v.BusName) == 1 {
			value = v.BusName[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["busName"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.BusNumber
		if len(v.BusNumber) == 1 {
			value = v.BusNumber[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["busNumber"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DepartureBusStop
		if len(v.DepartureBusStop) == 1 {
			value = v.DepartureBusStop[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["departureBusStop"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v BusTrip) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "BusTrip"

	return data, nil
}

func (v BusTrip) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
