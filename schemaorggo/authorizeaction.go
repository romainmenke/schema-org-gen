package schemaorggo

import "encoding/json"

// AuthorizeAction see : https://schema.org/AuthorizeAction
type AuthorizeAction struct {
	AllocateAction

	typeContext

	// Recipient see : https://schema.org/recipient
	// A sub property of participant. The participant who is at the receiving end of the action.
	Recipient interface{} `json:"recipient,omitempty"` // types : Audience ContactPoint Organization Person

}

func (v AuthorizeAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "AuthorizeAction"

	return json.Marshal(v)
}

func (v *AuthorizeAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "AuthorizeAction"

	return json.Marshal(*v)
}
