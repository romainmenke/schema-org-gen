package schemaorggo

import "encoding/json"

// OrderAction see : https://schema.org/OrderAction
type OrderAction struct {
	TradeAction

	typeContext

	// DeliveryMethod see : https://schema.org/deliveryMethod
	// A sub property of instrument. The method of delivery.
	DeliveryMethod *DeliveryMethod `json:"deliveryMethod,omitempty"` // types : DeliveryMethod

}

func (v OrderAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "OrderAction"

	return json.Marshal(v)
}

func (v *OrderAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "OrderAction"

	return json.Marshal(*v)
}
