package schemaorggo

import "encoding/json"

// SingleFamilyResidence see : https://schema.org/SingleFamilyResidence
type SingleFamilyResidence struct {
	House

	typeContext

	// NumberOfRooms see : https://schema.org/numberOfRooms
	// The number of rooms (excluding bathrooms and closets) of the acccommodation or lodging business.
	// Typical unit code(s): ROM for room or C62 for no unit. The type of room can be put in the unitText property of the QuantitativeValue.
	// types : Number QuantitativeValue
	NumberOfRooms []interface{} `json:"numberOfRooms,omitempty"`

	// Occupancy see : https://schema.org/occupancy
	// The allowed total occupancy for the accommodation in persons (including infants etc). For individual accommodations, this is not necessarily the legal maximum but defines the permitted usage as per the contractual agreement (e.g. a double room used by a single person).
	// Typical unit code(s): C62 for person
	// types : QuantitativeValue
	Occupancy []*QuantitativeValue `json:"occupancy,omitempty"`
}

func (v SingleFamilyResidence) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.House.IntoMap(intop)

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

	{
		var value interface{} = v.Occupancy
		if len(v.Occupancy) == 1 {
			value = v.Occupancy[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["occupancy"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v SingleFamilyResidence) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "SingleFamilyResidence"

	return data, nil
}

func (v SingleFamilyResidence) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
