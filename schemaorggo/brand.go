package schemaorggo

import "encoding/json"

// Brand see : https://schema.org/Brand
type Brand struct {
	Intangible

	typeContext
}

func (v Brand) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Intangible.IntoMap(intop)

	into := *intop

	*intop = into

	return nil
}

func (v Brand) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Brand"

	return data, nil
}

func (v Brand) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
