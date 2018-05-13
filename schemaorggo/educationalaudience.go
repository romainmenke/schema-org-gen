package schemaorggo

import "encoding/json"

// EducationalAudience see : https://schema.org/EducationalAudience
type EducationalAudience struct {
	Audience

	typeContext

	// EducationalRole see : https://schema.org/educationalRole
	// An educationalRole of an EducationalAudience.
	// types : Text
	EducationalRole []string `json:"educationalRole,omitempty"`
}

func (v EducationalAudience) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Audience.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.EducationalRole
		if len(v.EducationalRole) == 1 {
			value = v.EducationalRole[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["educationalRole"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v EducationalAudience) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "EducationalAudience"

	return data, nil
}

func (v EducationalAudience) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
