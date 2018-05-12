package schemaorggo

import "encoding/json"

// GeoCircle see : https://schema.org/GeoCircle
type GeoCircle struct {
	GeoShape

	typeContext

	// GeoMidpoint see : https://schema.org/geoMidpoint
	// Indicates the GeoCoordinates at the centre of a GeoShape e.g. GeoCircle.
	// types : GeoCoordinates
	GeoMidpoint *GeoCoordinates `json:"geoMidpoint,omitempty"`

	// GeoRadius see : https://schema.org/geoRadius
	// Indicates the approximate radius of a GeoCircle (metres unless indicated otherwise via Distance notation).
	// types : Distance Number Text
	GeoRadius interface{} `json:"geoRadius,omitempty"`
}

func (v GeoCircle) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "GeoCircle"

	return json.Marshal(v)
}

func (v *GeoCircle) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "GeoCircle"

	return json.Marshal(*v)
}
