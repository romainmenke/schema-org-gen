package schemaorggo

import "encoding/json"

// HotelRoom see : https://schema.org/HotelRoom
type HotelRoom struct {
	Room

	typeContext

	// Bed see : https://schema.org/bed
	// The type of bed or beds included in the accommodation. For the single case of just one bed of a certain type, you use bed directly with a text.
	//       If you want to indicate the quantity of a certain kind of bed, use an instance of BedDetails. For more detailed information, use the amenityFeature property.
	Bed interface{} `json:"bed"` // types : BedDetails BedType Text

	// Occupancy see : https://schema.org/occupancy
	// The allowed total occupancy for the accommodation in persons (including infants etc). For individual accommodations, this is not necessarily the legal maximum but defines the permitted usage as per the contractual agreement (e.g. a double room used by a single person).
	// Typical unit code(s): C62 for person
	Occupancy *QuantitativeValue `json:"occupancy"`
}

func (v *HotelRoom) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "HotelRoom"

	return json.Marshal(v)
}
