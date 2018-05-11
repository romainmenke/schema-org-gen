package schemaorg

import "encoding/json"

// GeoShape see : https://schema.org/GeoShape
type GeoShape struct {

typeContext

StructuredValue

// Address see : /address
// Physical address of the item.
Address interface{} `json:"address"` // types : PostalAddress Text

// AddressCountry see : /addressCountry
// The country. For example, USA. You can also provide the two-letter ISO 3166-1 alpha-2 country code (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_3166-1).
AddressCountry interface{} `json:"addressCountry"` // types : Country Text

// Box see : /box
// A box is the area enclosed by the rectangle formed by two points. The first point is the lower corner, the second point is the upper corner. A box is expressed as two points separated by a space character.
Box string `json:"box"`

// Circle see : /circle
// A circle is the circular region of a specified radius centered at a specified latitude and longitude. A circle is expressed as a pair followed by a radius in meters.
Circle string `json:"circle"`

// Elevation see : /elevation
// The elevation of a location (WGS 84 (see: https://schema.orghttps://en.wikipedia.org/wiki/World_Geodetic_System)).
Elevation interface{} `json:"elevation"` // types : Number Text

// Line see : /line
// A line is a point-to-point path consisting of two or more points. A line is expressed as a series of two or more point objects separated by space.
Line string `json:"line"`

// Polygon see : /polygon
// A polygon is the area enclosed by a point-to-point path for which the starting and ending points are the same. A polygon is expressed as a series of four or more space delimited points where the first and final points are identical.
Polygon string `json:"polygon"`

// PostalCode see : /postalCode
// The postal code. For example, 94043.
PostalCode string `json:"postalCode"`

}

func (v *GeoShape) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "GeoShape"

	return json.Marshal(v)
}
