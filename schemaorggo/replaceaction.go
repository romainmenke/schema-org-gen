package schemaorggo

import "encoding/json"

// ReplaceAction see : https://schema.org/ReplaceAction
type ReplaceAction struct {
	UpdateAction

	typeContext

	// Replacee see : https://schema.org/replacee
	// A sub property of object. The object that is being replaced.
	Replacee *Thing `json:"replacee,omitempty"` // types : Thing

	// Replacer see : https://schema.org/replacer
	// A sub property of object. The object that replaces.
	Replacer *Thing `json:"replacer,omitempty"` // types : Thing

}

func (v ReplaceAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ReplaceAction"

	return json.Marshal(v)
}

func (v *ReplaceAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "ReplaceAction"

	return json.Marshal(*v)
}
