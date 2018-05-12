package schemaorggo

import "encoding/json"

// UpdateAction see : https://schema.org/UpdateAction
type UpdateAction struct {
	Action

	typeContext

	// TargetCollection see : https://schema.org/targetCollection
	// A sub property of object. The collection target of the action. Supersedes collection (see: https://schema.org/collection).
	TargetCollection *Thing `json:"targetCollection"`
}

func (v *UpdateAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "UpdateAction"

	return json.Marshal(v)
}
