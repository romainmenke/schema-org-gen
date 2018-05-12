package schemaorggo

import "encoding/json"

// BookFormatType see : https://schema.org/BookFormatType
type BookFormatType struct {
	Enumeration

	typeContext

	// SupersededBy see : http://meta.schema.org/supersededBy
	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	SupersededBy interface{} `json:"supersededBy,omitempty"` // types : Class Enumeration Property

}

func (v BookFormatType) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "BookFormatType"

	return json.Marshal(v)
}

func (v *BookFormatType) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "BookFormatType"

	return json.Marshal(*v)
}
