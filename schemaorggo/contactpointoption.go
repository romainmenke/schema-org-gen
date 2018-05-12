package schemaorggo

import "encoding/json"

// ContactPointOption see : https://schema.org/ContactPointOption
type ContactPointOption struct {
	Enumeration

	typeContext

	// SupersededBy see : http://meta.schema.org/supersededBy
	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	SupersededBy interface{} `json:"supersededBy"` // types : Class Enumeration Property

}

func (v *ContactPointOption) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ContactPointOption"

	return json.Marshal(v)
}
