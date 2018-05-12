package schemaorggo

import "encoding/json"

// ItemAvailability see : https://schema.org/ItemAvailability
type ItemAvailability struct {
	Enumeration

	typeContext

	// SupersededBy see : http://meta.schema.org/supersededBy
	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	// types : Class Enumeration Property
	SupersededBy interface{} `json:"supersededBy,omitempty"`
}

func (v ItemAvailability) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ItemAvailability"

	return json.Marshal(v)
}

func (v *ItemAvailability) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "ItemAvailability"

	return json.Marshal(*v)
}
