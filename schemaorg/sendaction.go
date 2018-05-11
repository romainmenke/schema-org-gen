package schemaorg

import "encoding/json"

// SendAction see : https://schema.org/SendAction
type SendAction struct {

typeContext

TransferAction

// DeliveryMethod see : /deliveryMethod
// A sub property of instrument. The method of delivery.
DeliveryMethod *DeliveryMethod `json:"deliveryMethod"`

// Recipient see : /recipient
// A sub property of participant. The participant who is at the receiving end of the action.
Recipient interface{} `json:"recipient"` // types : Audience ContactPoint Organization Person

}

func (v *SendAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "SendAction"

	return json.Marshal(v)
}
