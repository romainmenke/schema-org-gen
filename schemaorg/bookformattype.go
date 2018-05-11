package schemaorg

import "encoding/json"

// BookFormatType see : https://schema.org/BookFormatType
type BookFormatType struct {

typeContext

Enumeration

// SupersededBy see : https://schema.orghttp://meta.schema.org/supersededBy
// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
SupersededBy interface{} `json:"supersededBy"` // types : Class Enumeration Property

}

func (v *BookFormatType) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "BookFormatType"

	return json.Marshal(v)
}
