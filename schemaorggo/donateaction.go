package schemaorggo

import "encoding/json"

// DonateAction see : https://schema.org/DonateAction
type DonateAction struct {
	TradeAction

	typeContext

	// Recipient see : https://schema.org/recipient
	// A sub property of participant. The participant who is at the receiving end of the action.
	// types : Audience ContactPoint Organization Person
	Recipient interface{} `json:"recipient,omitempty"`
}

func (v DonateAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "DonateAction"

	return json.Marshal(v)
}

func (v *DonateAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "DonateAction"

	return json.Marshal(*v)
}
