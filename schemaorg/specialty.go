package schemaorg

import "encoding/json"

// Specialty see : https://schema.org/Specialty
type Specialty struct {

Enumeration

typeContext

// SupersededBy see : http://meta.schema.org/supersededBy
// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
SupersededBy interface{} `json:"supersededBy"` // types : Class Enumeration Property

}

func (v *Specialty) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Specialty"

	return json.Marshal(v)
}
