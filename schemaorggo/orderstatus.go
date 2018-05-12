package schemaorggo

import "encoding/json"

// OrderStatus see : https://schema.org/OrderStatus
type OrderStatus struct {
	Enumeration

	typeContext

	// SupersededBy see : http://meta.schema.org/supersededBy
	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	SupersededBy interface{} `json:"supersededBy,omitempty"` // types : Class Enumeration Property

}

func (v OrderStatus) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "OrderStatus"

	return json.Marshal(v)
}

func (v *OrderStatus) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "OrderStatus"

	return json.Marshal(*v)
}
