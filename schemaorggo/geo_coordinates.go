package schemaorggo

import "encoding/json"

// GeoCoordinates see : https://schema.org/GeoCoordinates
type GeoCoordinates struct {
	StructuredValue

	typeContext

	// Address see : https://schema.org/address
	// Physical address of the item.
	// types : PostalAddress Text
	Address []interface{} `json:"address,omitempty"`

	// AddressCountry see : https://schema.org/addressCountry
	// The country. For example, USA. You can also provide the two-letter ISO 3166-1 alpha-2 country code (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_3166-1).
	// types : Country Text
	AddressCountry []interface{} `json:"addressCountry,omitempty"`

	// Elevation see : https://schema.org/elevation
	// The elevation of a location (WGS 84 (see: https://schema.orghttps://en.wikipedia.org/wiki/World_Geodetic_System)).
	// types : Number Text
	Elevation []interface{} `json:"elevation,omitempty"`

	// Latitude see : https://schema.org/latitude
	// The latitude of a location. For example 37.42242 (WGS 84 (see: https://schema.orghttps://en.wikipedia.org/wiki/World_Geodetic_System)).
	// types : Number Text
	Latitude []interface{} `json:"latitude,omitempty"`

	// Longitude see : https://schema.org/longitude
	// The longitude of a location. For example -122.08585 (WGS 84 (see: https://schema.orghttps://en.wikipedia.org/wiki/World_Geodetic_System)).
	// types : Number Text
	Longitude []interface{} `json:"longitude,omitempty"`

	// PostalCode see : https://schema.org/postalCode
	// The postal code. For example, 94043.
	// types : Text
	PostalCode []string `json:"postalCode,omitempty"`
}

func (v GeoCoordinates) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.StructuredValue.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.Address
		if len(v.Address) == 1 {
			value = v.Address[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["address"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.AddressCountry
		if len(v.AddressCountry) == 1 {
			value = v.AddressCountry[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["addressCountry"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Elevation
		if len(v.Elevation) == 1 {
			value = v.Elevation[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["elevation"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Latitude
		if len(v.Latitude) == 1 {
			value = v.Latitude[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["latitude"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Longitude
		if len(v.Longitude) == 1 {
			value = v.Longitude[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["longitude"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PostalCode
		if len(v.PostalCode) == 1 {
			value = v.PostalCode[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["postalCode"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v GeoCoordinates) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "GeoCoordinates"

	return data, nil
}

func (v GeoCoordinates) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
