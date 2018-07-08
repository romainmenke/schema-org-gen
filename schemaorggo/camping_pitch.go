package schemaorggo

import "encoding/json"

// CampingPitch see : https://schema.org/CampingPitch
type CampingPitch struct {
	Accommodation

	typeContext

	// AmenityFeature see : https://schema.org/amenityFeature
	// An amenity feature (e.g. a characteristic or service) of the Accommodation. This generic property does not make a statement about whether the feature is included in an offer for the main accommodation or available at extra costs.
	// types : LocationFeatureSpecification
	AmenityFeature []*LocationFeatureSpecification `json:"amenityFeature,omitempty"`

	// FloorSize see : https://schema.org/floorSize
	// The size of the accommodation, e.g. in square meter or squarefoot.
	// Typical unit code(s): MTK for square meter, FTK for square foot, or YDK for square yard
	// types : QuantitativeValue
	FloorSize []*QuantitativeValue `json:"floorSize,omitempty"`

	// NumberOfRooms see : https://schema.org/numberOfRooms
	// The number of rooms (excluding bathrooms and closets) of the accommodation or lodging business.
	// Typical unit code(s): ROM for room or C62 for no unit. The type of room can be put in the unitText property of the QuantitativeValue.
	// types : Number QuantitativeValue
	NumberOfRooms []interface{} `json:"numberOfRooms,omitempty"`

	// PermittedUsage see : https://schema.org/permittedUsage
	// Indications regarding the permitted usage of the accommodation.
	// types : Text
	PermittedUsage []string `json:"permittedUsage,omitempty"`

	// PetsAllowed see : https://schema.org/petsAllowed
	// Indicates whether pets are allowed to enter the accommodation or lodging business. More detailed information can be put in a text value.
	// types : Boolean Text
	PetsAllowed []interface{} `json:"petsAllowed,omitempty"`
}

func (v CampingPitch) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Accommodation.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.AmenityFeature
		if len(v.AmenityFeature) == 1 {
			value = v.AmenityFeature[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["amenityFeature"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.FloorSize
		if len(v.FloorSize) == 1 {
			value = v.FloorSize[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["floorSize"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.NumberOfRooms
		if len(v.NumberOfRooms) == 1 {
			value = v.NumberOfRooms[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["numberOfRooms"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PermittedUsage
		if len(v.PermittedUsage) == 1 {
			value = v.PermittedUsage[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["permittedUsage"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PetsAllowed
		if len(v.PetsAllowed) == 1 {
			value = v.PetsAllowed[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["petsAllowed"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v CampingPitch) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "CampingPitch"

	return data, nil
}

func (v CampingPitch) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
