package schemaorggo

import "encoding/json"

// LoseAction see : https://schema.org/LoseAction
type LoseAction struct {
	AchieveAction

	typeContext

	// Winner see : https://schema.org/winner
	// A sub property of participant. The winner of the action.
	// types : Person
	Winner []*Person `json:"winner,omitempty"`
}

func (v LoseAction) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.AchieveAction.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.Winner
		if len(v.Winner) == 1 {
			value = v.Winner[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["winner"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v LoseAction) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "LoseAction"

	return data, nil
}

func (v LoseAction) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
