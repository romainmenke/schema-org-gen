package schemaorggo

import "encoding/json"

// PayAction see : https://schema.org/PayAction
type PayAction struct {
	TradeAction

	typeContext

	// Purpose see : http://health-lifesci.schema.org/purpose
	// A goal towards an action is taken. Can be concrete or abstract.
	Purpose interface{} `json:"purpose"` // types : MedicalDevicePurpose Thing

	// Recipient see : https://schema.org/recipient
	// A sub property of participant. The participant who is at the receiving end of the action.
	Recipient interface{} `json:"recipient"` // types : Audience ContactPoint Organization Person

}

func (v *PayAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "PayAction"

	return json.Marshal(v)
}
