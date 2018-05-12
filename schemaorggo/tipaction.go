package schemaorggo

import "encoding/json"

// TipAction see : https://schema.org/TipAction
type TipAction struct {
	TradeAction

	typeContext

	// Recipient see : https://schema.org/recipient
	// A sub property of participant. The participant who is at the receiving end of the action.
	// types : Audience ContactPoint Organization Person
	Recipient interface{} `json:"recipient,omitempty"`
}

func (v TipAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "TipAction"

	return json.Marshal(v)
}

func (v *TipAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "TipAction"

	return json.Marshal(*v)
}
