package schemaorg

import "encoding/json"

// DonateAction see : https://schema.org/DonateAction
type DonateAction struct {

typeContext

TradeAction

// Recipient see : https://schema.org/recipient
// A sub property of participant. The participant who is at the receiving end of the action.
Recipient interface{} `json:"recipient"` // types : Audience ContactPoint Organization Person

}

func (v *DonateAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "DonateAction"

	return json.Marshal(v)
}
