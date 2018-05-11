package schemaorg

import "encoding/json"

// GeoCircle see : https://schema.org/GeoCircle
type GeoCircle struct {

typeContext

GeoShape

// GeoMidpoint see : /geoMidpoint
// Indicates the GeoCoordinates at the centre of a GeoShape e.g. GeoCircle.
GeoMidpoint *GeoCoordinates `json:"geoMidpoint"`

// GeoRadius see : /geoRadius
// Indicates the approximate radius of a GeoCircle (metres unless indicated otherwise via Distance notation).
GeoRadius interface{} `json:"geoRadius"` // types : Distance Number Text

}

func (v *GeoCircle) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "GeoCircle"

	return json.Marshal(v)
}
