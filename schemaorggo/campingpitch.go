package schemaorggo

import "encoding/json"

// CampingPitch see : https://schema.org/CampingPitch
type CampingPitch struct {
	Accommodation

	typeContext

	// AmenityFeature see : https://schema.org/amenityFeature
	// An amenity feature (e.g. a characteristic or service) of the Accommodation. This generic property does not make a statement about whether the feature is included in an offer for the main accommodation or available at extra costs.
	AmenityFeature *LocationFeatureSpecification `json:"amenityFeature,omitempty"` // types : LocationFeatureSpecification

	// FloorSize see : https://schema.org/floorSize
	// The size of the accommodation, e.g. in square meter or squarefoot.
	// Typical unit code(s): MTK for square meter, FTK for square foot, or YDK for square yard
	FloorSize *QuantitativeValue `json:"floorSize,omitempty"` // types : QuantitativeValue

	// NumberOfRooms see : https://schema.org/numberOfRooms
	// The number of rooms (excluding bathrooms and closets) of the acccommodation or lodging business.
	// Typical unit code(s): ROM for room or C62 for no unit. The type of room can be put in the unitText property of the QuantitativeValue.
	NumberOfRooms interface{} `json:"numberOfRooms,omitempty"` // types : Number QuantitativeValue

	// PermittedUsage see : https://schema.org/permittedUsage
	// Indications regarding the permitted usage of the accommodation.
	PermittedUsage string `json:"permittedUsage,omitempty"` // types : Text

	// PetsAllowed see : https://schema.org/petsAllowed
	// Indicates whether pets are allowed to enter the accommodation or lodging business. More detailed information can be put in a text value.
	PetsAllowed interface{} `json:"petsAllowed,omitempty"` // types : Boolean Text

}

func (v CampingPitch) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "CampingPitch"

	return json.Marshal(v)
}

func (v *CampingPitch) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "CampingPitch"

	return json.Marshal(*v)
}
