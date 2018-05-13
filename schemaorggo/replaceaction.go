package schemaorggo

import "encoding/json"

// ReplaceAction see : https://schema.org/ReplaceAction
type ReplaceAction struct {
	UpdateAction

	typeContext

	// Replacee see : https://schema.org/replacee
	// A sub property of object. The object that is being replaced.
	// types : Thing
	Replacee []*Thing `json:"replacee,omitempty"`

	// Replacer see : https://schema.org/replacer
	// A sub property of object. The object that replaces.
	// types : Thing
	Replacer []*Thing `json:"replacer,omitempty"`
}

func (v ReplaceAction) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.UpdateAction.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.Replacee
		if len(v.Replacee) == 1 {
			value = v.Replacee[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["replacee"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Replacer
		if len(v.Replacer) == 1 {
			value = v.Replacer[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["replacer"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v ReplaceAction) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "ReplaceAction"

	return data, nil
}

func (v ReplaceAction) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
