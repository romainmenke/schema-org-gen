package schemaorggo

import "encoding/json"

// House see : https://schema.org/House
type House struct {
	Accommodation

	typeContext

	// NumberOfRooms see : https://schema.org/numberOfRooms
	// The number of rooms (excluding bathrooms and closets) of the acccommodation or lodging business.
	// Typical unit code(s): ROM for room or C62 for no unit. The type of room can be put in the unitText property of the QuantitativeValue.
	// types : Number QuantitativeValue
	NumberOfRooms []interface{} `json:"numberOfRooms,omitempty"`
}

func (v House) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Accommodation.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.NumberOfRooms
		if len(v.NumberOfRooms) == 1 {
			value = v.NumberOfRooms[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["numberOfRooms"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v House) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "House"

	return data, nil
}

func (v House) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
