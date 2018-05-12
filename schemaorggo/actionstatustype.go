package schemaorggo

import "encoding/json"

// ActionStatusType see : https://schema.org/ActionStatusType
type ActionStatusType struct {
	Enumeration

	typeContext

	// SupersededBy see : http://meta.schema.org/supersededBy
	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	// types : Class Enumeration Property
	SupersededBy interface{} `json:"supersededBy,omitempty"`
}

func (v ActionStatusType) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ActionStatusType"

	return json.Marshal(v)
}

func (v *ActionStatusType) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "ActionStatusType"

	return json.Marshal(*v)
}
