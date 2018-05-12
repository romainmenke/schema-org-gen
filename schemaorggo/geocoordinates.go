package schemaorggo

import "encoding/json"

// GeoCoordinates see : https://schema.org/GeoCoordinates
type GeoCoordinates struct {
	StructuredValue

	typeContext

	// Address see : https://schema.org/address
	// Physical address of the item.
	// types : PostalAddress Text
	Address interface{} `json:"address,omitempty"`

	// AddressCountry see : https://schema.org/addressCountry
	// The country. For example, USA. You can also provide the two-letter ISO 3166-1 alpha-2 country code (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_3166-1).
	// types : Country Text
	AddressCountry interface{} `json:"addressCountry,omitempty"`

	// Elevation see : https://schema.org/elevation
	// The elevation of a location (WGS 84 (see: https://schema.orghttps://en.wikipedia.org/wiki/World_Geodetic_System)).
	// types : Number Text
	Elevation interface{} `json:"elevation,omitempty"`

	// Latitude see : https://schema.org/latitude
	// The latitude of a location. For example 37.42242 (WGS 84 (see: https://schema.orghttps://en.wikipedia.org/wiki/World_Geodetic_System)).
	// types : Number Text
	Latitude interface{} `json:"latitude,omitempty"`

	// Longitude see : https://schema.org/longitude
	// The longitude of a location. For example -122.08585 (WGS 84 (see: https://schema.orghttps://en.wikipedia.org/wiki/World_Geodetic_System)).
	// types : Number Text
	Longitude interface{} `json:"longitude,omitempty"`

	// PostalCode see : https://schema.org/postalCode
	// The postal code. For example, 94043.
	// types : Text
	PostalCode string `json:"postalCode,omitempty"`
}

func (v GeoCoordinates) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "GeoCoordinates"

	return json.Marshal(v)
}

func (v *GeoCoordinates) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "GeoCoordinates"

	return json.Marshal(*v)
}
