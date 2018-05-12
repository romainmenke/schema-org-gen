package schemaorg

import "encoding/json"

// BedDetails see : https://schema.org/BedDetails
type BedDetails struct {

Intangible

typeContext

// NumberOfBeds see : https://schema.org/numberOfBeds
// The quantity of the given bed type available in the HotelRoom, Suite, House, or Apartment.
NumberOfBeds float64 `json:"numberOfBeds"`

// TypeOfBed see : https://schema.org/typeOfBed
// The type of bed to which the BedDetail refers, i.e. the type of bed available in the quantity indicated by quantity.
TypeOfBed interface{} `json:"typeOfBed"` // types : BedType Text

}

func (v *BedDetails) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "BedDetails"

	return json.Marshal(v)
}
