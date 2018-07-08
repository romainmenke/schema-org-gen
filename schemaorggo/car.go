package schemaorggo

import "encoding/json"

// Car see : https://schema.org/Car
type Car struct {
	Vehicle

	typeContext

	// AcrissCode see : https://auto.schema.org/acrissCode
	// The ACRISS Car Classification Code is a code used by many car rental companies, for classifying vehicles. ACRISS stands for Association of Car Rental Industry Systems and Standards.
	// types : Text
	AcrissCode []string `json:"acrissCode,omitempty"`

	// RoofLoad see : https://auto.schema.org/roofLoad
	// The permitted total weight of cargo and installations (e.g. a roof rack) on top of the vehicle.
	//
	// Typical unit code(s): KGM for kilogram, LBR for pound
	//
	//
	// Note 1: You can indicate additional information in the name (see: https://schema.org/name) of the QuantitativeValue (see: https://schema.org/QuantitativeValue) node.
	// Note 2: You may also link to a QualitativeValue (see: https://schema.org/QualitativeValue) node that provides additional information using valueReference (see: https://schema.org/valueReference)
	// Note 3: Note that you can use minValue (see: https://schema.org/minValue) and maxValue (see: https://schema.org/maxValue) to indicate ranges.
	//
	//
	// types : QuantitativeValue
	RoofLoad []*QuantitativeValue `json:"roofLoad,omitempty"`
}

func (v Car) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Vehicle.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.AcrissCode
		if len(v.AcrissCode) == 1 {
			value = v.AcrissCode[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["acrissCode"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.RoofLoad
		if len(v.RoofLoad) == 1 {
			value = v.RoofLoad[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["roofLoad"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v Car) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Car"

	return data, nil
}

func (v Car) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
