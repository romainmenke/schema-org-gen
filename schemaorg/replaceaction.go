package schemaorg

import "encoding/json"

// ReplaceAction see : https://schema.org/ReplaceAction
type ReplaceAction struct {

typeContext

UpdateAction

// Replacee see : /replacee
// A sub property of object. The object that is being replaced.
Replacee *Thing `json:"replacee"`

// Replacer see : /replacer
// A sub property of object. The object that replaces.
Replacer *Thing `json:"replacer"`

}

func (v *ReplaceAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ReplaceAction"

	return json.Marshal(v)
}
