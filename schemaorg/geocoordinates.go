package schemaorg

import "encoding/json"

// GeoCoordinates see : https://schema.org/GeoCoordinates
type GeoCoordinates struct {

typeContext

StructuredValue

// Address see : /address
// Physical address of the item.
Address interface{} `json:"address"` // types : PostalAddress Text

// AddressCountry see : /addressCountry
// The country. For example, USA. You can also provide the two-letter ISO 3166-1 alpha-2 country code (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_3166-1).
AddressCountry interface{} `json:"addressCountry"` // types : Country Text

// Elevation see : /elevation
// The elevation of a location (WGS 84 (see: https://schema.orghttps://en.wikipedia.org/wiki/World_Geodetic_System)).
Elevation interface{} `json:"elevation"` // types : Number Text

// Latitude see : /latitude
// The latitude of a location. For example 37.42242 (WGS 84 (see: https://schema.orghttps://en.wikipedia.org/wiki/World_Geodetic_System)).
Latitude interface{} `json:"latitude"` // types : Number Text

// Longitude see : /longitude
// The longitude of a location. For example -122.08585 (WGS 84 (see: https://schema.orghttps://en.wikipedia.org/wiki/World_Geodetic_System)).
Longitude interface{} `json:"longitude"` // types : Number Text

// PostalCode see : /postalCode
// The postal code. For example, 94043.
PostalCode string `json:"postalCode"`

}

func (v *GeoCoordinates) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "GeoCoordinates"

	return json.Marshal(v)
}
