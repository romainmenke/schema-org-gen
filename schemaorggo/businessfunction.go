package schemaorggo

import "encoding/json"

// BusinessFunction see : https://schema.org/BusinessFunction
type BusinessFunction struct {
	Enumeration

	typeContext

	// SupersededBy see : http://meta.schema.org/supersededBy
	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	// types : Class Enumeration Property
	SupersededBy interface{} `json:"supersededBy,omitempty"`
}

func (v BusinessFunction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "BusinessFunction"

	return json.Marshal(v)
}

func (v *BusinessFunction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "BusinessFunction"

	return json.Marshal(*v)
}
