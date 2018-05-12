package schemaorggo

import "encoding/json"

// AddAction see : https://schema.org/AddAction
type AddAction struct {
	UpdateAction

	typeContext

	// TargetCollection see : https://schema.org/targetCollection
	// A sub property of object. The collection target of the action. Supersedes collection (see: https://schema.org/collection).
	TargetCollection *Thing `json:"targetCollection"`
}

func (v *AddAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "AddAction"

	return json.Marshal(v)
}
