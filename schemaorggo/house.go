package schemaorggo

import "encoding/json"

// House see : https://schema.org/House
type House struct {
	Accommodation

	typeContext

	// NumberOfRooms see : https://schema.org/numberOfRooms
	// The number of rooms (excluding bathrooms and closets) of the acccommodation or lodging business.
	// Typical unit code(s): ROM for room or C62 for no unit. The type of room can be put in the unitText property of the QuantitativeValue.
	NumberOfRooms interface{} `json:"numberOfRooms"` // types : Number QuantitativeValue

}

func (v *House) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "House"

	return json.Marshal(v)
}
