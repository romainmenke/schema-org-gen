package schemaorg

import "encoding/json"

// Apartment see : https://schema.org/Apartment
type Apartment struct {

typeContext

Accommodation

// NumberOfRooms see : /numberOfRooms
// The number of rooms (excluding bathrooms and closets) of the acccommodation or lodging business.
// Typical unit code(s): ROM for room or C62 for no unit. The type of room can be put in the unitText property of the QuantitativeValue.
NumberOfRooms interface{} `json:"numberOfRooms"` // types : Number QuantitativeValue

// Occupancy see : /occupancy
// The allowed total occupancy for the accommodation in persons (including infants etc). For individual accommodations, this is not necessarily the legal maximum but defines the permitted usage as per the contractual agreement (e.g. a double room used by a single person).
// Typical unit code(s): C62 for person
Occupancy *QuantitativeValue `json:"occupancy"`

}

func (v *Apartment) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Apartment"

	return json.Marshal(v)
}
