package schemaorggo

import "encoding/json"

// RsvpResponseType see : https://schema.org/RsvpResponseType
type RsvpResponseType struct {
	Enumeration

	typeContext

	// SupersededBy see : http://meta.schema.org/supersededBy
	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	// types : Class Enumeration Property
	SupersededBy interface{} `json:"supersededBy,omitempty"`
}

func (v RsvpResponseType) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "RsvpResponseType"

	return json.Marshal(v)
}

func (v *RsvpResponseType) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "RsvpResponseType"

	return json.Marshal(*v)
}
