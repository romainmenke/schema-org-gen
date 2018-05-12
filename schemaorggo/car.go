package schemaorggo

import "encoding/json"

// Car see : https://schema.org/Car
type Car struct {
	Vehicle

	typeContext

	// AcrissCode see : https://schema.org/acrissCode
	// The ACRISS Car Classification Code is a code used by many car rental companies, for classifying vehicles. ACRISS stands for Association of Car Rental Industry Systems and Standards.
	AcrissCode string `json:"acrissCode,omitempty"` // types : Text

	// RoofLoad see : https://schema.org/roofLoad
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
	RoofLoad *QuantitativeValue `json:"roofLoad,omitempty"` // types : QuantitativeValue

}

func (v Car) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Car"

	return json.Marshal(v)
}

func (v *Car) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "Car"

	return json.Marshal(*v)
}
