package schemaorg

import "encoding/json"

// OrderStatus see : https://schema.org/OrderStatus
type OrderStatus struct {

typeContext

Enumeration

// SupersededBy see : https://schema.orghttp://meta.schema.org/supersededBy
// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
SupersededBy interface{} `json:"supersededBy"` // types : Class Enumeration Property

}

func (v *OrderStatus) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "OrderStatus"

	return json.Marshal(v)
}
