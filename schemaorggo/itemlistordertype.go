package schemaorggo

import "encoding/json"

// ItemListOrderType see : https://schema.org/ItemListOrderType
type ItemListOrderType struct {
	Enumeration

	typeContext

	// SupersededBy see : http://meta.schema.org/supersededBy
	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	SupersededBy interface{} `json:"supersededBy"` // types : Class Enumeration Property

}

func (v *ItemListOrderType) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ItemListOrderType"

	return json.Marshal(v)
}
