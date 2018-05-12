package schemaorg

import "encoding/json"

// GiveAction see : https://schema.org/GiveAction
type GiveAction struct {

TransferAction

typeContext

// Recipient see : https://schema.org/recipient
// A sub property of participant. The participant who is at the receiving end of the action.
Recipient interface{} `json:"recipient"` // types : Audience ContactPoint Organization Person

}

func (v *GiveAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "GiveAction"

	return json.Marshal(v)
}
