package schemaorggo

import "encoding/json"

// MoveAction see : https://schema.org/MoveAction
type MoveAction struct {
	Action

	typeContext

	// FromLocation see : https://schema.org/fromLocation
	// A sub property of location. The original location of the object or the agent before the action.
	// types : Place
	FromLocation []*Place `json:"fromLocation,omitempty"`

	// ToLocation see : https://schema.org/toLocation
	// A sub property of location. The final location of the object or the agent after the action.
	// types : Place
	ToLocation []*Place `json:"toLocation,omitempty"`
}

func (v MoveAction) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Action.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.FromLocation
		if len(v.FromLocation) == 1 {
			value = v.FromLocation[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["fromLocation"] = json.RawMessage(b)
		}
	}

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

func (v MoveAction) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "MoveAction"

	return data, nil
}

func (v MoveAction) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
