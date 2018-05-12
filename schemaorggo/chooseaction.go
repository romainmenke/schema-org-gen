package schemaorggo

import "encoding/json"

// ChooseAction see : https://schema.org/ChooseAction
type ChooseAction struct {
	AssessAction

	typeContext

	// ActionOption see : https://schema.org/actionOption
	// A sub property of object. The options subject to this action. Supersedes option (see: https://schema.org/option).
	ActionOption interface{} `json:"actionOption,omitempty"` // types : Text Thing

}

func (v ChooseAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ChooseAction"

	return json.Marshal(v)
}

func (v *ChooseAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "ChooseAction"

	return json.Marshal(*v)
}
