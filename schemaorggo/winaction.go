package schemaorggo

import "encoding/json"

// WinAction see : https://schema.org/WinAction
type WinAction struct {
	AchieveAction

	typeContext

	// Loser see : https://schema.org/loser
	// A sub property of participant. The loser of the action.
	// types : Person
	Loser []*Person `json:"loser,omitempty"`
}

func (v WinAction) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.AchieveAction.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.Loser
		if len(v.Loser) == 1 {
			value = v.Loser[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["loser"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v WinAction) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "WinAction"

	return data, nil
}

func (v WinAction) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
