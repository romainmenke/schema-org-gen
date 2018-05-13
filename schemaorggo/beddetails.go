package schemaorggo

import "encoding/json"

// BedDetails see : https://schema.org/BedDetails
type BedDetails struct {
	Intangible

	typeContext

	// NumberOfBeds see : https://schema.org/numberOfBeds
	// The quantity of the given bed type available in the HotelRoom, Suite, House, or Apartment.
	// types : Number
	NumberOfBeds []float64 `json:"numberOfBeds,omitempty"`

	// TypeOfBed see : https://schema.org/typeOfBed
	// The type of bed to which the BedDetail refers, i.e. the type of bed available in the quantity indicated by quantity.
	// types : BedType Text
	TypeOfBed []interface{} `json:"typeOfBed,omitempty"`
}

func (v BedDetails) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Intangible.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.NumberOfBeds
		if len(v.NumberOfBeds) == 1 {
			value = v.NumberOfBeds[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["numberOfBeds"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.TypeOfBed
		if len(v.TypeOfBed) == 1 {
			value = v.TypeOfBed[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["typeOfBed"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v BedDetails) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "BedDetails"

	return data, nil
}

func (v BedDetails) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
