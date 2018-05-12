package schemaorggo

import "encoding/json"

// Suite see : https://schema.org/Suite
type Suite struct {
	Accommodation

	typeContext

	// Bed see : https://schema.org/bed
	// The type of bed or beds included in the accommodation. For the single case of just one bed of a certain type, you use bed directly with a text.
	//       If you want to indicate the quantity of a certain kind of bed, use an instance of BedDetails. For more detailed information, use the amenityFeature property.
	Bed interface{} `json:"bed,omitempty"` // types : BedDetails BedType Text

	// NumberOfRooms see : https://schema.org/numberOfRooms
	// The number of rooms (excluding bathrooms and closets) of the acccommodation or lodging business.
	// Typical unit code(s): ROM for room or C62 for no unit. The type of room can be put in the unitText property of the QuantitativeValue.
	NumberOfRooms interface{} `json:"numberOfRooms,omitempty"` // types : Number QuantitativeValue

	// Occupancy see : https://schema.org/occupancy
	// The allowed total occupancy for the accommodation in persons (including infants etc). For individual accommodations, this is not necessarily the legal maximum but defines the permitted usage as per the contractual agreement (e.g. a double room used by a single person).
	// Typical unit code(s): C62 for person
	Occupancy *QuantitativeValue `json:"occupancy,omitempty"` // types : QuantitativeValue

}

func (v Suite) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Suite"

	return json.Marshal(v)
}

func (v *Suite) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "Suite"

	return json.Marshal(*v)
}
