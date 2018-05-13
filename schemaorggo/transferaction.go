package schemaorggo

import "encoding/json"

// TransferAction see : https://schema.org/TransferAction
type TransferAction struct {
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

func (v TransferAction) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Action.IntoMap(intop)

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

func (v TransferAction) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "TransferAction"

	return data, nil
}

func (v TransferAction) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
