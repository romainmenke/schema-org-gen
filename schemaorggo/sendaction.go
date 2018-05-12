package schemaorggo

import "encoding/json"

// SendAction see : https://schema.org/SendAction
type SendAction struct {
	TransferAction

	typeContext

	// DeliveryMethod see : https://schema.org/deliveryMethod
	// A sub property of instrument. The method of delivery.
	DeliveryMethod *DeliveryMethod `json:"deliveryMethod,omitempty"`

	// Recipient see : https://schema.org/recipient
	// A sub property of participant. The participant who is at the receiving end of the action.
	Recipient interface{} `json:"recipient,omitempty"` // types : Audience ContactPoint Organization Person

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
