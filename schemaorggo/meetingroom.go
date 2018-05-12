package schemaorggo

import "encoding/json"

// MeetingRoom see : https://schema.org/MeetingRoom
type MeetingRoom struct {
	Room

	typeContext

	// AmenityFeature see : https://schema.org/amenityFeature
	// An amenity feature (e.g. a characteristic or service) of the Accommodation. This generic property does not make a statement about whether the feature is included in an offer for the main accommodation or available at extra costs.
	// types : LocationFeatureSpecification
	AmenityFeature *LocationFeatureSpecification `json:"amenityFeature,omitempty"`

	// FloorSize see : https://schema.org/floorSize
	// The size of the accommodation, e.g. in square meter or squarefoot.
	// Typical unit code(s): MTK for square meter, FTK for square foot, or YDK for square yard
	// types : QuantitativeValue
	FloorSize *QuantitativeValue `json:"floorSize,omitempty"`

	// NumberOfRooms see : https://schema.org/numberOfRooms
	// The number of rooms (excluding bathrooms and closets) of the acccommodation or lodging business.
	// Typical unit code(s): ROM for room or C62 for no unit. The type of room can be put in the unitText property of the QuantitativeValue.
	// types : Number QuantitativeValue
	NumberOfRooms interface{} `json:"numberOfRooms,omitempty"`

	// PermittedUsage see : https://schema.org/permittedUsage
	// Indications regarding the permitted usage of the accommodation.
	// types : Text
	PermittedUsage string `json:"permittedUsage,omitempty"`

	// PetsAllowed see : https://schema.org/petsAllowed
	// Indicates whether pets are allowed to enter the accommodation or lodging business. More detailed information can be put in a text value.
	// types : Boolean Text
	PetsAllowed interface{} `json:"petsAllowed,omitempty"`
}

func (v MeetingRoom) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MeetingRoom"

	return json.Marshal(v)
}

func (v *MeetingRoom) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "MeetingRoom"

	return json.Marshal(*v)
}
