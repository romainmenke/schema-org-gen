package schemaorggo

import "encoding/json"

// CommentAction see : https://schema.org/CommentAction
type CommentAction struct {
	CommunicateAction

	typeContext

	// ResultComment see : https://schema.org/resultComment
	// A sub property of result. The Comment created or sent as a result of this action.
	// types : Comment
	ResultComment []*Comment `json:"resultComment,omitempty"`
}

func (v CommentAction) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.CommunicateAction.IntoMap(intop)

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

func (v CommentAction) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "CommentAction"

	return data, nil
}

func (v CommentAction) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
