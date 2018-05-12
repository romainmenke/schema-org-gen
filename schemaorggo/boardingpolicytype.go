package schemaorggo

import "encoding/json"

// BoardingPolicyType see : https://schema.org/BoardingPolicyType
type BoardingPolicyType struct {
	Enumeration

	typeContext

	// SupersededBy see : http://meta.schema.org/supersededBy
	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	// types : Class Enumeration Property
	SupersededBy interface{} `json:"supersededBy,omitempty"`
}

func (v BoardingPolicyType) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "BoardingPolicyType"

	return json.Marshal(v)
}

func (v *BoardingPolicyType) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "BoardingPolicyType"

	return json.Marshal(*v)
}
