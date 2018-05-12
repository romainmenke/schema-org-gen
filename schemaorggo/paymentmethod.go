package schemaorggo

import "encoding/json"

// PaymentMethod see : https://schema.org/PaymentMethod
type PaymentMethod struct {
	Enumeration

	typeContext

	// SupersededBy see : http://meta.schema.org/supersededBy
	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	// types : Class Enumeration Property
	SupersededBy interface{} `json:"supersededBy,omitempty"`
}

func (v PaymentMethod) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "PaymentMethod"

	return json.Marshal(v)
}

func (v *PaymentMethod) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "PaymentMethod"

	return json.Marshal(*v)
}
