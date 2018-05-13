package schemaorggo

import "encoding/json"

// Airport see : https://schema.org/Airport
type Airport struct {
	CivicStructure

	typeContext

	// IataCode see : https://schema.org/iataCode
	// IATA identifier for an airline or airport.
	// types : Text
	IataCode []string `json:"iataCode,omitempty"`

	// IcaoCode see : https://schema.org/icaoCode
	// ICAO identifier for an airport.
	// types : Text
	IcaoCode []string `json:"icaoCode,omitempty"`
}

func (v Airport) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.CivicStructure.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.IataCode
		if len(v.IataCode) == 1 {
			value = v.IataCode[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["iataCode"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.IcaoCode
		if len(v.IcaoCode) == 1 {
			value = v.IcaoCode[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["icaoCode"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v Airport) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Airport"

	return data, nil
}

func (v Airport) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
