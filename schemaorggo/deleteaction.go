package schemaorggo

import "encoding/json"

// DeleteAction see : https://schema.org/DeleteAction
type DeleteAction struct {
	UpdateAction

	typeContext

	// TargetCollection see : https://schema.org/targetCollection
	// A sub property of object. The collection target of the action. Supersedes collection (see: https://schema.org/collection).
	TargetCollection *Thing `json:"targetCollection,omitempty"`
}

func (v DeleteAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "DeleteAction"

	return json.Marshal(v)
}

func (v *DeleteAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "DeleteAction"

	return json.Marshal(*v)
}
