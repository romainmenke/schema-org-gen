package schemaorggo

import "encoding/json"

// UpdateAction see : https://schema.org/UpdateAction
type UpdateAction struct {
	Action

	typeContext

	// TargetCollection see : https://schema.org/targetCollection
	// A sub property of object. The collection target of the action. Supersedes collection (see: https://schema.org/collection).
	TargetCollection *Thing `json:"targetCollection,omitempty"` // types : Thing

}

func (v UpdateAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "UpdateAction"

	return json.Marshal(v)
}

func (v *UpdateAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "UpdateAction"

	return json.Marshal(*v)
}
