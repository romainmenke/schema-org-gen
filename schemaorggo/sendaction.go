package schemaorggo

import "encoding/json"

// SendAction see : https://schema.org/SendAction
type SendAction struct {
	TransferAction

	typeContext

	// DeliveryMethod see : https://schema.org/deliveryMethod
	// A sub property of instrument. The method of delivery.
	// types : DeliveryMethod
	DeliveryMethod *DeliveryMethod `json:"deliveryMethod,omitempty"`

	// Recipient see : https://schema.org/recipient
	// A sub property of participant. The participant who is at the receiving end of the action.
	// types : Audience ContactPoint Organization Person
	Recipient interface{} `json:"recipient,omitempty"`
}

func (v SendAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "SendAction"

	return json.Marshal(v)
}

func (v *SendAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "SendAction"

	return json.Marshal(*v)
}
