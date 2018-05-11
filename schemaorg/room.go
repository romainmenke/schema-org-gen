package schemaorg

import "encoding/json"

// Room see : https://schema.org/Room
type Room struct {

typeContext

Accommodation

// AmenityFeature see : /amenityFeature
// An amenity feature (e.g. a characteristic or service) of the Accommodation. This generic property does not make a statement about whether the feature is included in an offer for the main accommodation or available at extra costs.
AmenityFeature *LocationFeatureSpecification `json:"amenityFeature"`

// FloorSize see : /floorSize
// The size of the accommodation, e.g. in square meter or squarefoot.
// Typical unit code(s): MTK for square meter, FTK for square foot, or YDK for square yard
FloorSize *QuantitativeValue `json:"floorSize"`

// NumberOfRooms see : /numberOfRooms
// The number of rooms (excluding bathrooms and closets) of the acccommodation or lodging business.
// Typical unit code(s): ROM for room or C62 for no unit. The type of room can be put in the unitText property of the QuantitativeValue.
NumberOfRooms interface{} `json:"numberOfRooms"` // types : Number QuantitativeValue

// PermittedUsage see : /permittedUsage
// Indications regarding the permitted usage of the accommodation.
PermittedUsage string `json:"permittedUsage"`

// PetsAllowed see : /petsAllowed
// Indicates whether pets are allowed to enter the accommodation or lodging business. More detailed information can be put in a text value.
PetsAllowed interface{} `json:"petsAllowed"` // types : Boolean Text

}

func (v *Room) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Room"

	return json.Marshal(v)
}
