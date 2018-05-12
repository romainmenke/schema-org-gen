package schemaorg

import "encoding/json"

// PaymentMethod see : https://schema.org/PaymentMethod
type PaymentMethod struct {

Enumeration

typeContext

// SupersededBy see : http://meta.schema.org/supersededBy
// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
SupersededBy interface{} `json:"supersededBy"` // types : Class Enumeration Property

}

func (v *PaymentMethod) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "PaymentMethod"

	return json.Marshal(v)
}
