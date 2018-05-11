package schemaorg

import "encoding/json"

// DeleteAction see : https://schema.org/DeleteAction
type DeleteAction struct {

typeContext

UpdateAction

// TargetCollection see : /targetCollection
// A sub property of object. The collection target of the action. Supersedes collection (see: https://schema.org/collection).
TargetCollection *Thing `json:"targetCollection"`

}

func (v *DeleteAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "DeleteAction"

	return json.Marshal(v)
}
