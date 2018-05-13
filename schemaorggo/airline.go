package schemaorggo

import "encoding/json"

// Airline see : https://schema.org/Airline
type Airline struct {
	Organization

	typeContext

	// BoardingPolicy see : https://schema.org/boardingPolicy
	// The type of boarding policy used by the airline (e.g. zone-based or group-based).
	// types : BoardingPolicyType
	BoardingPolicy []*BoardingPolicyType `json:"boardingPolicy,omitempty"`

	// IataCode see : https://schema.org/iataCode
	// IATA identifier for an airline or airport.
	// types : Text
	IataCode []string `json:"iataCode,omitempty"`
}

func (v Airline) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Organization.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.BoardingPolicy
		if len(v.BoardingPolicy) == 1 {
			value = v.BoardingPolicy[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["boardingPolicy"] = json.RawMessage(b)
		}
	}

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

	*intop = into

	return nil
}

func (v Airline) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Airline"

	return data, nil
}

func (v Airline) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
