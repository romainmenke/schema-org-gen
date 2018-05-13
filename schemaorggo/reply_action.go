package schemaorggo

import "encoding/json"

// ReplyAction see : https://schema.org/ReplyAction
type ReplyAction struct {
	CommunicateAction

	typeContext

	// ResultComment see : https://schema.org/resultComment
	// A sub property of result. The Comment created or sent as a result of this action.
	// types : Comment
	ResultComment []*Comment `json:"resultComment,omitempty"`
}

func (v ReplyAction) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.CommunicateAction.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.ResultComment
		if len(v.ResultComment) == 1 {
			value = v.ResultComment[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["resultComment"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v ReplyAction) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "ReplyAction"

	return data, nil
}

func (v ReplyAction) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
