package schemaorggo

import "encoding/json"

// BusinessEntityType see : https://schema.org/BusinessEntityType
type BusinessEntityType struct {
	Enumeration

	typeContext

	// SupersededBy see : http://meta.schema.org/supersededBy
	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	// types : Class Enumeration Property
	SupersededBy interface{} `json:"supersededBy,omitempty"`
}

func (v BusinessEntityType) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "BusinessEntityType"

	return json.Marshal(v)
}

func (v *BusinessEntityType) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "BusinessEntityType"

	return json.Marshal(*v)
}
