package schemaorggo

import "encoding/json"

// Specialty see : https://schema.org/Specialty
type Specialty struct {
	Enumeration

	typeContext

	// SupersededBy see : http://meta.schema.org/supersededBy
	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	// types : Class Enumeration Property
	SupersededBy interface{} `json:"supersededBy,omitempty"`
}

func (v Specialty) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Specialty"

	return json.Marshal(v)
}

func (v *Specialty) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "Specialty"

	return json.Marshal(*v)
}
