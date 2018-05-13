package schemaorggo

import "encoding/json"

// AddAction see : https://schema.org/AddAction
type AddAction struct {
	UpdateAction

	typeContext

	// TargetCollection see : https://schema.org/targetCollection
	// A sub property of object. The collection target of the action. Supersedes collection (see: https://schema.org/collection).
	// types : Thing
	TargetCollection []*Thing `json:"targetCollection,omitempty"`
}

func (v AddAction) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.UpdateAction.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.TargetCollection
		if len(v.TargetCollection) == 1 {
			value = v.TargetCollection[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["targetCollection"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v AddAction) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "AddAction"

	return data, nil
}

func (v AddAction) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
