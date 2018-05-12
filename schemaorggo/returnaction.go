package schemaorggo

import "encoding/json"

// ReturnAction see : https://schema.org/ReturnAction
type ReturnAction struct {
	TransferAction

	typeContext

	// Recipient see : https://schema.org/recipient
	// A sub property of participant. The participant who is at the receiving end of the action.
	Recipient interface{} `json:"recipient"` // types : Audience ContactPoint Organization Person

}

func (v *ReturnAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ReturnAction"

	return json.Marshal(v)
}
