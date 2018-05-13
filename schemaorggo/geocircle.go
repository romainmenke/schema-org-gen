package schemaorggo

import "encoding/json"

// GeoCircle see : https://schema.org/GeoCircle
type GeoCircle struct {
	GeoShape

	typeContext

	// GeoMidpoint see : https://schema.org/geoMidpoint
	// Indicates the GeoCoordinates at the centre of a GeoShape e.g. GeoCircle.
	// types : GeoCoordinates
	GeoMidpoint []*GeoCoordinates `json:"geoMidpoint,omitempty"`

	// GeoRadius see : https://schema.org/geoRadius
	// Indicates the approximate radius of a GeoCircle (metres unless indicated otherwise via Distance notation).
	// types : Distance Number Text
	GeoRadius []interface{} `json:"geoRadius,omitempty"`
}

func (v GeoCircle) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.GeoShape.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.GeoMidpoint
		if len(v.GeoMidpoint) == 1 {
			value = v.GeoMidpoint[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["geoMidpoint"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.GeoRadius
		if len(v.GeoRadius) == 1 {
			value = v.GeoRadius[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["geoRadius"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v GeoCircle) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "GeoCircle"

	return data, nil
}

func (v GeoCircle) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
