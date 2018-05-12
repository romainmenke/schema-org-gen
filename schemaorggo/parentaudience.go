package schemaorggo

import "encoding/json"

// ParentAudience see : https://schema.org/ParentAudience
type ParentAudience struct {
	PeopleAudience

	typeContext

	// ChildMaxAge see : https://schema.org/childMaxAge
	// Maximal age of the child.
	ChildMaxAge float64 `json:"childMaxAge,omitempty"` // types : Number

	// ChildMinAge see : https://schema.org/childMinAge
	// Minimal age of the child.
	ChildMinAge float64 `json:"childMinAge,omitempty"` // types : Number

}

func (v ParentAudience) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ParentAudience"

	return json.Marshal(v)
}

func (v *ParentAudience) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "ParentAudience"

	return json.Marshal(*v)
}
