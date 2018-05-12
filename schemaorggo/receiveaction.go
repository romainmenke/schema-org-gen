package schemaorggo

import "encoding/json"

// ReceiveAction see : https://schema.org/ReceiveAction
type ReceiveAction struct {
	TransferAction

	typeContext

	// DeliveryMethod see : https://schema.org/deliveryMethod
	// A sub property of instrument. The method of delivery.
	// types : DeliveryMethod
	DeliveryMethod *DeliveryMethod `json:"deliveryMethod,omitempty"`

	// Sender see : https://schema.org/sender
	// A sub property of participant. The participant who is at the sending end of the action.
	// types : Audience Organization Person
	Sender interface{} `json:"sender,omitempty"`
}

func (v ReceiveAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ReceiveAction"

	return json.Marshal(v)
}

func (v *ReceiveAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "ReceiveAction"

	return json.Marshal(*v)
}
