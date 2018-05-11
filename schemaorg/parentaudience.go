package schemaorg

import "encoding/json"

// ParentAudience see : https://schema.org/ParentAudience
type ParentAudience struct {

typeContext

PeopleAudience

// ChildMaxAge see : /childMaxAge
// Maximal age of the child.
ChildMaxAge float64 `json:"childMaxAge"`

// ChildMinAge see : /childMinAge
// Minimal age of the child.
ChildMinAge float64 `json:"childMinAge"`

}

func (v *ParentAudience) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ParentAudience"

	return json.Marshal(v)
}
