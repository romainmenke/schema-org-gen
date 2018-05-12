package schemaorggo

import "encoding/json"

// GenderType see : https://schema.org/GenderType
type GenderType struct {
	Enumeration

	typeContext

	// SupersededBy see : http://meta.schema.org/supersededBy
	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	SupersededBy interface{} `json:"supersededBy"` // types : Class Enumeration Property

}

func (v *GenderType) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "GenderType"

	return json.Marshal(v)
}
