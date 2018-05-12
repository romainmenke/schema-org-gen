package schemaorggo

import "encoding/json"

// RestrictedDiet see : https://schema.org/RestrictedDiet
type RestrictedDiet struct {
	Enumeration

	typeContext

	// SupersededBy see : http://meta.schema.org/supersededBy
	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	SupersededBy interface{} `json:"supersededBy,omitempty"` // types : Class Enumeration Property

}

func (v RestrictedDiet) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "RestrictedDiet"

	return json.Marshal(v)
}

func (v *RestrictedDiet) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "RestrictedDiet"

	return json.Marshal(*v)
}
