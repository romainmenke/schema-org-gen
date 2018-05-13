package schemaorggo

import "encoding/json"

// AskAction see : https://schema.org/AskAction
type AskAction struct {
	CommunicateAction

	typeContext

	// Question see : https://schema.org/question
	// A sub property of object. A question.
	// types : Question
	Question []*Question `json:"question,omitempty"`
}

func (v AskAction) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.CommunicateAction.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.Question
		if len(v.Question) == 1 {
			value = v.Question[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["question"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v AskAction) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "AskAction"

	return data, nil
}

func (v AskAction) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
