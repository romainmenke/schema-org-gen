package schemaorggo

import "encoding/json"

// RsvpResponseType see : https://schema.org/RsvpResponseType
type RsvpResponseType struct {
	Enumeration

	typeContext

	// SupersededBy see : http://meta.schema.org/supersededBy
	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	SupersededBy interface{} `json:"supersededBy"` // types : Class Enumeration Property

}

func (v *RsvpResponseType) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "RsvpResponseType"

	return json.Marshal(v)
}
