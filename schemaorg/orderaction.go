package schemaorg

import "encoding/json"

// OrderAction see : https://schema.org/OrderAction
type OrderAction struct {

typeContext

TradeAction

// DeliveryMethod see : https://schema.org/deliveryMethod
// A sub property of instrument. The method of delivery.
DeliveryMethod *DeliveryMethod `json:"deliveryMethod"`

}

func (v *OrderAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "OrderAction"

	return json.Marshal(v)
}
