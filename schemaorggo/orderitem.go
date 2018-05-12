package schemaorggo

import "encoding/json"

// OrderItem see : https://schema.org/OrderItem
type OrderItem struct {
	Intangible

	typeContext

	// OrderDelivery see : https://schema.org/orderDelivery
	// The delivery of the parcel related to this order or order item.
	// types : ParcelDelivery
	OrderDelivery *ParcelDelivery `json:"orderDelivery,omitempty"`

	// OrderItemNumber see : https://schema.org/orderItemNumber
	// The identifier of the order item.
	// types : Text
	OrderItemNumber string `json:"orderItemNumber,omitempty"`

	// OrderItemStatus see : https://schema.org/orderItemStatus
	// The current status of the order item.
	// types : OrderStatus
	OrderItemStatus *OrderStatus `json:"orderItemStatus,omitempty"`

	// OrderQuantity see : https://schema.org/orderQuantity
	// The number of the item ordered. If the property is not set, assume the quantity is one.
	// types : Number
	OrderQuantity float64 `json:"orderQuantity,omitempty"`

	// OrderedItem see : https://schema.org/orderedItem
	// The item ordered.
	// types : OrderItem Product
	OrderedItem interface{} `json:"orderedItem,omitempty"`
}

func (v OrderItem) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "OrderItem"

	return json.Marshal(v)
}

func (v *OrderItem) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "OrderItem"

	return json.Marshal(*v)
}
