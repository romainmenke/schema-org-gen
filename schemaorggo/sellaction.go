package schemaorggo

import "encoding/json"

// SellAction see : https://schema.org/SellAction
type SellAction struct {
	TradeAction

	typeContext

	// Buyer see : https://schema.org/buyer
	// A sub property of participant. The participant/person/organization that bought the object.
	// types : Person
	Buyer *Person `json:"buyer,omitempty"`
}

func (v SellAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "SellAction"

	return json.Marshal(v)
}

func (v *SellAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "SellAction"

	return json.Marshal(*v)
}
