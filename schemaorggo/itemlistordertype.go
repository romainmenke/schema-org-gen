package schemaorggo

import "encoding/json"

// ItemListOrderType see : https://schema.org/ItemListOrderType
type ItemListOrderType struct {
	Enumeration

	typeContext

	// SupersededBy see : http://meta.schema.org/supersededBy
	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	// types : Class Enumeration Property
	SupersededBy interface{} `json:"supersededBy,omitempty"`
}

func (v ItemListOrderType) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ItemListOrderType"

	return json.Marshal(v)
}

func (v *ItemListOrderType) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "ItemListOrderType"

	return json.Marshal(*v)
}
