package schemaorg

import "encoding/json"

// AuthorizeAction see : https://schema.org/AuthorizeAction
type AuthorizeAction struct {

typeContext

AllocateAction

// Recipient see : https://schema.org/recipient
// A sub property of participant. The participant who is at the receiving end of the action.
Recipient interface{} `json:"recipient"` // types : Audience ContactPoint Organization Person

}

func (v *AuthorizeAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "AuthorizeAction"

	return json.Marshal(v)
}
