package schemaorggo

import "encoding/json"

// ReplyAction see : https://schema.org/ReplyAction
type ReplyAction struct {
	CommunicateAction

	typeContext

	// ResultComment see : https://schema.org/resultComment
	// A sub property of result. The Comment created or sent as a result of this action.
	// types : Comment
	ResultComment *Comment `json:"resultComment,omitempty"`
}

func (v ReplyAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ReplyAction"

	return json.Marshal(v)
}

func (v *ReplyAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "ReplyAction"

	return json.Marshal(*v)
}
