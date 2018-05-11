package schemaorg

import "encoding/json"

// TipAction see : https://schema.org/TipAction
type TipAction struct {

typeContext

TradeAction

// Recipient see : https://schema.org/recipient
// A sub property of participant. The participant who is at the receiving end of the action.
Recipient interface{} `json:"recipient"` // types : Audience ContactPoint Organization Person

}

func (v *TipAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "TipAction"

	return json.Marshal(v)
}
