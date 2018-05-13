package schemaorggo

import "encoding/json"

// Suite see : https://schema.org/Suite
type Suite struct {
	Accommodation

	typeContext

	// Bed see : https://schema.org/bed
	// The type of bed or beds included in the accommodation. For the single case of just one bed of a certain type, you use bed directly with a text.
	//       If you want to indicate the quantity of a certain kind of bed, use an instance of BedDetails. For more detailed information, use the amenityFeature property.
	// types : BedDetails BedType Text
	Bed []interface{} `json:"bed,omitempty"`

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

func (v Suite) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Accommodation.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.Bed
		if len(v.Bed) == 1 {
			value = v.Bed[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["bed"] = json.RawMessage(b)
		}
	}

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

func (v Suite) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Suite"

	return data, nil
}

func (v Suite) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
