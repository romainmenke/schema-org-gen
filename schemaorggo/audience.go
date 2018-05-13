package schemaorggo

import "encoding/json"

// Audience see : https://schema.org/Audience
type Audience struct {
	Intangible

	typeContext

	// AudienceType see : https://schema.org/audienceType
	// The target group associated with a given audience (e.g. veterans, car owners, musicians, etc.).
	// types : Text
	AudienceType []string `json:"audienceType,omitempty"`

	// GeographicArea see : https://schema.org/geographicArea
	// The geographic area associated with the audience.
	// types : AdministrativeArea
	GeographicArea []*AdministrativeArea `json:"geographicArea,omitempty"`
}

func (v Audience) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Intangible.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.AudienceType
		if len(v.AudienceType) == 1 {
			value = v.AudienceType[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["audienceType"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.GeographicArea
		if len(v.GeographicArea) == 1 {
			value = v.GeographicArea[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["geographicArea"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v Audience) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Audience"

	return data, nil
}

func (v Audience) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
