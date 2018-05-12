package schemaorggo

import "encoding/json"

// CommentAction see : https://schema.org/CommentAction
type CommentAction struct {
	CommunicateAction

	typeContext

	// ResultComment see : https://schema.org/resultComment
	// A sub property of result. The Comment created or sent as a result of this action.
	// types : Comment
	ResultComment *Comment `json:"resultComment,omitempty"`
}

func (v CommentAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "CommentAction"

	return json.Marshal(v)
}

func (v *CommentAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "CommentAction"

	return json.Marshal(*v)
}
