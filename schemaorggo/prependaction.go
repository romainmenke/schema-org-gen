package schemaorggo

import "encoding/json"

// PrependAction see : https://schema.org/PrependAction
type PrependAction struct {
	InsertAction

	typeContext

	// ToLocation see : https://schema.org/toLocation
	// A sub property of location. The final location of the object or the agent after the action.
	// types : Place
	ToLocation []*Place `json:"toLocation,omitempty"`
}

func (v PrependAction) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.InsertAction.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.ToLocation
		if len(v.ToLocation) == 1 {
			value = v.ToLocation[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["toLocation"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v PrependAction) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "PrependAction"

	return data, nil
}

func (v PrependAction) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
