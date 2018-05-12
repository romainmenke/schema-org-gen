package schemaorggo

import "encoding/json"

// GeoShape see : https://schema.org/GeoShape
type GeoShape struct {
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

	// Box see : https://schema.org/box
	// A box is the area enclosed by the rectangle formed by two points. The first point is the lower corner, the second point is the upper corner. A box is expressed as two points separated by a space character.
	// types : Text
	Box string `json:"box,omitempty"`

	// Circle see : https://schema.org/circle
	// A circle is the circular region of a specified radius centered at a specified latitude and longitude. A circle is expressed as a pair followed by a radius in meters.
	// types : Text
	Circle string `json:"circle,omitempty"`

	// Elevation see : https://schema.org/elevation
	// The elevation of a location (WGS 84 (see: https://schema.orghttps://en.wikipedia.org/wiki/World_Geodetic_System)).
	// types : Number Text
	Elevation interface{} `json:"elevation,omitempty"`

	// Line see : https://schema.org/line
	// A line is a point-to-point path consisting of two or more points. A line is expressed as a series of two or more point objects separated by space.
	// types : Text
	Line string `json:"line,omitempty"`

	// Polygon see : https://schema.org/polygon
	// A polygon is the area enclosed by a point-to-point path for which the starting and ending points are the same. A polygon is expressed as a series of four or more space delimited points where the first and final points are identical.
	// types : Text
	Polygon string `json:"polygon,omitempty"`

	// PostalCode see : https://schema.org/postalCode
	// The postal code. For example, 94043.
	// types : Text
	PostalCode string `json:"postalCode,omitempty"`
}

func (v GeoShape) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "GeoShape"

	return json.Marshal(v)
}

func (v *GeoShape) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "GeoShape"

	return json.Marshal(*v)
}
