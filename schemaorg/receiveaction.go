package schemaorg

import "encoding/json"

// ReceiveAction see : https://schema.org/ReceiveAction
type ReceiveAction struct {

typeContext

TransferAction

// DeliveryMethod see : /deliveryMethod
// A sub property of instrument. The method of delivery.
DeliveryMethod *DeliveryMethod `json:"deliveryMethod"`

// Sender see : /sender
// A sub property of participant. The participant who is at the sending end of the action.
Sender interface{} `json:"sender"` // types : Audience Organization Person

}

func (v *ReceiveAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ReceiveAction"

	return json.Marshal(v)
}
