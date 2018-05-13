package schemaorggo

import "encoding/json"

// PerformanceRole see : https://schema.org/PerformanceRole
type PerformanceRole struct {
	Role

	typeContext

	// CharacterName see : https://schema.org/characterName
	// The name of a character played in some acting or performing role, i.e. in a PerformanceRole.
	// types : Text
	CharacterName []string `json:"characterName,omitempty"`
}

func (v PerformanceRole) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Role.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.CharacterName
		if len(v.CharacterName) == 1 {
			value = v.CharacterName[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["characterName"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v PerformanceRole) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "PerformanceRole"

	return data, nil
}

func (v PerformanceRole) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
