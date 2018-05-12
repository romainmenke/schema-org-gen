package schemaorggo

import "encoding/json"

// ContactPointOption see : https://schema.org/ContactPointOption
type ContactPointOption struct {
	Enumeration

	typeContext

	// SupersededBy see : http://meta.schema.org/supersededBy
	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	SupersededBy interface{} `json:"supersededBy,omitempty"` // types : Class Enumeration Property

}

func (v ContactPointOption) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ContactPointOption"

	return json.Marshal(v)
}

func (v *ContactPointOption) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "ContactPointOption"

	return json.Marshal(*v)
}
