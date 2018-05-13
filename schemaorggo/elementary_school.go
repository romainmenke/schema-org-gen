package schemaorggo

import "encoding/json"

// ElementarySchool see : https://schema.org/ElementarySchool
type ElementarySchool struct {
	EducationalOrganization

	typeContext

	// Alumni see : https://schema.org/alumni
	// Alumni of an organization. Inverse property: alumniOf (see: https://schema.org/alumniOf).
	// types : Person
	Alumni []*Person `json:"alumni,omitempty"`
}

func (v ElementarySchool) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.EducationalOrganization.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.Alumni
		if len(v.Alumni) == 1 {
			value = v.Alumni[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["alumni"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v ElementarySchool) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "ElementarySchool"

	return data, nil
}

func (v ElementarySchool) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}