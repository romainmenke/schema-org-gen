package schemaorggo

import "encoding/json"

// PayAction see : https://schema.org/PayAction
type PayAction struct {
	TradeAction

	typeContext

	// Purpose see : http://health-lifesci.schema.org/purpose
	// A goal towards an action is taken. Can be concrete or abstract.
	// types : MedicalDevicePurpose Thing
	Purpose interface{} `json:"purpose,omitempty"`

	// Recipient see : https://schema.org/recipient
	// A sub property of participant. The participant who is at the receiving end of the action.
	// types : Audience ContactPoint Organization Person
	Recipient interface{} `json:"recipient,omitempty"`
}

func (v PayAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "PayAction"

	return json.Marshal(v)
}

func (v *PayAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "PayAction"

	return json.Marshal(*v)
}
