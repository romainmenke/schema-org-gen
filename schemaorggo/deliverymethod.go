package schemaorggo

import "encoding/json"

// DeliveryMethod see : https://schema.org/DeliveryMethod
type DeliveryMethod struct {
	Enumeration

	typeContext

	// SupersededBy see : http://meta.schema.org/supersededBy
	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	// types : Class Enumeration Property
	SupersededBy interface{} `json:"supersededBy,omitempty"`
}

func (v DeliveryMethod) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "DeliveryMethod"

	return json.Marshal(v)
}

func (v *DeliveryMethod) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "DeliveryMethod"

	return json.Marshal(*v)
}
