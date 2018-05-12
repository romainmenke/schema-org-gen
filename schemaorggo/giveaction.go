package schemaorggo

import "encoding/json"

// GiveAction see : https://schema.org/GiveAction
type GiveAction struct {
	TransferAction

	typeContext

	// Recipient see : https://schema.org/recipient
	// A sub property of participant. The participant who is at the receiving end of the action.
	// types : Audience ContactPoint Organization Person
	Recipient interface{} `json:"recipient,omitempty"`
}

func (v GiveAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "GiveAction"

	return json.Marshal(v)
}

func (v *GiveAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "GiveAction"

	return json.Marshal(*v)
}
