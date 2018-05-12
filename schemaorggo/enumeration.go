package schemaorggo

import "encoding/json"

// Enumeration see : https://schema.org/Enumeration
type Enumeration struct {
	Intangible

	typeContext

	// SupersededBy see : http://meta.schema.org/supersededBy
	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	SupersededBy interface{} `json:"supersededBy,omitempty"` // types : Class Enumeration Property

}

func (v Enumeration) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Enumeration"

	return json.Marshal(v)
}

func (v *Enumeration) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "Enumeration"

	return json.Marshal(*v)
}
