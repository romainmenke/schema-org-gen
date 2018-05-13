package schemaorggo

import "encoding/json"

// GeoShape see : https://schema.org/GeoShape
type GeoShape struct {
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

	// Box see : https://schema.org/box
	// A box is the area enclosed by the rectangle formed by two points. The first point is the lower corner, the second point is the upper corner. A box is expressed as two points separated by a space character.
	// types : Text
	Box []string `json:"box,omitempty"`

	// Circle see : https://schema.org/circle
	// A circle is the circular region of a specified radius centered at a specified latitude and longitude. A circle is expressed as a pair followed by a radius in meters.
	// types : Text
	Circle []string `json:"circle,omitempty"`

	// Elevation see : https://schema.org/elevation
	// The elevation of a location (WGS 84 (see: https://schema.orghttps://en.wikipedia.org/wiki/World_Geodetic_System)).
	// types : Number Text
	Elevation []interface{} `json:"elevation,omitempty"`

	// Line see : https://schema.org/line
	// A line is a point-to-point path consisting of two or more points. A line is expressed as a series of two or more point objects separated by space.
	// types : Text
	Line []string `json:"line,omitempty"`

	// Polygon see : https://schema.org/polygon
	// A polygon is the area enclosed by a point-to-point path for which the starting and ending points are the same. A polygon is expressed as a series of four or more space delimited points where the first and final points are identical.
	// types : Text
	Polygon []string `json:"polygon,omitempty"`

	// PostalCode see : https://schema.org/postalCode
	// The postal code. For example, 94043.
	// types : Text
	PostalCode []string `json:"postalCode,omitempty"`
}

func (v GeoShape) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.StructuredValue.IntoMap(intop)

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
		var value interface{} = v.Box
		if len(v.Box) == 1 {
			value = v.Box[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["box"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Circle
		if len(v.Circle) == 1 {
			value = v.Circle[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["circle"] = json.RawMessage(b)
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
		var value interface{} = v.Line
		if len(v.Line) == 1 {
			value = v.Line[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["line"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Polygon
		if len(v.Polygon) == 1 {
			value = v.Polygon[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["polygon"] = json.RawMessage(b)
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

func (v GeoShape) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "GeoShape"

	return data, nil
}

func (v GeoShape) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
