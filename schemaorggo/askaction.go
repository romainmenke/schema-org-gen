package schemaorggo

import "encoding/json"

// AskAction see : https://schema.org/AskAction
type AskAction struct {
	CommunicateAction

	typeContext

	// Question see : https://schema.org/question
	// A sub property of object. A question.
	// types : Question
	Question *Question `json:"question,omitempty"`
}

func (v AskAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "AskAction"

	return json.Marshal(v)
}

func (v *AskAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "AskAction"

	return json.Marshal(*v)
}
