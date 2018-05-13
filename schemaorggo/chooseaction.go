package schemaorggo

import "encoding/json"

// ChooseAction see : https://schema.org/ChooseAction
type ChooseAction struct {
	AssessAction

	typeContext

	// ActionOption see : https://schema.org/actionOption
	// A sub property of object. The options subject to this action. Supersedes option (see: https://schema.org/option).
	// types : Text Thing
	ActionOption []interface{} `json:"actionOption,omitempty"`
}

func (v ChooseAction) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.AssessAction.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.ActionOption
		if len(v.ActionOption) == 1 {
			value = v.ActionOption[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["actionOption"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v ChooseAction) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "ChooseAction"

	return data, nil
}

func (v ChooseAction) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
