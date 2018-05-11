package schemaorg

import "encoding/json"

// DeliveryMethod see : https://schema.org/DeliveryMethod
type DeliveryMethod struct {

typeContext

Enumeration

// SupersededBy see : https://schema.orghttp://meta.schema.org/supersededBy
// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
SupersededBy interface{} `json:"supersededBy"` // types : Class Enumeration Property

}

func (v *DeliveryMethod) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "DeliveryMethod"

	return json.Marshal(v)
}
