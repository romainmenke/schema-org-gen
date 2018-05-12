package schemaorg

import "encoding/json"

// AskAction see : https://schema.org/AskAction
type AskAction struct {

CommunicateAction

typeContext

// Question see : https://schema.org/question
// A sub property of object. A question.
Question *Question `json:"question"`

}

func (v *AskAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "AskAction"

	return json.Marshal(v)
}
