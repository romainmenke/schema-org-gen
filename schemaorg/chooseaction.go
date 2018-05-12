package schemaorg

import "encoding/json"

// ChooseAction see : https://schema.org/ChooseAction
type ChooseAction struct {

AssessAction

typeContext

// ActionOption see : https://schema.org/actionOption
// A sub property of object. The options subject to this action. Supersedes option (see: https://schema.org/option).
ActionOption interface{} `json:"actionOption"` // types : Text Thing

}

func (v *ChooseAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ChooseAction"

	return json.Marshal(v)
}
