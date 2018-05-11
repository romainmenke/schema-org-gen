package schemaorg

import "encoding/json"

// CommentAction see : https://schema.org/CommentAction
type CommentAction struct {

typeContext

CommunicateAction

// ResultComment see : /resultComment
// A sub property of result. The Comment created or sent as a result of this action.
ResultComment *Comment `json:"resultComment"`

}

func (v *CommentAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "CommentAction"

	return json.Marshal(v)
}
