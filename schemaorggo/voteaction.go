package schemaorggo

import "encoding/json"

// VoteAction see : https://schema.org/VoteAction
type VoteAction struct {
	ChooseAction

	typeContext

	// Candidate see : https://schema.org/candidate
	// A sub property of object. The candidate subject of this action.
	// types : Person
	Candidate []*Person `json:"candidate,omitempty"`
}

func (v VoteAction) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.ChooseAction.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.Candidate
		if len(v.Candidate) == 1 {
			value = v.Candidate[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["candidate"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v VoteAction) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "VoteAction"

	return data, nil
}

func (v VoteAction) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
