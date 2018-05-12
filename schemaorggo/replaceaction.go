package schemaorggo

import "encoding/json"

// ReplaceAction see : https://schema.org/ReplaceAction
type ReplaceAction struct {
	UpdateAction

	typeContext

	// Replacee see : https://schema.org/replacee
	// A sub property of object. The object that is being replaced.
	// types : Thing
	Replacee *Thing `json:"replacee,omitempty"`

	// Replacer see : https://schema.org/replacer
	// A sub property of object. The object that replaces.
	// types : Thing
	Replacer *Thing `json:"replacer,omitempty"`
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
